package userRepository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	nftPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftPb"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/grpcConn"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/jwtAuth"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	UserRepositoryService interface {
		IsUserExists(pctx context.Context, email, username string) bool
		InsertUser(pctx context.Context, user *user.User) (primitive.ObjectID, error)
		FindOneUserProfile(pctx context.Context, userId string) (*user.UserProfileBson, error)

		AddToWallet(pctx context.Context, req *user.UserTransaction) error
		GetUserWalletAccount(pctx context.Context, userId string) (*user.UserWalletAccount, error)
		UpdateUserTransaction(pctx context.Context, orderId, paymentId string) error

		FindOneEmail(pctx context.Context, email string) (*user.User, error)
		FindOneUserProfileToRefresh(pctx context.Context, userId string) (*user.User, error)
		BlockOrUnblockUser(pctx context.Context, userId string, isActive bool) error

		// ----- Wish List -----
		AddToWishList(pctx context.Context, userId, nftId string) (any, error)
		GetWishList(pctx context.Context, userId string) (any, error)
		RemoveFromWishList(pctx context.Context, userId, nftId string) error

		// ----- Address -----
		AddAddress(pctx context.Context, userId string, req *user.CreateUserAddressReq) (*user.AddressModel, error)
		GetAddress(pctx context.Context, userId string) (*[]user.AddressModel, error)
		UpdateAddress(pctx context.Context, userId string, address_id string, req *user.CreateUserAddressReq) (*user.AddressModel, error)
		DeleteAddress(pctx context.Context, userId, addressId string) error
	}

	userRepository struct {
		db *mongo.Client
	}
)

func NewUserRepository(db *mongo.Client) UserRepositoryService {
	return &userRepository{db: db}
}

func (r *userRepository) userDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("user_db")
}

// logic to check if user exists
func (r *userRepository) IsUserExists(pctx context.Context, email, username string) bool {

	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	collection := db.Collection("users")

	//user := collection.FindOne(ctx, bson.M{"email": email, "username": username})

	user := new(user.User)
	if err := collection.FindOne(
		ctx,
		bson.M{
			"$or": []bson.M{
				{"email": email},
				{"username": username},
			},
		},
	).Decode(user); err != nil {
		log.Printf("Error: IsUserExists %v", err.Error())
		return true
	}

	return false
}

// logic to insert user
func (r *userRepository) InsertUser(pctx context.Context, user *user.User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	collection := db.Collection("users")

	newUser, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Error: InsertUser %v", err.Error())
		return primitive.NilObjectID, errors.New("Error: InsertUser " + err.Error())
	}

	return newUser.InsertedID.(primitive.ObjectID), nil
}

func (r *userRepository) FindOneUserProfile(pctx context.Context, userId string) (*user.UserProfileBson, error) {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("users")

	result := new(user.UserProfileBson)

	if err := col.FindOne(
		ctx,
		bson.M{"_id": utils.ConvertToObjectId(userId)},
		options.FindOne().SetProjection(
			bson.M{
				"_id":           1,
				"email":         1,
				"username":      1,
				"created_at":    1,
				"updated_at":    1,
				"is_blocked":    1,
				"wallet_amount": 1,
			},
		),
	).Decode(result); err != nil {
		log.Printf("Error: FindOneUserProfile: %s", err.Error())
		return nil, errors.New("error: user profile not found")
	}

	return result, nil
}

func (r *userRepository) AddToWallet(pctx context.Context, req *user.UserTransaction) error {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("user_transactions")
	// colUser := db.Collection("users")
	// userId := req.UserId
	// userIdTrim := strings.TrimPrefix(userId, "user:")

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: AddToWallet: %s", err.Error())
		return errors.New("error: AddToWallet transaction failed")
	}
	log.Printf("Result: AddToWallet: %v", result.InsertedID)

	// // Get user wallet account
	// userWallet, err := r.GetUserWalletAccount(ctx, userId)
	// // Update user wallet account in the database
	// _, err = colUser.UpdateOne(
	// 	ctx,
	// 	bson.M{"_id": utils.ConvertToObjectId(userIdTrim)},
	// 	bson.M{"$set": bson.M{"wallet_amount": userWallet.Balance}},
	// )
	// if err != nil {
	// 	log.Printf("Error: AddToWallet: %s", err.Error())
	// 	return errors.New("error: failed to update user wallet account")
	// }

	return nil
}

