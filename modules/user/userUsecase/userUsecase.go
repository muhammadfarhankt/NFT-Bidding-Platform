package userUsecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user"
	userPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userPb"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userRepository"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type (
	UserUsecaseService interface {
		InsertUser(pctx context.Context, req *user.CreateUserReq) (string, error)
		FindOneUserProfile(pctx context.Context, userId string) (*user.UserProfile, error)
		AddToWallet(pctx context.Context, req *user.CreateUserTransactionReq) (*user.UserWalletAccount, error)
		GetUserWalletAccount(pctx context.Context, userId string) (*user.UserWalletAccount, error)
		FindOneEmail(pctx context.Context, password, email string) (*userPb.UserProfile, error)
		FindOneUserProfileToRefresh(pctx context.Context, userId string) (*userPb.UserProfile, error)
		BlockOrUnblockUser(pctx context.Context, userId string) (bool, error)
	}

	userUsecase struct {
		userRepository userRepository.UserRepositoryService
	}
)

func NewUserUsecase(userRepository userRepository.UserRepositoryService) UserUsecaseService {
	return &userUsecase{userRepository: userRepository}
}

func (u *userUsecase) InsertUser(pctx context.Context, req *user.CreateUserReq) (string, error) {

	if req.Email == "" || req.Password == "" || req.Username == "" {
		return "", errors.New("error: Email, Password Or Username cannot be empty")
	}

	if !u.userRepository.IsUserExists(pctx, req.Email, req.Username) {
		return "", errors.New("error: User already exists")
	}

	// hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 7)
	if err != nil {
		return "", errors.New("error: Failed to hash password")
	}

	// create user
	userId, err := u.userRepository.InsertUser(pctx, &user.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		Username:  req.Username,
		CreatedAt: utils.LocalTime(),
		UpdatedAt: utils.LocalTime(),
		UserRoles: []user.UserRole{
			{
				RoleTitle: "user",
				RoleCode:  0,
			},
		},
		IsBlocked: false,
	})

	if err != nil {
		return "", errors.New("error: Failed to insert user")
	}

	return userId.Hex(), nil
}

func (u *userUsecase) FindOneUserProfile(pctx context.Context, userId string) (*user.UserProfile, error) {

	result, err := u.userRepository.FindOneUserProfile(pctx, userId)
	if err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("Asia/Calcutta")

	return &user.UserProfile{
		Id:        result.Id.Hex(),
		Email:     result.Email,
		Username:  result.Username,
		CreatedAt: result.CreatedAt.In(loc),
		UpdatedAt: result.UpdatedAt.In(loc),
		IsBlocked: result.IsBlocked,
	}, nil
}

func (u *userUsecase) AddToWallet(pctx context.Context, req *user.CreateUserTransactionReq) (*user.UserWalletAccount, error) {

	log.Println("req", req)
	if err := u.userRepository.AddToWallet(pctx, &user.UserTransaction{
		UserId:    req.UserId,
		Amount:    req.Amount,
		CreatedAt: utils.LocalTime(),
	}); err != nil {
		return nil, err
	}

	// Get user saving account
	return u.userRepository.GetUserWalletAccount(pctx, req.UserId)
}

func (u *userUsecase) GetUserWalletAccount(pctx context.Context, userId string) (*user.UserWalletAccount, error) {
	return u.userRepository.GetUserWalletAccount(pctx, userId)
}

func (u *userUsecase) FindOneEmail(pctx context.Context, password, email string) (*userPb.UserProfile, error) {
	result, err := u.userRepository.FindOneEmail(pctx, email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		log.Printf("Error: FindOneEmail: %s", err.Error())
		return nil, errors.New("error: password is invalid")
	}

	roleCode := 0
	for _, role := range result.UserRoles {
		roleCode += role.RoleCode
	}

	loc, _ := time.LoadLocation("Asia/Calcutta")

	return &userPb.UserProfile{
		Id:        result.Id.Hex(),
		Email:     result.Email,
		Username:  result.Username,
		RoleCode:  int32(roleCode),
		CreatedAt: result.CreatedAt.In(loc).String(),
		UpdatedAt: result.UpdatedAt.In(loc).String(),
	}, nil
}

func (u *userUsecase) FindOneUserProfileToRefresh(pctx context.Context, userId string) (*userPb.UserProfile, error) {
	result, err := u.userRepository.FindOneUserProfileToRefresh(pctx, userId)
	if err != nil {
		return nil, err
	}

	roleCode := 0
	for _, v := range result.UserRoles {
		roleCode += v.RoleCode
	}

	loc, _ := time.LoadLocation("Asia/Calcutta")

	return &userPb.UserProfile{
		Id:        result.Id.Hex(),
		Email:     result.Email,
		Username:  result.Username,
		RoleCode:  int32(roleCode),
		CreatedAt: result.CreatedAt.In(loc).String(),
		UpdatedAt: result.UpdatedAt.In(loc).String(),
	}, nil
}

func (u *userUsecase) BlockOrUnblockUser(pctx context.Context, userId string) (bool, error) {
	result, err := u.userRepository.FindOneUserProfile(pctx, userId)
	if err != nil {
		return false, err
	}

	if err := u.userRepository.BlockOrUnblockUser(pctx, userId, !result.IsBlocked); err != nil {
		return false, err
	}

	return !result.IsBlocked, nil
}

// func (u *userUsecase) FindAllUsers(pctx context.Context) (*[]user.UserProfile, error) {

// 	result, err := u.userRepository.FindOneUserProfile(pctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	loc, _ := time.LoadLocation("Asia/Calcutta")

// 	return &[]user.UserProfile{
// 		Id:        result.Id.Hex(),
// 		Email:     result.Email,
// 		Username:  result.Username,
// 		RoleCode:  int32(roleCode),
// 		CreatedAt: result.CreatedAt.In(loc).String(),
// 		UpdatedAt: result.UpdatedAt.In(loc).String(),
// 	}, nil
// }
