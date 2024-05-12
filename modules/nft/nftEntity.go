package nft

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Nft struct {
		Id                primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Title             string             `json:"title" bson:"title"`
		Price             float64            `json:"price" bson:"price"`
		ImageUrl          string             `json:"image_url" bson:"image_url"`
		UsageStatus       bool               `json:"usage_status" bson:"usage_status"`
		Description       string             `json:"description" bson:"description"`
		AuthorId          primitive.ObjectID `json:"author_id" bson:"author_id"`
		OwnerId           primitive.ObjectID `json:"owner_id" bson:"owner_id"`
		Category          primitive.ObjectID `json:"category" bson:"category"`
		ListingType       string             `json:"listing_type" bson:"listing_type"`
		CreatedAt         time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt         time.Time          `json:"updated_at" bson:"updated_at"`
		IsDeleted         bool               `json:"is_deleted" bson:"is_deleted"`
		IsCategoryBlocked bool               `json:"is_category_blocked" bson:"is_category_blocked"`
	}

	NftCategory struct {
		Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Title       string             `json:"title" bson:"title"`
		Description string             `json:"description" bson:"description"`
		UsageStatus bool               `json:"usage_status" bson:"usage_status"`
		IsDeleted   bool               `json:"is_deleted" bson:"is_deleted"`
		CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	}
)