func (r *userRepository) GetUserWalletAccount(pctx context.Context, userId string) (*user.UserWalletAccount, error) {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("user_transactions")

	log.Println("userId : ", userId)

	filter := bson.A{
		bson.D{{"$match", bson.D{{"user_id", userId}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$user_id"},
					{"balance", bson.D{{"$sum", "$amount"}}},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"user_id", "$_id"},
					{"_id", 0},
					{"balance", 1},
				},
			},
		},
	}

	cursors, err := col.Aggregate(ctx, filter)
	if err != nil {
		log.Printf("Error: GetUserWalletAccount: %s", err.Error())
		return nil, errors.New("error: failed to get user wallet account")
	}

	result := new(user.UserWalletAccount)
	for cursors.Next(ctx) {
		if err := cursors.Decode(result); err != nil {
			log.Printf("Error: GetUserWalletAccount: %s", err.Error())
			return nil, errors.New("error: failed to get user wallet account")
		}
	}

	return result, nil
}

func (r *userRepository) FindOneEmail(pctx context.Context, email string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("users")

	result := new(user.User)

	if err := col.FindOne(ctx, bson.M{"email": email}).Decode(result); err != nil {
		log.Printf("Error: FindOneEmail: %s", err.Error())
		return nil, errors.New("error: email is invalid")
	}

	return result, nil
}

func (r *userRepository) FindOneUserProfileToRefresh(pctx context.Context, userId string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("users")

	result := new(user.User)

	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(userId)}).Decode(result); err != nil {
		log.Printf("Error: FindOneUserProfileToRefresh: %s", err.Error())
		return nil, errors.New("error: user profile not found")
	}

	return result, nil
}

func (r *userRepository) BlockOrUnblockUser(pctx context.Context, userId string, isActive bool) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("users")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(userId)}, bson.M{"$set": bson.M{"is_blocked": isActive}})
	if err != nil {
		log.Printf("Error: BlockOrUnblockUser failed: %s", err.Error())
		return errors.New("error: BlockOrUnblockUser failed")
	}
	log.Printf("BlockOrUnblockNft result: %v", result.ModifiedCount)

	return nil
}

// func (r *userRepository) FindAllUsers(pctx context.Context) ([]*user.UserProfileBson, error) {
// 	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
// 	defer cancel()

// 	db := r.userDbConn(ctx)
// 	col := db.Collection("users")

// 	result := new([]user.UserProfileBson)

// 	if err := col.Find().Decode(result); err != nil {
// 		log.Printf("Error: FindOneUserProfile: %s", err.Error())
// 		return nil, errors.New("error: user profile not found")
// 	}

// 	return result, nil
// }

// ------------ Wish List ------------ //
func (r *userRepository) AddToWishList(pctx context.Context, userId, nftId string) (any, error) {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("users")

	// Check if nft already in users wish list
	user := new(user.User)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(userId), "wishlist": utils.ConvertToObjectId(nftId)}).Decode(user); err == nil {
		return nil, errors.New("error: nft already exists in the wish list")
	}

	// fmt.Println("nftId: ", nftId)
	// fmt.Println("userId: ", userId)

	// Check if nft exists in the nft database

	// Set api key in context
	jwtAuth.SetApiKeyInContext(&ctx)
	// fmt.Println("ctx: ", ctx)

	// Connect to grpc server
	conn, err := grpcConn.NewGrpcClient("0.0.0.0:1623")
	if err != nil {
		log.Printf("failed to connect to grpc server: %s", err.Error())
		return nil, errors.New("failed to connect to grpc server")
	}
	// fmt.Println("conn: ", conn)

	req := &nftPb.FindNftsInIdsReq{
		Ids: []string{nftId},
	}
	// fmt.Println("req: ", req)

	result, err := conn.Nft().FindNftsInIds(ctx, req)

	fmt.Println("result: ", result)
	if err != nil {
		log.Printf("failed to find nfts in ids: %s", err.Error())
		return nil, errors.New("failed to find nfts in ids")
	}

	if result == nil {
		log.Printf("\nError: FindNftsInIds: nfts not found : %s", err.Error())
		return nil, errors.New("nft not found")
	}

	if len(result.Nfts) != 1 {
		log.Printf("\nError: FindNftsInIds: nfts not found : %s", err.Error())
		return nil, errors.New("nft not found")
	}

	// Add nft to wish list
	update := bson.M{"$addToSet": bson.M{"wishlist": utils.ConvertToObjectId(nftId)}}
	if _, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(userId)}, update); err != nil {
		log.Printf("Error: AddToWishList: %s", err.Error())
		return nil, errors.New("error: failed to add nft to wish list")
	}

	fmt.Println("result.Nfts: ", result.Nfts)

	// return current nft

	return result.Nfts, nil
}

