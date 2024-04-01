package nft

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Item struct {
		Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Title       string             `json:"title" bson:"title"`
		Price       float64            `json:"price" bson:"price"`
		ImageUrl    string             `json:"image_url" bson:"image_url"`
		UsageStatus bool               `json:"usage_status" bson:"usage_status"`
		Description string             `json:"description" bson:"description"`
		AuthorId    string             `json:"author_id" bson:"author_id"`
		OwnerId     string             `json:"owner_id" bson:"owner_id"`
		Category    string             `json:"category" bson:"category"`
		ListingType string             `json:"listing_type" bson:"listing_type"`
		CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	}
)
