package nft

import "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/models"

type (
	CreateItemReq struct {
		Title       string  `json:"title" validate:"required,max=64"`
		Price       float64 `json:"price" validate:"required"`
		ImageUrl    string  `json:"image_url" validate:"required,max=255"`
		Description string  `json:"description" validate:"required,max=255"`
		AuthorId    string  `json:"author_id" validate:"required"`
		OwnerId     string  `json:"owner_id" validate:"required"`
		Category    string  `json:"category" validate:"required"`
		ListingType string  `json:"listing_type" validate:"required"`
	}

	ItemShowCase struct {
		ItemId      string  `json:"item_id"`
		Title       string  `json:"title"`
		Price       float64 `json:"price"`
		ImageUrl    string  `json:"image_url"`
		Description string  `json:"description"`
		AuthorId    string  `json:"author_id"`
		OwnerId     string  `json:"owner_id"`
		Category    string  `json:"category"`
		ListingType string  `json:"listing_type"`
	}

	ItemSearchReq struct {
		Title string `json:"title"`
		models.PaginateReq
	}

	ItemUpdateReq struct {
		Title       string  `json:"title" validate:"required,max=64"`
		Price       float64 `json:"price" validate:"required"`
		ImageUrl    string  `json:"image_url" validate:"required,max=255"`
		Description string  `json:"description" validate:"required,max=255"`
		Category    string  `json:"category" validate:"required"`
		ListingType string  `json:"listing_type" validate:"required"`
	}

	EnableOrDisableItemReq struct {
		UsageStatus bool `json:"usage_status"`
	}
)
