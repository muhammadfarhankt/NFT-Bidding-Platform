package inventory

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Inventory struct {
		Id     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		UserId string             `json:"user_id" bson:"user_id"`
		NftId  string             `json:"nft_id" bson:"nft_id"`
	}
)
