package nftRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	NftRepositoryService interface {
		// ------------------- NFT ------------------- //
		IsUniqueNft(pctx context.Context, title string) bool
		InsertOneNft(pctx context.Context, req *nft.Nft) (primitive.ObjectID, error)
		FindOneNft(pctx context.Context, nftId string) (*nft.Nft, error)
		FindManyNfts(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*nft.NftShowCase, error)
		CountNfts(pctx context.Context, filter primitive.D) (int64, error)
		UpdateOneNft(pctx context.Context, nftId string, req primitive.M) error
		BlockOrUnblockNft(pctx context.Context, nftId string, isActive bool) error
		DeleteNft(pctx context.Context, nftId string) error

		FindTopWishlistNfts(pctx context.Context, filter primitive.D) ([]*nft.NftShowCase, error)
		FindTopBiddingNfts(pctx context.Context, filter primitive.D) ([]*nft.NftShowCase, error)

		// ------------------- NFT gRPC Method ------------------- //
		IncrementWishlistCount(pctx context.Context, nftId string) error

		// ------------------- Category ------------------- //
		InsertOneCategory(pctx context.Context, req *nft.NftCategory) (primitive.ObjectID, error)
		FindOneCategory(pctx context.Context, categoryId string) (*nft.NftCategory, error)
		FindManyCategories(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*nft.NftCategory, error)
		UpdateOneCategory(pctx context.Context, categoryId string, req primitive.M) error
		BlockOrUnblockCategory(pctx context.Context, categoryId string, isActive bool) error
		DeleteCategory(pctx context.Context, categoryId string) error

		// ------------------- Bidding ------------------- //
		FindManyBids(pctx context.Context, userId string) ([]*nft.Bid, error)
		CreateBid(pctx context.Context, userId string, req *nft.CreateBidReq) (primitive.ObjectID, error)
		FindOneBid(pctx context.Context, bidId string) (*nft.Bid, error)
		EditBid(pctx context.Context, bidId string, req primitive.M) error
		DeleteBid(pctx context.Context, bidId string) error

		// ------------------- NFT Bidding User ------------------- //
		FindOneUserBid(pctx context.Context, userId, bidId string) (*nft.SingleBid, error)
		FindUserBids(pctx context.Context, userId string) (any, error)
		BidNft(pctx context.Context, userId, nftId, price string) (primitive.ObjectID, error)
		WithdrawBid(pctx context.Context, bidId string) error

		// ------------------- NFT Bidding Admin ------------------- //
		ExecuteBids(pctx context.Context) (any, error)
	}

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

// ------------------- NFT gRPC Method ------------------- //
func (r *nftRepository) IncrementWishlistCount(pctx context.Context, nftId string) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("nfts")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(nftId)}, bson.M{"$inc": bson.M{"wishlist_count": 1}})

	if result.ModifiedCount == 0 {
		log.Printf("Error: IncrementWishlistCount failed: %s", err.Error())
		return errors.New("error: increment wishlist count failed")
	}

	if err != nil {
		log.Printf("Error: IncrementWishlistCount failed: %s", err.Error())
		return errors.New("error: increment wishlist count failed")
	}

	return nil
}

func (r *nftRepository) IsUniqueNft(pctx context.Context, title string) bool {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("nfts")

	result := new(nft.Nft)
	if err := col.FindOne(ctx, bson.M{"title": title}).Decode(result); err != nil {
		log.Printf("Error: IsUniqueNft: %s", err.Error())
		return true
	}
	return false
}

func (r *nftRepository) InsertOneNft(pctx context.Context, req *nft.Nft) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("nfts")

	nftId, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOneNft: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one nft failed")
	}

	return nftId.InsertedID.(primitive.ObjectID), nil
}

func (r *nftRepository) FindOneNft(pctx context.Context, nftId string) (*nft.Nft, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("nfts")

	result := new(nft.Nft)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(nftId)}).Decode(result); err != nil {
		log.Printf("Error: FindOneNft failed: %s", err.Error())
		return nil, errors.New("error: nft not found")
	}

	return result, nil
}

func (r *nftRepository) FindManyNfts(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*nft.NftShowCase, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("nfts")

	cursors, err := col.Find(ctx, filter, opts...)
	if err != nil {
		log.Printf("Error: FindManyNfts failed: %s", err.Error())
		return make([]*nft.NftShowCase, 0), errors.New("error: find many nfts failed")
	}

	results := make([]*nft.NftShowCase, 0)
	for cursors.Next(ctx) {
		result := new(nft.Nft)
		if err := cursors.Decode(result); err != nil {
			log.Printf("Error: FindManyNfts failed: %s", err.Error())
			return make([]*nft.NftShowCase, 0), errors.New("error: find many nfts failed")
		}

		results = append(results, &nft.NftShowCase{
			NftId:       "nft:" + result.Id.Hex(),
			Title:       result.Title,
			Price:       result.Price,
			Description: result.Description,
			AuthorId:    result.AuthorId.Hex(),
			OwnerId:     result.OwnerId.Hex(),
			Category:    result.Category.Hex(),
			ListingType: result.ListingType,
			UsageStatus: result.UsageStatus,
			ImageUrl:    result.ImageUrl,
		})
	}

	return results, nil
}

