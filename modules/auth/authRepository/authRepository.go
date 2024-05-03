package authRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/auth"
	userPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userPb"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/grpcConn"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/jwtAuth"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	AuthRepositoryService interface {
		CredentialSearch(pctx context.Context, grpcUrl string, req *userPb.CredentialSearchReq) (*userPb.UserProfile, error)
		InsertOneUserCredential(pctx context.Context, req *auth.Credential) (primitive.ObjectID, error)
		FindOneUserCredential(pctx context.Context, credentialId string) (*auth.Credential, error)
		FindOneUserProfileToRefresh(pctx context.Context, grpcUrl string, req *userPb.FindOneUserProfileToRefreshReq) (*userPb.UserProfile, error)
		UpdateOneUserCredential(pctx context.Context, credentialId string, req *auth.UpdateRefreshTokenReq) error
		DeleteOneUserCredential(pctx context.Context, credentialId string) (int64, error)
		FindOneAccessToken(pctx context.Context, accessToken string) (*auth.Credential, error)
		RolesCount(pctx context.Context) (int64, error)
		InsertOneOtp(pctx context.Context, req *auth.OtpRequestReq, otp string) error
		//OtpVerification(pctx context.Context, req *auth.OtpVerificationReq) (*userPb.UserProfile, error)
	}

	authRepository struct {
		db *mongo.Client
	}
)

func NewAuthRepository(db *mongo.Client) AuthRepositoryService {
	return &authRepository{db}
}

func (r *authRepository) authDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("auth_db")
}

func (r *authRepository) CredentialSearch(pctx context.Context, grpcUrl string, req *userPb.CredentialSearchReq) (*userPb.UserProfile, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	jwtAuth.SetApiKeyInContext(&ctx)
	conn, err := grpcConn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}

	result, err := conn.User().CredentialSearch(ctx, req)
	if err != nil {
		log.Printf("Error: CredentialSearch failed: %s", err.Error())
		return nil, errors.New("error : Invalid Credentials")
	}

	return result, nil
}

func (r *authRepository) InsertOneUserCredential(pctx context.Context, req *auth.Credential) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	coll := db.Collection("auth")

	result, err := coll.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOneUserCredential failed: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one user credential failed")
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *authRepository) FindOneUserCredential(pctx context.Context, credentialId string) (*auth.Credential, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	result := new(auth.Credential)

	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(credentialId)}).Decode(result); err != nil {
		log.Printf("Error: FindOneUserCredential failed: %s", err.Error())
		return nil, errors.New("error: find one user credential failed")
	}

	return result, nil
}

func (r *authRepository) FindOneUserProfileToRefresh(pctx context.Context, grpcUrl string, req *userPb.FindOneUserProfileToRefreshReq) (*userPb.UserProfile, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	jwtAuth.SetApiKeyInContext(&ctx)
	conn, err := grpcConn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}

	result, err := conn.User().FindOneUserProfileToRefresh(ctx, req)
	if err != nil {
		log.Printf("Error: FindOneUserProfileToRefresh failed: %s", err.Error())
		return nil, errors.New("error: user profile not found")
	}

	return result, nil
}

func (r *authRepository) UpdateOneUserCredential(pctx context.Context, credentialId string, req *auth.UpdateRefreshTokenReq) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	_, err := col.UpdateOne(
		ctx,
		bson.M{"_id": utils.ConvertToObjectId(credentialId)},
		bson.M{
			"$set": bson.M{
				"user_id":       req.UserId,
				"access_token":  req.AccessToken,
				"refresh_token": req.RefreshToken,
				"updated_at":    req.UpdatedAt,
			},
		},
	)
	if err != nil {
		log.Printf("Error: UpdateOneUserCredential failed: %s", err.Error())
		return errors.New("error: user credential not found")
	}

	return nil
}

func (r *authRepository) DeleteOneUserCredential(pctx context.Context, credentialId string) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	result, err := col.DeleteOne(ctx, bson.M{"_id": utils.ConvertToObjectId(credentialId)})
	if err != nil {
		log.Printf("Error: DeleteOneUserCredential failed: %s", err.Error())
		return -1, errors.New("error: delete user credential failed")
	}
	log.Printf("DeleteOneUserCredential result: %v", result)

	return result.DeletedCount, nil
}

func (r *authRepository) FindOneAccessToken(pctx context.Context, accessToken string) (*auth.Credential, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	credential := new(auth.Credential)
	if err := col.FindOne(ctx, bson.M{"access_token": accessToken}).Decode(credential); err != nil {
		log.Printf("Error: FindOneAccessToken failed: %s", err.Error())
		return nil, errors.New("error: access token not found")
	}

	return credential, nil
}

func (r *authRepository) RolesCount(pctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("roles")

	count, err := col.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Printf("Error: RolesCount failed: %s", err.Error())
		return -1, errors.New("error: roles count failed")
	}

	return count, nil
}

func (r *authRepository) InsertOneOtp(pctx context.Context, req *auth.OtpRequestReq, otp string) error {
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("otp")

	// Check if the email exists in the user database
	userCol := db.Collection("users")
	user := new(auth.UserLoginReq)
	if err := userCol.FindOne(ctx, bson.M{"email": req.Email}).Decode(user); err != nil {
		log.Printf("Error: InsertOneOtp failed: %s", err.Error())
		return errors.New("error: email does not exist")
	}

	// Save the OTP in the "otp" collection
	otpDoc := bson.M{
		"email": req.Email,
		"otp":   otp,
	}
	_, err := col.InsertOne(ctx, otpDoc)
	if err != nil {
		log.Printf("Error: InsertOneOtp failed: %s", err.Error())
		return errors.New("error: failed to save OTP")
	}

	return nil
}

// func (r *authRepository) OtpVerification(pctx context.Context, req *auth.OtpVerificationReq) (*userPb.UserProfile, error) {
// 	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
// 	defer cancel()

// 	db := r.authDbConn(ctx)
// 	col := db.Collection("otp")

// 	otpDoc := new(auth.Otp)
// 	if err := col.FindOne(ctx, bson.M{"email": req.Email, "otp": req.Otp}).Decode(otpDoc); err != nil {
// 		log.Printf("Error: OtpVerification failed: %s", err.Error())
// 		return nil, errors.New("error: invalid OTP")
// 	}

// 	// Delete the OTP from the "otp" collection
// 	_, err := col.DeleteOne(ctx, bson.M{"email": req.Email})
// 	if err != nil {
// 		log.Printf("Error: OtpVerification failed: %s", err.Error())
// 		return nil, errors.New("error: failed to delete OTP")
// 	}

// 	// Check if the email exists in the user database
// 	userCol := db.Collection("users")
// 	user := new(auth.User)
// 	if err := userCol.FindOne(ctx, bson.M{"email": req.Email}).Decode(user); err != nil {
// 		log.Printf("Error: OtpVerification failed: %s", err.Error())
// 		return nil, errors.New("error: email does not exist")
// 	}

// 	// Generate a new access token
// 	accessToken, err := jwtAuth.GenerateAccessToken(user.Id)

// 	return userProfile, nil
// }
