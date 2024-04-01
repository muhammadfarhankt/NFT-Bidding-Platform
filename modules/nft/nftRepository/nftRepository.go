package nftRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	NftRepositoryService interface{}

	nftRepository struct {
		db *mongo.Client
	}
)

func NewNftRepository(db *mongo.Client) NftRepositoryService {
	return &nftRepository{db}
}

func (r *nftRepository) nftDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("nft_db")
}
