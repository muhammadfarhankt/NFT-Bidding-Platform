package nftUsecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/models"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
	files "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
	nftPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftPb"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftRepository"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	NftUsecaseService interface {
		CreateNft(pctx context.Context, req *nft.CreateNftReq) (any, error)
		FindOneNft(pctx context.Context, nftId string) (*nft.NftShowCase, error)
		FindManyNfts(pctx context.Context, basePaginateUrl string, req *nft.NftSearchReq) (*models.PaginateRes, error)
		EditNft(pctx context.Context, nftId string, req *nft.NftUpdateReq) (*nft.NftShowCase, error)
		BlockOrUnblockNft(pctx context.Context, nftId string) (bool, error)
		DeleteNft(pctx context.Context, nftId string) (bool, error)

		FindNftsInIds(pctx context.Context, req *nftPb.FindNftsInIdsReq) (*nftPb.FindNftsInIdsRes, error)

		// -------------------- Category -------------------- //
		CreateCategory(pctx context.Context, req *nft.NftCategoryReq) (any, error)
		FindOneCategory(pctx context.Context, categoryId string) (*nft.NftCategory, error)
		FindManyCategories(pctx context.Context, basePaginateUrl string, req *nft.NftSearchReq) (*models.PaginateRes, error)
		EditCategory(pctx context.Context, categoryId string, req *nft.NftCategory) (*nft.NftCategory, error)
		BlockOrUnblockCategory(pctx context.Context, categoryId string) (bool, error)
		DeleteCategory(pctx context.Context, categoryId string) (bool, error)

		// -------------------- NFT Image -------------------- //
		UploadToGCP(req []*files.FileReq) ([]*files.FileRes, error)
		DeleteFileFromGCP(req []*files.DeleteFileReq) error
	}

	nftUsecase struct {
		nftRepository nftRepository.NftRepositoryService
	}
)

func NewNftUsecase(nftRepository nftRepository.NftRepositoryService) NftUsecaseService {
	return &nftUsecase{nftRepository}
}

func (u *nftUsecase) CreateNft(pctx context.Context, req *nft.CreateNftReq) (any, error) {
	if !u.nftRepository.IsUniqueNft(pctx, req.Title) {
		return nil, errors.New("error: this title is already exist")
	}

	// req.AuthorId = strings.TrimPrefix(claims.UserId, "user:")
	// req.OwnerId = strings.TrimPrefix(claims.UserId, "user:")
	authorId, _ := primitive.ObjectIDFromHex(req.AuthorId)
	ownerId, _ := primitive.ObjectIDFromHex(req.OwnerId)
	categoryId, _ := primitive.ObjectIDFromHex(req.Category)

	nftId, err := u.nftRepository.InsertOneNft(pctx, &nft.Nft{
		Title:       req.Title,
		Price:       req.Price,
		UsageStatus: true,
		IsDeleted:   false,
		AuthorId:    authorId,
		OwnerId:     ownerId,
		Category:    categoryId,
		ListingType: req.ListingType,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		CreatedAt:   utils.LocalTime(),
		UpdatedAt:   utils.LocalTime(),
	})
	if err != nil {
		return nil, err
	}
	// return nftId.Hex(), nil
	return u.FindOneNft(pctx, nftId.Hex())
}

func (u *nftUsecase) FindOneNft(pctx context.Context, nftId string) (*nft.NftShowCase, error) {
	result, err := u.nftRepository.FindOneNft(pctx, nftId)
	if err != nil {
		return nil, err
	}

	return &nft.NftShowCase{
		NftId:       "nft:" + result.Id.Hex(),
		Title:       result.Title,
		Price:       result.Price,
		Description: result.Description,
		ImageUrl:    result.ImageUrl,
		AuthorId:    result.AuthorId.Hex(),
		OwnerId:     result.OwnerId.Hex(),
		Category:    result.Category.Hex(),
		ListingType: result.ListingType,
		UsageStatus: result.UsageStatus,
	}, nil
}

