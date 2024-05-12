package inventoryUsecase

import (
	"context"
	"fmt"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/inventory"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/inventory/inventoryRepository"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/models"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
	nftPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftPb"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	InventoryUsecaseService interface {
		FindUserNfts(pctx context.Context, cfg *config.Config, userId string, req *inventory.InventorySearchReq) (*models.PaginateRes, error)
	}

	inventoryUsecase struct {
		inventoryRepository inventoryRepository.InventoryRepositoryService
	}
)

func NewInventoryUsecase(inventoryRepository inventoryRepository.InventoryRepositoryService) InventoryUsecaseService {
	return &inventoryUsecase{
		inventoryRepository: inventoryRepository,
	}
}

func (u *inventoryUsecase) FindUserNfts(pctx context.Context, cfg *config.Config, userId string, req *inventory.InventorySearchReq) (*models.PaginateRes, error) {
	// filter
	filter := bson.D{}

	if req.Start != "" {
		filter = append(filter, bson.E{"_id", bson.D{{"$gt", utils.ConvertToObjectId(req.Start)}}})
	}
	userIdTrimmed := userId[5:]
	filter = append(filter, bson.E{"user_id", utils.ConvertToObjectId(userIdTrimmed)})
	// filter = append(filter, bson.E{"author_id", utils.ConvertToObjectId(userIdTrimmed)})
	// fmt.Println("User ID: ", userIdTrimmed)

	// option
	opts := make([]*options.FindOptions, 0)

	opts = append(opts, options.Find().SetSort(bson.D{{"_id", 1}}))
	opts = append(opts, options.Find().SetLimit(int64(req.Limit)))

	// find
	inventoryData, err := u.inventoryRepository.FindUserNfts(pctx, filter, opts)
	if err != nil {
		return nil, err
	}

	nftData, err := u.inventoryRepository.FindNftsInIds(pctx, cfg.Grpc.NftUrl, &nftPb.FindNftsInIdsReq{
		Ids: func() []string {
			nftIds := make([]string, 0)
			for _, v := range inventoryData {
				nftIds = append(nftIds, v.NftId)
			}
			return nftIds
		}(),
	})

	nftMaps := make(map[string]*nft.NftShowCase)
	for _, v := range nftData.Nfts {
		nftMaps[v.Id] = &nft.NftShowCase{
			NftId:    v.Id,
			Title:    v.Title,
			Price:    v.Price,
			ImageUrl: v.ImageUrl,
		}
	}

	results := make([]*inventory.NftInInventory, 0)

	//fmt.Println("Inventory Data: ", inventoryData)
	//fmt.Println("NFT Data: ", nftData)
	// fmt.Println("NFT Maps: ", nftMaps)
	// fmt.Println("User ID Trimmed: ", userIdTrimmed)
	for key, val := range nftMaps {
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", val)
	}

	for _, v := range inventoryData {
		results = append(results, &inventory.NftInInventory{
			InventoryId: v.Id.Hex(),
			UserId:      v.UserId,
			NftShowCase: &nft.NftShowCase{
				NftId:    v.NftId,
				Title:    nftMaps[v.NftId].Title,
				Price:    nftMaps[v.NftId].Price,
				ImageUrl: nftMaps[v.NftId].ImageUrl,
				// Description: nftMaps[v.NftId].Description,
				// AuthorId:    nftMaps[v.NftId].AuthorId,
				// OwnerId:     nftMaps[v.NftId].OwnerId,
				// Category:    nftMaps[v.NftId].Category,
				// UsageStatus: nftMaps[v.NftId].UsageStatus,
			},
		})
		fmt.Println("nft id : ", v.NftId)
	}

	// count
	total, err := u.inventoryRepository.CountUserNfts(pctx, userIdTrimmed)
	fmt.Println("Total usecase : ", total)

	if err != nil {
		return nil, err
	}

	fmt.Println("Results: ", results)

	if len(results) == 0 {
		return &models.PaginateRes{
			Data:  results,
			Total: total,
			Limit: req.Limit,
			First: models.FirstPaginate{
				Href: fmt.Sprintf("%s/%s?limit=%d", cfg.Paginate.InventoryNextPageBasedUrl, userId, req.Limit),
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
			Href: fmt.Sprintf("%s/%s?limit=%d", cfg.Paginate.InventoryNextPageBasedUrl, userId, req.Limit),
		},
		Next: models.NextPaginate{
			Start: results[len(results)-1].InventoryId,
			Href:  fmt.Sprintf("%s/%s?limit=%d&start=%s", cfg.Paginate.InventoryNextPageBasedUrl, userId, req.Limit, results[len(results)-1].InventoryId),
		},
	}, nil
}
