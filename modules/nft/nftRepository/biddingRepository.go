package nftRepository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/grpcConn"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/jwtAuth"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	userPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userPb"
)

// ------------------- Bidding User Side ------------------- //

// ------------------- Find One User Bid ------------------- //
func (r *nftRepository) FindOneUserBid(pctx context.Context, userId, bidId string) (*nft.SingleBid, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("user_bids")

	result := new(nft.SingleBid)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(bidId), "user_id": utils.ConvertToObjectId(userId)}).Decode(result); err != nil {
		log.Printf("Error: FindOneUserBid: %s", err.Error())
		return nil, errors.New("error: user bid not found")
	}

	return result, nil

}

func (r *nftRepository) BidNft(pctx context.Context, userId, nftId, price string) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	bidsCol := db.Collection("bids")
	userBidsCol := db.Collection("user_bids")

	// find bid id for this nft from bids collection
	bid := new(nft.Bid)
	if err := bidsCol.FindOne(ctx, bson.M{"nft_id": utils.ConvertToObjectId(nftId), "is_deleted": false}).Decode(bid); err != nil {
		log.Printf("Error: BidNft: %s", "this nft is not available for bidding")
		return primitive.NilObjectID, errors.New("error: this nft is not available for bidding")
	}

	// check this nft is already bidded by this user in user_bids collection using bid._id
	userBid := new(nft.SingleBid)
	if err := userBidsCol.FindOne(ctx, bson.M{"bid_id": bid.Id, "user_id": utils.ConvertToObjectId(userId), "is_deleted": false}).Decode(userBid); err == nil {
		log.Printf("Error: BidNft: %s", "this nft is already bidded by this user")
		return primitive.NilObjectID, errors.New("error: this nft is already bidded by this user")
	}

	// convert price string to float64
	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Printf("Error: BidNft: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: price conversion failed")
	}

	// check priceFloat is greater than  or equal to bid price
	if priceFloat < bid.Price {
		log.Printf("Error: BidNft: %s", "price must be greater than bid price")
		return primitive.NilObjectID, errors.New("error: price must be greater than bid price")
	}

	// check price is available in userWallet
	// userWallet := new(user.UserWalletAccount)

	// req := &userPb.GetUserWalletAccountReq{
	// 	UserId: userId,
	// }

	// jwtAuth.SetApiKeyInContext(&ctx)
	// conn, err := grpcConn.NewGrpcClient("0.0.0.0:1923")
	// if err != nil {
	// 	log.Printf("Error: gRPC connection failed: %s", err.Error())
	// 	return primitive.NilObjectID, errors.New("error: gRPC connection failed")
	// }

	// result, err := conn.User().GetUserWalletAccount(ctx, req)
	// if err != nil {
	// 	log.Printf("Error: GetUserWalletAccount failed: %s", err.Error())
	// 	return primitive.NilObjectID, errors.New("error: user wallet not found")
	// }

	// // fmt.Println("result: ", result)

	// if result.Balance < priceFloat {
	// 	log.Printf("Error: BidNft: %s", "insufficient balance")
	// 	return primitive.NilObjectID, errors.New("error: insufficient balance")
	// }

	// // deduct price from userWallet
	// deductReq := &userPb.DeductWalletAmountReq{
	// 	UserId: userId,
	// 	Amount: priceFloat,
	// }

	// deductResult, err := conn.User().DeductWalletAmount(ctx, deductReq)
	// if err != nil {
	// 	log.Printf("Error: DeductWalletAmount failed: %s", err.Error())
	// 	return primitive.NilObjectID, errors.New("error: deduct wallet amount failed")
	// }

	// fmt.Println("deductResult: ", deductResult)

	// insert user bid
	userBid = &nft.SingleBid{
		BidId:     bid.Id,
		UserId:    utils.ConvertToObjectId(userId),
		Price:     priceFloat,
		IsDeleted: false,
	}

	userBidId, err := userBidsCol.InsertOne(ctx, userBid)
	if err != nil {
		log.Printf("Error: BidNft: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: user bid failed")
	}

	// increment bid count in nft collection
	nftCol := db.Collection("nfts")
	_, err = nftCol.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(nftId)}, bson.M{"$inc": bson.M{"bid_count": 1}})
	if err != nil {
		log.Printf("Error: BidNft: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: increment bid count failed")
	}

	return userBidId.InsertedID.(primitive.ObjectID), nil

}

