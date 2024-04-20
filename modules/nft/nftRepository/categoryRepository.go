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
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	// only return if is_deleted is false
	//filter := bson.M{"_id": utils.ConvertToObjectId(categoryId), "is_deleted": false}
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(categoryId)}).Decode(result); err != nil {
		//if err := col.FindOne(ctx, filter).Decode(result); err != nil {
		log.Printf("Error: FindOneCategory failed: %s", err.Error())
		return nil, errors.New("error: category not found")
	}

	// if result.IsDeleted && !result.UsageStatus {
	// 	return nil, errors.New("error: category is deleted")
	// }

	// if !result.UsageStatus {
	// 	return nil, errors.New("error: category is blocked")
	// }

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
	colNft := db.Collection("nfts")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(categoryId)}, bson.M{"$set": bson.M{"usage_status": isActive}})
	if err != nil {
		log.Printf("Error: BlockOrUnblockCategory failed: %s", err.Error())
		return errors.New("error: BlockOrUnblockCategory failed")
	}

	// soft block/unblock all nfts in this category
	_, err = colNft.UpdateMany(ctx, bson.M{"category": utils.ConvertToObjectId(categoryId)}, bson.M{"$set": bson.M{"is_category_blocked": !isActive}})
	if err != nil {
		log.Printf("Error: BlockOrUnblockCategory failed: %s", err.Error())
		return errors.New("error: BlockOrUnblockCategory failed")
	}

	log.Printf("BlockOrUnblockNft result: %v", result.ModifiedCount)

	return nil
}

func (r *nftRepository) DeleteCategory(pctx context.Context, categoryId string) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("categories")
	colNft := db.Collection("nfts")

	_, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(categoryId)}, bson.M{"$set": bson.M{"is_deleted": true, "usage_status": false}})
	if err != nil {
		log.Printf("Error: DeleteCategory failed: %s", err.Error())
		return errors.New("error: DeleteCategory failed")
	}

	// soft delete all nfts in this category
	result, err := colNft.UpdateMany(ctx, bson.M{"category": utils.ConvertToObjectId(categoryId)}, bson.M{"$set": bson.M{"is_category_blocked": true}})
	if err != nil {
		log.Printf("Error: DeleteCategory failed: %s", err.Error())
		return errors.New("error: DeleteCategory failed")
	}
	log.Printf("DeleteCategory Success\n Count of Nfts Blocked: %v", result.ModifiedCount)

	return nil
}