func (r *userRepository) GetWishList(pctx context.Context, userId string) (any, error) {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("users")

	user := new(user.User)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(userId)}).Decode(user); err != nil {
		log.Printf("Error: GetWishList: %s", err.Error())
		return nil, errors.New("error: user not found")
	}

	// Check if wish list is empty
	if len(user.WishList) == 0 {
		return nil, errors.New("wish list is empty")
	}

	// fmt.Println("user: ", user)

	// Set api key in context
	jwtAuth.SetApiKeyInContext(&ctx)

	// Connect to grpc server
	conn, err := grpcConn.NewGrpcClient("0.0.0.0:1623")
	if err != nil {
		log.Printf("failed to connect to grpc server: %s", err.Error())
		return nil, errors.New("failed to connect to grpc server")
	}
	// fmt.Println("conn: ", conn)

	req := &nftPb.FindNftsInIdsReq{
		Ids: user.WishList,
	}
	// fmt.Println("req: ", req)

	result, err := conn.Nft().FindNftsInIds(ctx, req)

	fmt.Println("result: ", result)
	if err != nil {
		log.Printf("failed to find nfts in ids: %s", err.Error())
		return nil, errors.New("failed to find nfts in ids")
	}

	if result == nil {
		log.Printf("\nError: FindNftsInIds: nfts not found : %s", err.Error())
		return nil, errors.New("nft not found")
	}

	return result, nil
}

func (r *userRepository) RemoveFromWishList(pctx context.Context, userId, nftId string) error {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("users")

	// check if nft exists in the wish list
	user := new(user.User)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(userId), "wishlist": utils.ConvertToObjectId(nftId)}).Decode(user); err != nil {
		log.Printf("Error: RemoveFromWishList: %s", err.Error())
		return errors.New("error: nft not found in wish list")
	}

	// Remove nft from wish list
	update := bson.M{"$pull": bson.M{"wishlist": utils.ConvertToObjectId(nftId)}}
	if _, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(userId)}, update); err != nil {
		log.Printf("Error: RemoveFromWishList: %s", err.Error())
		return errors.New("error: failed to remove nft from wish list")
	}

	return nil
}

// ------------ Address ------------ //
func (r *userRepository) AddAddress(pctx context.Context, userId string, req *user.CreateUserAddressReq) (*user.AddressModel, error) {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("users")

	// user1 := new(user.User)

	// add address to address array in users collection and return the address id
	addressId := primitive.NewObjectID()
	update := bson.M{"$addToSet": bson.M{"addresses": bson.M{
		"id":         addressId,
		"name":       req.Name,
		"phone":      req.Phone,
		"street":     req.Street,
		"city":       req.City,
		"state":      req.State,
		"pincode":    req.Pincode,
		"country":    req.Country,
		"created_at": utils.LocalTime(),
		"updated_at": utils.LocalTime(),
	}}}

	if _, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(userId)}, update); err != nil {
		log.Printf("Error: AddAddress: %s", err.Error())
		return nil, errors.New("error: failed to add address")
	}

	// return address, nil
	return &user.AddressModel{
		Id:        addressId,
		Name:      req.Name,
		Phone:     req.Phone,
		Street:    req.Street,
		City:      req.City,
		State:     req.State,
		Pincode:   req.Pincode,
		Country:   req.Country,
		CreatedAt: utils.LocalTime(),
		UpdatedAt: utils.LocalTime(),
	}, nil
}

