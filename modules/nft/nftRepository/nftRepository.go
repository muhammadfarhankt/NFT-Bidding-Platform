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
		IsUniqueNft(pctx context.Context, title string) bool
		InsertOneNft(pctx context.Context, req *nft.Nft) (primitive.ObjectID, error)
		FindOneNft(pctx context.Context, nftId string) (*nft.Nft, error)
		FindManyNfts(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*nft.NftShowCase, error)
		CountNfts(pctx context.Context, filter primitive.D) (int64, error)
		UpdateOneNft(pctx context.Context, nftId string, req primitive.M) error
		BlockOrUnblockNft(pctx context.Context, nftId string, isActive bool) error
		InsertOneCategory(pctx context.Context, req *nft.NftCategory) (primitive.ObjectID, error)
		FindOneCategory(pctx context.Context, categoryId string) (*nft.NftCategory, error)
		FindManyCategories(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*nft.NftCategory, error)
		UpdateOneCategory(pctx context.Context, categoryId string, req primitive.M) error
		BlockOrUnblockCategory(pctx context.Context, categoryId string, isActive bool) error
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
			AuthorId:    result.AuthorId,
			OwnerId:     result.OwnerId,
			Category:    result.Category,
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

func (r *nftRepository) InsertOneCategory(pctx context.Context, req *nft.NftCategory) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("categories")

	nftId, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOneCategory: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one category failed")
	}

	return nftId.InsertedID.(primitive.ObjectID), nil
}

func (r *nftRepository) FindOneCategory(pctx context.Context, categoryId string) (*nft.NftCategory, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("categories")

	result := new(nft.NftCategory)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(categoryId)}).Decode(result); err != nil {
		log.Printf("Error: FindOneCategory failed: %s", err.Error())
		return nil, errors.New("error: category not found")
	}

	return result, nil
}

func (r *nftRepository) FindManyCategories(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*nft.NftCategory, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("categories")

	cursors, err := col.Find(ctx, filter, opts...)
	if err != nil {
		log.Printf("Error: FindManyCategories failed: %s", err.Error())
		return make([]*nft.NftCategory, 0), errors.New("error: find many categories failed")
	}

	results := make([]*nft.NftCategory, 0)
	for cursors.Next(ctx) {
		result := new(nft.NftCategory)
		if err := cursors.Decode(result); err != nil {
			log.Printf("Error: FindManyCategories failed: %s", err.Error())
			return make([]*nft.NftCategory, 0), errors.New("error: FindManyCategories failed")
		}

		results = append(results, &nft.NftCategory{
			Id:          result.Id,
			Title:       result.Title,
			Description: result.Description,
			UsageStatus: result.UsageStatus,
		})
	}

	return results, nil
}

func (r *nftRepository) UpdateOneCategory(pctx context.Context, categoryId string, req primitive.M) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("categories")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(categoryId)}, bson.M{"$set": req})
	if err != nil {
		log.Printf("Error: UpdateOneCategory failed: %s", err.Error())
		return errors.New("error: update one category failed")
	}
	log.Printf("UpdateOneCategory result: %v", result.ModifiedCount)

	return nil
}

func (r *nftRepository) BlockOrUnblockCategory(pctx context.Context, categoryId string, isActive bool) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("categories")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(categoryId)}, bson.M{"$set": bson.M{"usage_status": isActive}})
	if err != nil {
		log.Printf("Error: BlockOrUnblockCategory failed: %s", err.Error())
		return errors.New("error: BlockOrUnblockCategory failed")
	}
	log.Printf("BlockOrUnblockNft result: %v", result.ModifiedCount)

	return nil
}
