package nftUsecase

import (
	"context"
	"errors"
	"strings"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
)

// ------------- Bidding NFT Owner Usecase ------------- //

func (u *nftUsecase) FindManyBids(pctx context.Context, userId string) (any, error) {

	results, err := u.nftRepository.FindManyBids(pctx, userId)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (u *nftUsecase) CreateBid(pctx context.Context, req *nft.CreateBidReq, userId string) (any, error) {

	nftId := strings.TrimPrefix(req.NftId, "nft:")
	// check nft exist
	nft, err := u.nftRepository.FindOneNft(pctx, nftId)
	if err != nil {
		return nil, errors.New("error: nft not found")
	}

	// check current user is the owner of this nft
	if nft.OwnerId.Hex() != userId {
		return nil, errors.New("error: you are not the owner of this nft")
	}

	// check bid exist
	// bidResult, err := u.nftRepository.FindOneBid(pctx, nftId, userId)
	// if err != nil {
	// 	return nil, errors.New("error: bid not found")
	// }

	// if bidResult != nil {
	// 	return nil, errors.New("error: you have already bid this nft")
	// }

	bidId, err := u.nftRepository.CreateBid(pctx, userId, req)
	if err != nil {
		return nil, err
	}

	// return bid details with bid id
	return u.nftRepository.FindOneBid(pctx, bidId.Hex())

	// return "Bid created successfully" + bidId.Hex(), nil
}

// delete bid
func (u *nftUsecase) DeleteBid(pctx context.Context, bidId string, userId string) error {

	// check bid exist
	bid, err := u.nftRepository.FindOneBid(pctx, bidId)
	if err != nil {
		return errors.New("error: bid not found")
	}

	// check current user is the owner of this bid
	if bid.UserId.Hex() != userId {
		return errors.New("error: you are not the owner of this bid")
	}

	if bid.IsDeleted {
		return errors.New("error: bid is already soft deleted")
	}

	if err := u.nftRepository.DeleteBid(pctx, bidId); err != nil {
		return err
	}

	return nil
}

// edit bid
func (u *nftUsecase) EditBid(pctx context.Context, bidId string, userId string, req *nft.CreateBidReq) (*nft.Bid, error) {

	// check bid exist
	bid, err := u.nftRepository.FindOneBid(pctx, bidId)
	if err != nil {
		return nil, errors.New("error: bid not found")
	}

	// check current user is the owner of this bid
	if bid.UserId.Hex() != userId {
		return nil, errors.New("error: you are not the owner of this bid")
	}

	// Update logic
	updateReq := bson.M{}
	if req.Price >= 0 {
		updateReq["price"] = req.Price
	}
	updateReq["updated_at"] = utils.LocalTime()
	updateReq["expiry_date"] = req.ExpiryDate

	if err := u.nftRepository.EditBid(pctx, bidId, updateReq); err != nil {
		return nil, err
	}

	return u.nftRepository.FindOneBid(pctx, bidId)
}

// ---------- Bidding NFT User Usecase ---------- //

func (u *nftUsecase) FindUserBids(pctx context.Context, userId string) (any, error) {

	results, err := u.nftRepository.FindUserBids(pctx, userId)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (u *nftUsecase) BidNft(pctx context.Context, nftId, userId, price string) (any, error) {

	// check nft exist
	nft, err := u.nftRepository.FindOneNft(pctx, nftId)
	if err != nil {
		return nil, errors.New("error: nft not found")
	}

	// check current user is not the owner of this nft
	if nft.OwnerId.Hex() == userId {
		return nil, errors.New("error: you are the owner of this nft")
	}

	bidId, err := u.nftRepository.BidNft(pctx, userId, nftId, price)
	if err != nil {
		return nil, err
	}

	// check bid exist
	// bidResult, err := u.nftRepository.FindOneBidByNft(pctx, nftId)
	// if err != nil {
	// 	return nil, errors.New("error: bid not found")
	// }

	return "Bid created successfully. Bid Id : " + bidId.Hex(), nil
}

func (u *nftUsecase) WithdrawBid(pctx context.Context, bidId string, userId string) error {

	// check bid exist
	bid, err := u.nftRepository.FindOneUserBid(pctx, userId, bidId)
	// fmt.Println("bid: ", bid)
	if err != nil {
		return errors.New("error: bid not found")
	}

	// check current user is the owner of this bid
	if bid.UserId.Hex() != userId {
		return errors.New("error: you are not the owner of this bid")
	}

	if bid.IsDeleted {
		return errors.New("error: bid is already soft deleted")
	}

	if err := u.nftRepository.WithdrawBid(pctx, bidId); err != nil {
		return err
	}

	return nil
}