func (r *userRepository) GetAddress(pctx context.Context, userId string) (*[]user.AddressModel, error) {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("users")

	// get all address from address array from users collection
	user := new(user.User)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(userId)}).Decode(user); err != nil {
		log.Printf("Error: GetAddress: %s", err.Error())
		return nil, errors.New("error: user not found")
	}

	// check if address array is empty
	if len(user.Addressess) == 0 {
		return nil, errors.New("address is empty")
	}

	return &user.Addressess, nil

}

func (r *userRepository) UpdateAddress(pctx context.Context, userId string, addressId string, req *user.CreateUserAddressReq) (*user.AddressModel, error) {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("users")

	// check if address exists in the address array
	userDetails := new(user.User)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(userId), "addresses.id": utils.ConvertToObjectId(addressId)}).Decode(userDetails); err != nil {
		log.Printf("Error: UpdateAddress: %s", err.Error())
		return nil, errors.New("error: address not found")
	}

	// Update address in address array
	update := bson.M{"$set": bson.M{
		"addresses.$.name":       req.Name,
		"addresses.$.phone":      req.Phone,
		"addresses.$.street":     req.Street,
		"addresses.$.city":       req.City,
		"addresses.$.state":      req.State,
		"addresses.$.pincode":    req.Pincode,
		"addresses.$.country":    req.Country,
		"addresses.$.updated_at": utils.LocalTime(),
	}}

	if _, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(userId), "addresses.id": utils.ConvertToObjectId(addressId)}, update); err != nil {
		log.Printf("Error: UpdateAddress: %s", err.Error())
		return nil, errors.New("error: failed to update address")
	}

	// return updated address
	return &user.AddressModel{
		Id:        utils.ConvertToObjectId(addressId),
		Name:      req.Name,
		Phone:     req.Phone,
		Street:    req.Street,
		City:      req.City,
		State:     req.State,
		Pincode:   req.Pincode,
		Country:   req.Country,
		CreatedAt: userDetails.CreatedAt,
		UpdatedAt: utils.LocalTime(),
	}, nil

}

func (r *userRepository) DeleteAddress(pctx context.Context, userId, addressId string) error {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("users")

	// check if address exists in the address array
	user := new(user.User)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(userId), "addresses.id": utils.ConvertToObjectId(addressId)}).Decode(user); err != nil {
		log.Printf("Error: DeleteAddress: %s", err.Error())
		return errors.New("error: address not found")
	}

	// Remove address from address array
	update := bson.M{"$pull": bson.M{"addresses": bson.M{"id": utils.ConvertToObjectId(addressId)}}}
	if _, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(userId)}, update); err != nil {
		log.Printf("Error: DeleteAddress: %s", err.Error())
		return errors.New("error: failed to delete address")
	}

	return nil
}

// func (r *userRepository) CheckAddressExists(pctx context.Context, userId, addressId string) bool {
// 	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
// 	defer cancel()

// 	db := r.userDbConn(ctx)
// 	col := db.Collection("users")

// 	user := new(user.User)
// 	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(userId), "addresses.id": utils.ConvertToObjectId(addressId)}).Decode(user); err != nil {
// 		return false
// 	}

// 	return true
// }

func (r *userRepository) UpdateUserTransaction(pctx context.Context, orderId, paymentId string) error {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("user_transactions")

	// check if order exists in the user transactions
	user := new(user.UserTransaction)
	if err := col.FindOne(ctx, bson.M{"order_id": orderId}).Decode(user); err != nil {
		log.Printf("Error: UpdateUserTransaction: %s", err.Error())
		return errors.New("error: order not found")
	}

	// Update order success in user transactions
	update := bson.M{"$set": bson.M{"order_success": paymentId}}
	if _, err := col.UpdateOne(ctx, bson.M{"order_id": orderId}, update); err != nil {
		log.Printf("Error: UpdateUserTransaction: %s", err.Error())
		return errors.New("error: failed to update order success")
	}

	// add amount to user wallet
	_, err := r.GetUserWalletAccount(ctx, user.UserId)
	if err != nil {
		log.Printf("Error: UpdateUserTransaction: %s", err.Error())
		return errors.New("error: failed to get user wallet account")
	}

	return nil
}