func (r *nftRepository) FindUserBids(pctx context.Context, userId string) (any, error) {

	// find the user bids from user_bids collection
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("user_bids")

	cursors, err := col.Find(ctx, bson.M{"user_id": utils.ConvertToObjectId(userId)})
	if err != nil {
		log.Printf("Error: FindUserBids: %s", err.Error())
		return nil, errors.New("error: find user bids failed")
	}

	results := make([]primitive.M, 0)
	for cursors.Next(ctx) {
		result := new(nft.SingleBid)
		if err := cursors.Decode(result); err != nil {
			log.Printf("Error: FindUserBids: %s", err.Error())
			return nil, errors.New("error: find user bids failed")
		}

		results = append(results, bson.M{
			"_id":    result.Id.Hex(),
			"bid_id": "bid:" + result.BidId.Hex(),
			// "user_id": "user:" + result.UserId.Hex(),
			"price":      result.Price,
			"is_deleted": result.IsDeleted,
		})
	}

	// if results is empty array
	if len(results) == 0 {
		return nil, errors.New("error: no bids found")
	}

	return results, nil
}

func (r *nftRepository) WithdrawBid(pctx context.Context, bidId string) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("user_bids")

	// check this user bid exist
	userBid := new(nft.SingleBid)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(bidId), "is_deleted": false}).Decode(userBid); err != nil {
		log.Printf("Error: WithdrawBid: %s", "user bid not found")
		return errors.New("error: user bid not found")
	}

	// check this user bid is already soft deleted
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(bidId), "is_deleted": true}).Decode(userBid); err == nil {
		log.Printf("Error: WithdrawBid: %s", "user bid is already soft deleted")
		return errors.New("error: user bid is already soft deleted")
	}

	// add the price to userWallet
	req := &userPb.AddWalletAmountReq{
		UserId: userBid.UserId.Hex(),
		Amount: userBid.Price,
	}

	jwtAuth.SetApiKeyInContext(&ctx)
	conn, err := grpcConn.NewGrpcClient("0.0.0.0:1923")

	if err != nil {
		log.Printf("Error: gRPC connection failed: %s", err.Error())
		return errors.New("error: gRPC connection failed")
	}

	addWalletRes, err := conn.User().AddWalletAmount(ctx, req)
	if err != nil {
		log.Printf("Error: AddWalletAmount failed: %s", err.Error())
		return errors.New("error: add wallet amount failed")
	}

	fmt.Println("addWalletRes: ", addWalletRes)

	// delete the user bid

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(bidId)}, bson.M{"$set": bson.M{"is_deleted": true}})
	if err != nil {
		log.Printf("Error: WithdrawBid failed: %s", err.Error())
		return errors.New("error: withdraw bid failed")
	}
	log.Printf("WithdrawBid result: %v", result.ModifiedCount)

	return nil
}

// ------------------- Bidding NFT Owner Side ------------------- //

// ------------------- Find One Bid ------------------- //
func (r *nftRepository) FindOneBid(pctx context.Context, bidId string) (*nft.Bid, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("bids")

	result := new(nft.Bid)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(bidId)}).Decode(result); err != nil {
		log.Printf("Error: FindOneBid failed: %s", err.Error())
		return nil, errors.New("error: bid not found")
	}

	return result, nil
}

// ------------------- Find Many Bids for owner ------------------- //
func (r *nftRepository) FindManyBids(pctx context.Context, userId string) ([]*nft.Bid, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("bids")

	fmt.Println("FindManyBids userId: ", userId)

	cursors, err := col.Find(ctx, bson.M{"user_id": utils.ConvertToObjectId(userId), "is_deleted": false})
	if err != nil {
		log.Printf("Error: FindManyBids: %s", err.Error())
		return make([]*nft.Bid, 0), errors.New("error: find many bids failed")
	}

	results := make([]*nft.Bid, 0)
	for cursors.Next(ctx) {
		result := new(nft.Bid)
		if err := cursors.Decode(result); err != nil {
			log.Printf("Error: FindManyBids: %s", err.Error())
			return make([]*nft.Bid, 0), errors.New("error: find many bids failed")
		}

		results = append(results, result)
	}

	// if results is empty array
	if len(results) == 0 {
		return nil, errors.New("error: no bids found")
	}

	return results, nil
}

