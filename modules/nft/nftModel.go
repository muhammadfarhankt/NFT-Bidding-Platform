package nft

import (
	"time"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/models"
)

type (
	CreateNftReq struct {
		Title       string  `json:"title" validate:"required,max=64"`
		Price       float64 `json:"price" validate:"required"`
		ImageUrl    string  `json:"image_url" validate:"required,max=255"`
		Description string  `json:"description" validate:"required,max=255"`
		Category    string  `json:"category" validate:"required"`
		ListingType string  `json:"listing_type" validate:"required"`
		UsageStatus bool    `json:"usage_status"`
	}

	NftShowCase struct {
		NftId             string  `json:"nft_id"`
		Title             string  `json:"title"`
		Price             float64 `json:"price"`
		ImageUrl          string  `json:"image_url"`
		Description       string  `json:"description"`
		AuthorId          string  `json:"author_id"`
		OwnerId           string  `json:"owner_id"`
		Category          string  `json:"category"`
		ListingType       string  `json:"listing_type"`
		UsageStatus       bool    `json:"usage_status"`
		IsDeleted         bool    `json:"is_deleted"`
		IsCategoryBlocked bool    `json:"is_category_blocked"`
		WishlistCount     int     `json:"wishlist_count"`
		BidCount          int     `json:"bid_count"`
	}

	NftSearchReq struct {
		Title    string `query:"title" validate:"max=64"`
		Category string `query:"category" validate:"max=64"`
		models.PaginateReq
	}

	NftUpdateReq struct {
		Title       string  `json:"title" validate:"required,max=64"`
		Price       float64 `json:"price" validate:"required"`
		ImageUrl    string  `json:"image_url" validate:"required,max=255"`
		Description string  `json:"description" validate:"required,max=255"`
		Category    string  `json:"category" validate:"required"`
		ListingType string  `json:"listing_type" validate:"required"`
	}

	EnableOrDisableNftReq struct {
		UsageStatus bool `json:"usage_status"`
	}

	NftCategoryReq struct {
		Title       string `json:"title" validate:"required,max=64"`
		Description string `json:"description" validate:"required,max=255"`
	}

	EnableOrDisableNftCategoryReq struct {
		UsageStatus bool `json:"usage_status"`
	}

	BidShowCase struct {
		BidId      string    `json:"bid_id"`
		NftId      string    `json:"nft_id"`
		Price      float64   `json:"price"`
		ExpiryDate time.Time `json:"expiry_date"`
		IsDeleted  bool      `json:"is_deleted"`
	}

	CreateBidReq struct {
		NftId      string    `json:"nft_id" validate:"required"`
		Price      float64   `json:"price" validate:"required"`
		ExpiryDate time.Time `json:"expiry_date" validate:"required"`
	}

	CreateUserBidReq struct {
		NftId string  `json:"nft_id" validate:"required"`
		Price float64 `json:"price" validate:"required"`
	}
)
