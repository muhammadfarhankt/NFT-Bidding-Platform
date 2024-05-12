package nftUsecase

import (
	"context"
	"errors"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (u *nftUsecase) CreateCategory(pctx context.Context, req *nft.NftCategoryReq) (any, error) {
	// if !u.nftRepository.IsUniqueNft(pctx, req.Title) {
	// 	return nil, errors.New("error: this title is already exist")
	// }

	categoryId, err := u.nftRepository.InsertOneCategory(pctx, &nft.NftCategory{
		Title:       req.Title,
		UsageStatus: true,
		Description: req.Description,
		CreatedAt:   utils.LocalTime(),
		UpdatedAt:   utils.LocalTime(),
	})
	if err != nil {
		return nil, err
	}
	//return categoryId.Hex(), nil
	return u.FindOneCategory(pctx, categoryId.Hex())
}

func (u *nftUsecase) FindOneCategory(pctx context.Context, categoryId string) (*nft.NftCategory, error) {
	result, err := u.nftRepository.FindOneCategory(pctx, categoryId)
	if err != nil {
		return nil, err
	}

	return &nft.NftCategory{
		Id:          result.Id,
		Title:       result.Title,
		Description: result.Description,
		UsageStatus: result.UsageStatus,
		IsDeleted:   result.IsDeleted,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}, nil
}

func (u *nftUsecase) FindManyCategories(pctx context.Context, basePaginateUrl string, req *nft.NftSearchReq) ([]*nft.NftCategory, error) {
	findNftsFilter := bson.D{}
	findNftsOpts := make([]*options.FindOptions, 0)

	findNftsFilter = append(findNftsFilter, bson.E{"usage_status", true})
	findNftsFilter = append(findNftsFilter, bson.E{"is_deleted", false})

	// Options
	findNftsOpts = append(findNftsOpts, options.Find().SetSort(bson.D{{"_id", 1}}))

	// Find
	results, err := u.nftRepository.FindManyCategories(pctx, findNftsFilter, findNftsOpts)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return make([]*nft.NftCategory, 0), errors.New("error: categories not found")
	}

	return results, nil
}

func (u *nftUsecase) EditCategory(pctx context.Context, categoryId string, req *nft.NftCategory) (*nft.NftCategory, error) {
	// Update logical
	updateReq := bson.M{}
	if req.Title != "" {
		// if !u.nftRepository.IsUniqueNft(pctx, req.Title) {
		// 	log.Println("Error: EditNft failed: this title is already exist")
		// 	return nil, errors.New("error: this title is already exist")
		// }

		updateReq["title"] = req.Title
	}

	if req.Description != "" {
		updateReq["description"] = req.Description
	}
	// if req.Category != "" {
	// 	updateReq["category"] = req.Category
	// }

	updateReq["updated_at"] = utils.LocalTime()

	if err := u.nftRepository.UpdateOneCategory(pctx, categoryId, updateReq); err != nil {
		return nil, err
	}

	return u.FindOneCategory(pctx, categoryId)
}

func (u *nftUsecase) BlockOrUnblockCategory(pctx context.Context, categoryId string) (bool, error) {
	result, err := u.nftRepository.FindOneCategory(pctx, categoryId)
	if err != nil {
		return false, err
	}

	if err := u.nftRepository.BlockOrUnblockCategory(pctx, categoryId, !result.UsageStatus); err != nil {
		return false, err
	}

	return !result.UsageStatus, nil
}

func (u *nftUsecase) DeleteCategory(pctx context.Context, categoryId string) (bool, error) {
	// commenting below code because category check already done in handler
	// _, err := u.nftRepository.FindOneCategory(pctx, categoryId)
	// if err != nil {
	// 	return false, err
	// }

	if err := u.nftRepository.DeleteCategory(pctx, categoryId); err != nil {
		return false, err
	}

	return false, nil
}
