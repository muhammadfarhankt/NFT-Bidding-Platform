package inventoryRepository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/inventory"
	nftPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftPb"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/grpcConn"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/jwtAuth"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	InventoryRepositoryService interface {
		FindNftsInIds(pctx context.Context, grpcUrl string, req *nftPb.FindNftsInIdsReq) (*nftPb.FindNftsInIdsRes, error)
		FindUserNfts(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*inventory.Inventory, error)
		CountUserNfts(pctx context.Context, userId string) (int64, error)
	}

	inventoryRepository struct {
		db *mongo.Client
	}
)

func NewInventoryRepository(db *mongo.Client) InventoryRepositoryService {
	return &inventoryRepository{db}
}

func (r *inventoryRepository) inventoryDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("inventory_db")
}

// FindNftsInIds find nfts in ids
func (r *inventoryRepository) FindNftsInIds(pctx context.Context, grpcUrl string, req *nftPb.FindNftsInIdsReq) (*nftPb.FindNftsInIdsRes, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	// Set api key in context
	jwtAuth.SetApiKeyInContext(&ctx)

	// Connect to grpc server
	conn, err := grpcConn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("failed to connect to grpc server: %s", err.Error())
		return nil, errors.New("failed to connect to grpc server")
	}

	// Find nfts in ids
	result, err := conn.Nft().FindNftsInIds(ctx, req)
	if err != nil {
		log.Printf("failed to find nfts in ids: %s", err.Error())
		return nil, errors.New("failed to find nfts in ids")
	}

	if result == nil {
		log.Printf("\nError: FindNftsInIds: nfts not found : %s", err.Error())
		return nil, errors.New("nfts not found")
	}

	if len(result.Nfts) == 0 {
		log.Printf("\nError: FindNftsInIds: nfts not found : %s", err.Error())
		return nil, errors.New("nfts not found")
	}

	return result, nil
}

// FindUserNfts find user nfts
func (r *inventoryRepository) FindUserNfts(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*inventory.Inventory, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	// Connect to inventory db and collection
	collection := r.inventoryDbConn(ctx).Collection("users_inventory")

	// Find user nfts
	cursors, err := collection.Find(ctx, filter, opts...)
	if err != nil {
		log.Printf("Error: failed to find user nfts: %s", err.Error())
		return nil, errors.New("failed to find user nfts")
	}

	results := make([]*inventory.Inventory, 0)

	// Decode user nfts
	for cursors.Next(ctx) {
		result := new(inventory.Inventory)
		if err := cursors.Decode(result); err != nil {
			log.Printf("failed to decode user nfts: %s", err.Error())
			return nil, errors.New("error: failed to decode user nfts")
		}
		results = append(results, result)
		// fmt.Println("result (inventoryRepository) : ", result)
	}

	return results, nil
}

// CountUserNfts count user nfts
func (r *inventoryRepository) CountUserNfts(pctx context.Context, userId string) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	// Connect to inventory db and collection
	collection := r.inventoryDbConn(ctx).Collection("users_inventory")

	// Count user nfts
	count, err := collection.CountDocuments(ctx, primitive.D{{"user_id", utils.ConvertToObjectId(userId)}})
	fmt.Println("user id : ", userId)

	if err != nil {
		log.Printf("Error: failed to count user nfts: %s", err.Error())
		return -1, errors.New("error: failed to count user nfts")
	}
	return count, nil
}
