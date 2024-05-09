package nftHandler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	files "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/response"
)

func RandFileName(ext string) string {
	fileName := fmt.Sprintf("%s_%v", strings.ReplaceAll(uuid.NewString()[:6], "-", ""), time.Now().UnixMilli())
	if ext != "" {
		fileName = fmt.Sprintf("%s.%s", fileName, ext)
	}
	return fileName
}

func (f *nftHttpHandler) UploadToGCP(c echo.Context) error {
	fmt.Println("uploadToGCP")
	req := make([]*files.FileReq, 0)
	form, err := c.MultipartForm()
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "uploadToGCPErr : "+err.Error())
	}

	filesReq := form.File["files"]
	destination := c.FormValue("destination")

	if filesReq == nil {
		return response.ErrResponse(c, http.StatusBadRequest, "uploadToGCPErr : file/ files is required")
	}

	if destination == "" {
		return response.ErrResponse(c, http.StatusBadRequest, "uploadToGCPErr : destination is required")
	}

	// files  extension validaton
	imageExtensions := map[string]bool{
		"jpg":  true,
		"jpeg": true,
		"png":  true,
	}

	for _, file := range filesReq {
		extension := strings.TrimPrefix(filepath.Ext(file.Filename), ".")
		if _, ok := imageExtensions[extension]; !ok {
			return response.ErrResponse(c, http.StatusBadRequest, "uploadToGCPErr : invalid file extension")
		}
		// 4097kb
		if file.Size > int64(4097000) {
			return response.ErrResponse(c, http.StatusBadRequest, "uploadToGCPErr : file size too large")
		}

		filename := RandFileName(extension)
		req = append(req, &files.FileReq{
			File:        file,
			Destination: destination + "/" + filename,
			FileName:    filename,
			Extension:   extension,
		},
		)
	}
	// for _, file := range filesReq {
	// 	fmt.Println("file.Filename", file.Filename)
	// 	fmt.Println("file.Destination", destination)
	// }

	res, err := f.nftUsecase.UploadToGCP(req)
	// res, err := "success", nil
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, "uploadToGCPErr : "+err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (f *nftHttpHandler) DeleteFromGCP(c echo.Context) error {

	req := make([]*files.DeleteFileReq, 0)

	if err := c.Bind(&req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "deleteFromGCPErr Bind Error : "+err.Error())
	}

	for _, file := range req {
		fmt.Println("file.Destination", file.Destination)
	}

	if err := f.nftUsecase.DeleteFileFromGCP(req); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, "deleteFromGCPErr : "+err.Error())
	}

	// return entities.NewResponse(c).Success(
	// 	fiber.StatusOK,
	// 	"File Deleted Successfully",
	// ).Res()
	return response.SuccessResponse(c, http.StatusOK, "File Deleted Successfully")
}