// ------------------- Create Bid ------------------- //
func (r *nftRepository) CreateBid(pctx context.Context, userId string, req *nft.CreateBidReq) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("bids")

	// check this nft is already bidded by this user
	bid := new(nft.Bid)
	if err := col.FindOne(ctx, bson.M{"nft_id": utils.ConvertToObjectId(req.NftId), "user_id": utils.ConvertToObjectId(userId), "is_deleted": false}).Decode(bid); err == nil {
		log.Printf("Error: CreateBid: %s", "this nft is already bidded by this user")
		return primitive.NilObjectID, errors.New("error: this nft is already bidded by this user")
	}

	bid = &nft.Bid{
		NftId:      utils.ConvertToObjectId(req.NftId),
		UserId:     utils.ConvertToObjectId(userId),
		Price:      req.Price,
		ExpiryDate: req.ExpiryDate,
		IsDeleted:  false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	bidId, err := col.InsertOne(ctx, bid)

	if err != nil {
		log.Printf("Error: CreateBid: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: create bid failed")
	}

	return bidId.InsertedID.(primitive.ObjectID), nil
}

func (r *nftRepository) DeleteBid(pctx context.Context, bidId string) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("bids")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(bidId)}, bson.M{"$set": bson.M{"is_deleted": true}})
	if err != nil {
		log.Printf("Error: DeleteBid failed: %s", err.Error())
		return errors.New("error: delete bid failed")
	}
	log.Printf("DeleteBid result: %v", result.ModifiedCount)

	return nil
}

func (r *nftRepository) EditBid(pctx context.Context, bidId string, req primitive.M) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	col := db.Collection("bids")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(bidId)}, bson.M{"$set": req})
	if err != nil {
		log.Printf("Error: EditBid failed: %s", err.Error())
		return errors.New("error: edit bid failed")
	}
	log.Printf("EditBid result: %v", result.ModifiedCount)

	return nil
}