func (r *nftRepository) CountNfts(pctx context.Context, filter primitive.D) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("nfts")

	count, err := col.CountDocuments(ctx, filter)
	if err != nil {
		log.Printf("Error: CountNfts failed: %s", err.Error())
		return -1, errors.New("error: count nfts failed")
	}

	return count, nil
}

func (r *nftRepository) UpdateOneNft(pctx context.Context, nftId string, req primitive.M) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("nfts")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(nftId)}, bson.M{"$set": req})
	if err != nil {
		log.Printf("Error: UpdateOneNft failed: %s", err.Error())
		return errors.New("error: update one nft failed")
	}
	log.Printf("UpdateOneNft result: %v", result.ModifiedCount)

	return nil
}

func (r *nftRepository) BlockOrUnblockNft(pctx context.Context, nftId string, isActive bool) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("nfts")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(nftId)}, bson.M{"$set": bson.M{"usage_status": isActive}})
	if err != nil {
		log.Printf("Error: BlockOrUnblockNft failed: %s", err.Error())
		return errors.New("error: BlockOrUnblockNft failed")
	}
	log.Printf("BlockOrUnblockNft result: %v", result.ModifiedCount)

	return nil
}

func (r *nftRepository) DeleteNft(pctx context.Context, nftId string) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("nfts")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(nftId)}, bson.M{"$set": bson.M{"usage_status": false, "is_deleted": true}})
	if err != nil {
		log.Printf("Error: DeleteNft failed: %s", err.Error())
		return errors.New("error: DeleteNft failed")
	}
	log.Printf("DeleteNft result: %v", result.ModifiedCount)

	return nil
}

func (r *nftRepository) FindTopWishlistNfts(pctx context.Context, filter primitive.D) ([]*nft.NftShowCase, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("nfts")

	// find top 10 wishlist nfts
	cursor, err := col.Find(ctx, filter, options.Find().SetSort(bson.D{{"wishlist_count", -1}}).SetLimit(10))
	if err != nil {
		log.Printf("Error: FindTopWishlistNfts failed: %s", err.Error())
		return make([]*nft.NftShowCase, 0), errors.New("error: find top wishlist nfts failed")
	}

	results := make([]*nft.NftShowCase, 0)
	for cursor.Next(ctx) {
		result := new(nft.Nft)
		if err := cursor.Decode(result); err != nil {
			log.Printf("Error: FindTopWishlistNfts failed: %s", err.Error())
			return make([]*nft.NftShowCase, 0), errors.New("error: find top wishlist nfts failed")
		}

		results = append(results, &nft.NftShowCase{
			NftId:         result.Id.Hex(),
			Title:         result.Title,
			Price:         result.Price,
			Description:   result.Description,
			AuthorId:      result.AuthorId.Hex(),
			OwnerId:       result.OwnerId.Hex(),
			Category:      result.Category.Hex(),
			ListingType:   result.ListingType,
			UsageStatus:   result.UsageStatus,
			ImageUrl:      result.ImageUrl,
			WishlistCount: result.WishlistCount,
			BidCount:      result.BidCount,
		})
	}

	return results, nil

}

func (r *nftRepository) FindTopBiddingNfts(pctx context.Context, filter primitive.D) ([]*nft.NftShowCase, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("nfts")

	// find top 10 bidding nfts
	cursor, err := col.Find(ctx, filter, options.Find().SetSort(bson.D{{"bid_count", -1}}).SetLimit(10))
	if err != nil {
		log.Printf("Error: FindTopBiddingNfts failed: %s", err.Error())
		return make([]*nft.NftShowCase, 0), errors.New("error: find top bidding nfts failed")
	}

	results := make([]*nft.NftShowCase, 0)
	for cursor.Next(ctx) {
		result := new(nft.Nft)
		if err := cursor.Decode(result); err != nil {
			log.Printf("Error: FindTopBiddingNfts failed: %s", err.Error())
			return make([]*nft.NftShowCase, 0), errors.New("error: find top bidding nfts failed")
		}

		results = append(results, &nft.NftShowCase{
			NftId:         result.Id.Hex(),
			Title:         result.Title,
			Price:         result.Price,
			Description:   result.Description,
			AuthorId:      result.AuthorId.Hex(),
			OwnerId:       result.OwnerId.Hex(),
			Category:      result.Category.Hex(),
			ListingType:   result.ListingType,
			UsageStatus:   result.UsageStatus,
			ImageUrl:      result.ImageUrl,
			WishlistCount: result.WishlistCount,
			BidCount:      result.BidCount,
		})
	}

	return results, nil
}
