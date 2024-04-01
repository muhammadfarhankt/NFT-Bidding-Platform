package inventory

import (
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/models"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
)

type (
	UpdateInventoryReq struct {
		UserId string `json:"user_id" validate:"required,max=64"`
		NftId  string `json:"nft_id" validate:"required,max=64"`
	}

	NftInInventory struct {
		InventoryId string `json:"inventory_id"`
		*nft.NftShowCase
	}

	UserInventory struct {
		UserId string `json:"user_id"`
		*models.PaginateRes
	}
)