func (u *nftUsecase) FindManyNfts(pctx context.Context, basePaginateUrl string, req *nft.NftSearchReq) (*models.PaginateRes, error) {
	findNftsFilter := bson.D{}
	findNftsOpts := make([]*options.FindOptions, 0)

	countNftsFilter := bson.D{}

	// Filter
	if req.Start != "" {
		req.Start = strings.TrimPrefix(req.Start, "nft:")
		findNftsFilter = append(findNftsFilter, bson.E{"_id", bson.D{{"$gt", utils.ConvertToObjectId(req.Start)}}})
	}

	if req.Title != "" {
		findNftsFilter = append(findNftsFilter, bson.E{"title", primitive.Regex{Pattern: req.Title, Options: "i"}})
		countNftsFilter = append(countNftsFilter, bson.E{"title", primitive.Regex{Pattern: req.Title, Options: "i"}})
	}

	if req.Category != "" {
		findNftsFilter = append(findNftsFilter, bson.E{"category", primitive.Regex{Pattern: req.Category, Options: "i"}})
		countNftsFilter = append(countNftsFilter, bson.E{"category", primitive.Regex{Pattern: req.Category, Options: "i"}})
	}

	findNftsFilter = append(findNftsFilter, bson.E{"usage_status", true})
	countNftsFilter = append(countNftsFilter, bson.E{"usage_status", true})

	// Options
	findNftsOpts = append(findNftsOpts, options.Find().SetSort(bson.D{{"_id", 1}}))
	findNftsOpts = append(findNftsOpts, options.Find().SetLimit(int64(req.Limit)))

	// Find
	results, err := u.nftRepository.FindManyNfts(pctx, findNftsFilter, findNftsOpts)
	if err != nil {
		return nil, err
	}

	// Count
	total, err := u.nftRepository.CountNfts(pctx, countNftsFilter)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return &models.PaginateRes{
			Data:  make([]*nft.NftShowCase, 0),
			Total: total,
			Limit: req.Limit,
			First: models.FirstPaginate{
				Href: fmt.Sprintf("%s?limit=%d&title=%s", basePaginateUrl, req.Limit, req.Title),
			},
			Next: models.NextPaginate{
				Start: "",
				Href:  "",
			},
		}, nil
	}

	return &models.PaginateRes{
		Data:  results,
		Total: total,
		Limit: req.Limit,
		First: models.FirstPaginate{
			Href: fmt.Sprintf("%s?limit=%d&title=%s", basePaginateUrl, req.Limit, req.Title),
		},
		Next: models.NextPaginate{
			Start: results[len(results)-1].NftId,
			Href:  fmt.Sprintf("%s?limit=%d&title=%s&start=%s", basePaginateUrl, req.Limit, req.Title, results[len(results)-1].NftId),
		},
	}, nil
}

func (u *nftUsecase) EditNft(pctx context.Context, nftId string, req *nft.NftUpdateReq) (*nft.NftShowCase, error) {
	// Update logical
	updateReq := bson.M{}
	if req.Title != "" {
		if !u.nftRepository.IsUniqueNft(pctx, req.Title) {
			log.Println("Error: EditNft failed: this title is already exist")
			return nil, errors.New("error: this title is already exist")
		}

		updateReq["title"] = req.Title
	}
	if req.ImageUrl != "" {
		updateReq["image_url"] = req.ImageUrl
	}
	if req.Description != "" {
		updateReq["description"] = req.Description
	}
	// if req.Category != "" {
	// 	updateReq["category"] = req.Category
	// }
	if req.ListingType != "" {
		updateReq["listing_type"] = req.ListingType
	}
	if req.Price >= 0 {
		updateReq["price"] = req.Price
	}
	updateReq["updated_at"] = utils.LocalTime()

	if err := u.nftRepository.UpdateOneNft(pctx, nftId, updateReq); err != nil {
		return nil, err
	}

	return u.FindOneNft(pctx, nftId)
}

func (u *nftUsecase) BlockOrUnblockNft(pctx context.Context, nftId string) (bool, error) {
	result, err := u.nftRepository.FindOneNft(pctx, nftId)
	if err != nil {
		return false, err
	}

	if err := u.nftRepository.BlockOrUnblockNft(pctx, nftId, !result.UsageStatus); err != nil {
		return false, err
	}

	return !result.UsageStatus, nil
}

func (u *nftUsecase) DeleteNft(pctx context.Context, nftId string) (bool, error) {
	_, err := u.nftRepository.FindOneNft(pctx, nftId)
	if err != nil {
		return false, err
	}

	if err := u.nftRepository.DeleteNft(pctx, nftId); err != nil {
		return false, err
	}

	return false, nil
}

func (u *nftUsecase) FindNftsInIds(pctx context.Context, req *nftPb.FindNftsInIdsReq) (*nftPb.FindNftsInIdsRes, error) {

	filter := bson.D{}

	objectIds := make([]primitive.ObjectID, 0)
	for _, nftId := range req.Ids {
		objectIds = append(objectIds, utils.ConvertToObjectId(strings.TrimPrefix(nftId, "nft:")))
		log.Println("nftId: (FindNftsInIds nftUseCase)", nftId)
	}

	filter = append(filter, bson.E{"_id", bson.D{{"$in", objectIds}}})
	filter = append(filter, bson.E{"usage_status", true})
	// filter = append(filter, bson.E{"is_deleted", false})
	// filter = append(filter, bson.E{"is_category_blocked", false})

	results, err := u.nftRepository.FindManyNfts(pctx, filter, nil)
	if err != nil {
		return nil, err
	}

	resultsToRes := make([]*nftPb.Nft, 0)
	for _, result := range results {
		resultsToRes = append(resultsToRes, &nftPb.Nft{
			Id:    result.NftId,
			Title: result.Title,
			Price: result.Price,
			// Description: result.Description,
			ImageUrl: result.ImageUrl,
			// AuthorId:    result.AuthorId.Hex(),
			// OwnerId:     result.OwnerId.Hex(),
			// Category:    result.Category.Hex(),
			// ListingType: result.ListingType,
			// UsageStatus: result.UsageStatus,
		})
	}

	return &nftPb.FindNftsInIdsRes{
		Nfts: resultsToRes,
	}, nil
}