// ------------------- Admin ------------------- //
func (r *nftRepository) ExecuteBids(pctx context.Context) (any, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.nftDbConn(ctx)
	bidsColl := db.Collection("bids")
	user_bids := db.Collection("user_bids")
	nftCollection := db.Collection("nfts")

	// find all bids that are expired and not deleted
	cursors, err := bidsColl.Find(ctx, bson.M{"expiry_date": bson.M{"$lte": time.Now()}, "is_deleted": false})
	if err != nil {
		log.Printf("Error: ExecuteBids: %s", err.Error())
		return nil, errors.New("error: execute bids failed")
	}

	// if results is empty array
	if cursors.RemainingBatchLength() == 0 {
		return nil, errors.New("error: no bids found")
	}

	results := make([]primitive.M, 0)

	// iterate through results array
	for cursors.Next(ctx) {

		currentBid := new(nft.Bid)
		if err := cursors.Decode(currentBid); err != nil {
			log.Printf("Error: ExecuteBids: %s", err.Error())
			return nil, errors.New("error: execute bids failed")
		}

		fmt.Println("currentBid: ", currentBid.Id.Hex())

		// for each bid in results find all bids that are done by users in user_bids collection
		opts := options.Find().SetSort(bson.D{{"price", -1}})
		userBids, err := user_bids.Find(ctx, bson.M{"bid_id": currentBid.Id, "is_deleted": false}, opts)
		if err != nil {
			log.Printf("Error: ExecuteBids: %s", err.Error())
			return nil, errors.New("error: execute bids failed")
		}

		fmt.Println("userBids: ", userBids.RemainingBatchLength())

		if userBids.RemainingBatchLength() == 0 {
			// mark bid as done
			_, err := bidsColl.UpdateOne(ctx, bson.M{"_id": currentBid.Id}, bson.M{"$set": bson.M{"is_deleted": true}})
			if err != nil {
				log.Printf("Error: No one bidded this bid: ExecuteBids: %s", err.Error())
				//return nil, errors.New("error: execute bids failed since no one bids this bid")
			}
			results = append(results, bson.M{
				"bid_id":      currentBid.Id.Hex(),
				"nft_id":      currentBid.NftId.Hex(),
				"user_id":     currentBid.UserId.Hex(),
				"floor_price": currentBid.Price,
				"status":      "bid cancelled since no one bidded this bid",
			})
			continue
		}
		flag := false
		for userBids.Next(ctx) {
			result := new(nft.SingleBid)
			if err := userBids.Decode(result); err != nil {
				log.Printf("Error: ExecuteBids: %s", err.Error())
				return nil, errors.New("error: execute bids failed")
			}

			// results = append(results, bson.M{
			// 	"_id":        result.Id.Hex(),
			// 	"bid_id":     "bid:" + result.BidId.Hex(),
			// 	"user_id":    "user:" + result.UserId.Hex(),
			// 	"price":      result.Price,
			// 	"is_deleted": result.IsDeleted,
			// })

			// check price is available in userWallet
			// userId trimmed user:
			userId := strings.TrimPrefix(result.UserId.Hex(), "user:")

			req := &userPb.GetUserWalletAccountReq{
				UserId: userId,
			}

			jwtAuth.SetApiKeyInContext(&ctx)
			conn, err := grpcConn.NewGrpcClient("0.0.0.0:1923")
			if err != nil {
				log.Printf("Error: gRPC connection failed: %s", err.Error())
				return primitive.NilObjectID, errors.New("error: gRPC connection failed")
			}

			gRPCresult, err := conn.User().GetUserWalletAccount(ctx, req)
			if err != nil {
				log.Printf("Error: GetUserWalletAccount failed: %s", err.Error())
				return primitive.NilObjectID, errors.New("error: user wallet not found")
			}

			fmt.Println("result: ", result)

			if gRPCresult.Balance < result.Price {
				log.Printf("Error: BidNft: %s", "insufficient balance")
				continue
				// return primitive.NilObjectID, errors.New("error: insufficient balance")
			}

			// deduct price from userWallet
			deductReq := &userPb.DeductWalletAmountReq{
				UserId: userId,
				Amount: result.Price,
			}

			deductResult, err := conn.User().DeductWalletAmount(ctx, deductReq)
			if err != nil {
				log.Printf("Error: DeductWalletAmount failed: %s", err.Error())
				return primitive.NilObjectID, errors.New("error: deduct wallet amount failed")
			}

			fmt.Println("deductResult: ", deductResult)

			if deductResult.Balance == gRPCresult.Balance-result.Price {
				// mark bid as done
				updateResult, err := bidsColl.UpdateOne(ctx, bson.M{"_id": currentBid.Id}, bson.M{"$set": bson.M{"is_deleted": true}})
				if err != nil {
					log.Printf("Error: ExecuteBids: %s", err.Error())
					return primitive.NilObjectID, errors.New("error: execute bids failed")
				}
				log.Printf("ExecuteBids result 2: %v", updateResult.ModifiedCount)

				// change nft ownership
				nftUpdateResult, err := nftCollection.UpdateOne(ctx, bson.M{"_id": currentBid.NftId}, bson.M{"$set": bson.M{"owner_id": result.UserId}})
				if err != nil {
					log.Printf("Error: ExecuteBids: %s", err.Error())
					return primitive.NilObjectID, errors.New("error: execute bids failed")
				}
				log.Printf("ExecuteBids result 3: %v", nftUpdateResult.ModifiedCount)

				// send amount to current owner by deducting 2.5 %
				sendReq := &userPb.AddWalletAmountReq{
					UserId: string(currentBid.UserId.Hex()),
					Amount: result.Price * 0.975,
				}

				sendResult, err := conn.User().AddWalletAmount(ctx, sendReq)
				if err != nil {
					log.Printf("Error: AddWalletAmount failed: %s", err.Error())
					return primitive.NilObjectID, errors.New("error: add wallet amount failed")
				}

				fmt.Println("sendResult: ", sendResult)

				results = append(results, bson.M{
					"bid_id":      currentBid.Id.Hex(),
					"nft_id":      currentBid.NftId.Hex(),
					"user_id":     currentBid.UserId.Hex(),
					"floor_price": currentBid.Price,
					"bid_rate":    result.Price,
					"bid_user_id": result.UserId.Hex(),
				})
				flag = true
				break
			}

		}

		if !flag {
			// mark bid as done
			updateResult, err := bidsColl.UpdateOne(ctx, bson.M{"_id": currentBid.Id}, bson.M{"$set": bson.M{"is_deleted": true}})
			if err != nil {
				log.Printf("Error: ExecuteBids: %s", err.Error())
				return primitive.NilObjectID, errors.New("error: execute bids failed")
			}
			log.Printf("ExecuteBids result 2: %v", updateResult.ModifiedCount)
			results = append(results, bson.M{
				"bid_id":      currentBid.Id.Hex(),
				"nft_id":      currentBid.NftId.Hex(),
				"user_id":     currentBid.UserId.Hex(),
				"floor_price": currentBid.Price,
				"status":      "no user have wallet balance",
			})
		}
	}

	return results, nil
}
