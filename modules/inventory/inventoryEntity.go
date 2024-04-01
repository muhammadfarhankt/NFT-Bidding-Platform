package inventory

type (
	Inventory struct {
		Id       string `json:"_id" bson:"_id,omitempty"`
		UserId string `json:"user_id" bson:"user_id"`
		NftId   string `json:"nft_id" bson:"nft_id"`
	}
)
