package userRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user"
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
		FindOneEmail(pctx context.Context, email string) (*user.User, error)
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
				"_id":        1,
				"email":      1,
				"username":   1,
				"created_at": 1,
				"updated_at": 1,
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

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: AddToWallet: %s", err.Error())
		return errors.New("error: AddToWallet transcation failed")
	}
	log.Printf("Result: AddToWallet: %v", result.InsertedID)

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
