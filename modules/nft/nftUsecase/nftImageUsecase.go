package nftUsecase

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"

	files "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
)

type filesPublic struct {
	bucket      string
	destination string
	file        *files.FileRes
}

func (u *nftUsecase) streamFileUpload(ctx context.Context, client *storage.Client, jobs <-chan *files.FileReq, results chan<- *files.FileRes, errs chan<- error) {

	for job := range jobs {
		container, err := job.File.Open()
		if err != nil {
			errs <- fmt.Errorf("error job.File.Open: %w", err)
			return
		}
		b, err := io.ReadAll(container)
		if err != nil {
			errs <- fmt.Errorf("error io.ReadAll: %w", err)
			return
		}

		buf := bytes.NewBuffer(b)

		// Upload an object with storage.Writer.
		wc := client.Bucket("nft-marketplace-dev-bucket").Object(job.Destination).NewWriter(ctx)
		//wc.ChunkSize = 0 // note retries are not supported for chunk size 0.

		if _, err = io.Copy(wc, buf); err != nil {
			errs <- fmt.Errorf("error io.Copy: %w", err)
			return
		}
		// Data can continue to be added to the file until the writer is closed.
		if err := wc.Close(); err != nil {
			errs <- fmt.Errorf("error Writer.Close: %w", err)
			return
		}
		fmt.Printf("%v uploaded to %v.\n", job.FileName, job.Extension)

		newFile := &filesPublic{
			file: &files.FileRes{
				FileName: job.FileName,
				Url:      fmt.Sprintf("https://storage.googleapis.com/%s/%s", "nft-marketplace-dev-bucket", job.Destination),
			},
			bucket:      "nft-marketplace-dev-bucket",
			destination: job.Destination,
		}

		if err := newFile.makePublic(ctx, client); err != nil {
			errs <- fmt.Errorf("error newFile.makePublic: %w", err)
			return
		}

		errs <- nil
		results <- newFile.file
	}
}

// makePublic gives all users read access to an object.
func (f *filesPublic) makePublic(ctx context.Context, client *storage.Client) error {

	acl := client.Bucket(f.bucket).Object(f.destination).ACL()
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return fmt.Errorf("error : ACLHandle.Set: %w", err)
	}
	fmt.Printf("Blob %v is now publicly accessible.\n", f.destination)
	return nil
}

func (u *nftUsecase) UploadToGCP(req []*files.FileReq) ([]*files.FileRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	jobsCh := make(chan *files.FileReq, len(req))
	resultsCh := make(chan *files.FileRes, len(req))
	errsCh := make(chan error, len(req))

	results := make([]*files.FileRes, 0)

	for _, request := range req {
		jobsCh <- request
	}
	close(jobsCh)

	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		//  worker
		go u.streamFileUpload(ctx, client, jobsCh, resultsCh, errsCh)

	}

	for a := 0; a < len(req); a++ {
		err := <-errsCh
		if err != nil {
			return nil, fmt.Errorf("errChannel error: %w", err)
		}
		result := <-resultsCh
		results = append(results, result)
	}

	return results, nil
}

// deleteFile removes specified object.
func (u *nftUsecase) deleteFileWorker(ctx context.Context, client *storage.Client, jobs <-chan *files.DeleteFileReq, errs chan<- error) {

	for job := range jobs {
		o := client.Bucket("nft-marketplace-dev-bucket").Object(job.Destination)

		// Optional: set a generation-match precondition to avoid potential race
		// conditions and data corruptions. The request to delete the file is aborted
		// if the object's generation number does not match your precondition.
		attrs, err := o.Attrs(ctx)
		if err != nil {
			errs <- fmt.Errorf("object.Attrs: %w", err)
			return
		}
		o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

		if err := o.Delete(ctx); err != nil {
			errs <- fmt.Errorf("object(%q).Delete: %w", job.Destination, err)
			return
		}
		fmt.Printf("Blob %v deleted.\n", job.Destination)
		errs <- nil
	}
}

func (u *nftUsecase) DeleteFileFromGCP(req []*files.DeleteFileReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	jobsCh := make(chan *files.DeleteFileReq, len(req))
	errsCh := make(chan error, len(req))

	for _, request := range req {
		jobsCh <- request
	}
	close(jobsCh)

	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		//  worker
		go u.deleteFileWorker(ctx, client, jobsCh, errsCh)

	}

	for a := 0; a < len(req); a++ {
		err := <-errsCh
		if err != nil {
			return fmt.Errorf("errChannel error: %w", err)
		}
	}
	return nil
}
