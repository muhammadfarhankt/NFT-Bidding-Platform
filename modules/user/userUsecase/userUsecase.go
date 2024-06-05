package userUsecase

import (
	"context"
	"errors"
	"fmt"
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
		ResetPassword(pctx context.Context, userId string, req *user.ResetPasswordReq) error

		AddToWallet(pctx context.Context, req *user.CreateUserTransactionReq, orderId string) (*user.UserWalletAccount, error)
		GetUserWalletAccount(pctx context.Context, userId string) (*userPb.GetUserWalletAccountRes, error)
		UpdateUserTransaction(pctx context.Context, orderId, paymentId string) error

		FindOneEmail(pctx context.Context, password, email string) (*userPb.UserProfile, error)
		FindOneUserProfileToRefresh(pctx context.Context, userId string) (*userPb.UserProfile, error)
		FindOneUserOnEmail(pctx context.Context, email string) (*userPb.UserProfile, error)

		// ----- Reports -----
		UserPaymentReport(pctx context.Context, userId string) (any, error)
		SingleOrderPaymentReport(pctx context.Context, userId string) (any, error)

		// --- admin ----
		BlockOrUnblockUser(pctx context.Context, userId string) (bool, error)
		SalesReport(pctx context.Context) (any, error)

		// Wish List
		AddToWishList(pctx context.Context, userId, nftId string) (any, error)
		GetWishList(pctx context.Context, userId string) (any, error)
		RemoveFromWishList(pctx context.Context, userId, nftId string) error

		// Address
		AddAddress(pctx context.Context, userId string, req *user.CreateUserAddressReq) (*user.AddressModel, error)
		GetAddress(pctx context.Context, userId string) (*[]user.AddressModel, error)
		UpdateAddress(pctx context.Context, userId string, address_id string, req *user.CreateUserAddressReq) (*user.AddressModel, error)
		DeleteAddress(pctx context.Context, userId, addressId string) error

		// Bidding wallet amount deduction and refund
		AddWalletAmount(pctx context.Context, userId string, amount float64) (*userPb.GetUserWalletAccountRes, error)
		DeductWalletAmount(pctx context.Context, userId string, amount float64) (*userPb.GetUserWalletAccountRes, error)
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

	fmt.Println("result findOneUserProfile : \n", result)

	return &user.UserProfile{
		Id:           result.Id.Hex(),
		Email:        result.Email,
		Username:     result.Username,
		WalletAmount: result.WalletAmount,
		CreatedAt:    result.CreatedAt.In(loc),
		UpdatedAt:    result.UpdatedAt.In(loc),
		IsBlocked:    result.IsBlocked,
	}, nil
}

func (u *userUsecase) AddToWallet(pctx context.Context, req *user.CreateUserTransactionReq, orderId string) (*user.UserWalletAccount, error) {

	log.Println("req", req)
	if err := u.userRepository.AddToWallet(pctx, &user.UserTransaction{
		UserId:    req.UserId,
		Amount:    req.Amount,
		OrderId:   orderId,
		AddressId: req.AddressId,
		CreatedAt: utils.LocalTime(),
		UpdatedAt: utils.LocalTime(),
	}); err != nil {
		return nil, err
	}

	// Get user saving account
	return u.userRepository.GetUserWalletAccount(pctx, req.UserId)
}

func (u *userUsecase) GetUserWalletAccount(pctx context.Context, userId string) (*userPb.GetUserWalletAccountRes, error) {

	result, err := u.userRepository.GetUserWalletAccount(pctx, userId)
	if err != nil {
		return nil, err
	}

	return &userPb.GetUserWalletAccountRes{
		Balance: result.Balance,
		UserId:  result.UserId,
	}, nil

	// return u.userRepository.GetUserWalletAccount(pctx, userId)
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

	fmt.Println("result findOneEmail \n", result)

	return &userPb.UserProfile{
		Id:           result.Id.Hex(),
		Email:        result.Email,
		Username:     result.Username,
		WalletAmount: result.WalletAmount,
		RoleCode:     int32(roleCode),
		CreatedAt:    result.CreatedAt.In(loc).String(),
		UpdatedAt:    result.UpdatedAt.In(loc).String(),
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

// ---------------- Wish List ---------------- //
func (u *userUsecase) AddToWishList(pctx context.Context, userId, nftId string) (any, error) {
	return u.userRepository.AddToWishList(pctx, userId, nftId)
}

func (u *userUsecase) GetWishList(pctx context.Context, userId string) (any, error) {
	return u.userRepository.GetWishList(pctx, userId)
	// return nil, nil
}

func (u *userUsecase) RemoveFromWishList(pctx context.Context, userId, nftId string) error {
	return u.userRepository.RemoveFromWishList(pctx, userId, nftId)
	// return nil
}

// ---------------- Address ---------------- //
func (u *userUsecase) AddAddress(pctx context.Context, userId string, req *user.CreateUserAddressReq) (*user.AddressModel, error) {
	errorMessage := ""
	if req.Name == "" {
		errorMessage += "Name cannot be empty.    "
	}
	if req.Phone == "" {
		errorMessage += "Phone cannot be empty.    "
	} else if len(req.Phone) != 10 {
		errorMessage += "Phone must be 10 digits.    "
	}
	if req.Pincode == "" {
		errorMessage += "Pincode cannot be empty.    "
	} else if len(req.Pincode) != 6 {
		errorMessage += "Pincode must be 6 digits.    "
	}
	if req.Street == "" {
		errorMessage += "Street cannot be empty.    "
	}
	if req.City == "" {
		errorMessage += "City cannot be empty.    "
	}
	if req.State == "" {
		errorMessage += "State cannot be empty.    "
	}

	if errorMessage != "" {
		return nil, errors.New(errorMessage)
	}

	return u.userRepository.AddAddress(pctx, userId, req)
}

func (u *userUsecase) GetAddress(pctx context.Context, userId string) (*[]user.AddressModel, error) {
	return u.userRepository.GetAddress(pctx, userId)
}

func (u *userUsecase) UpdateAddress(pctx context.Context, userId string, address_id string, req *user.CreateUserAddressReq) (*user.AddressModel, error) {
	// fmt.Println("address_id: ", address_id)

	errorMessage := ""
	if req.Name == "" {
		errorMessage += "Name cannot be empty.    "
	}
	if req.Phone == "" {
		errorMessage += "Phone cannot be empty.    "
	} else if len(req.Phone) != 10 {
		errorMessage += "Phone must be 10 digits.    "
	}
	if req.Pincode == "" {
		errorMessage += "Pincode cannot be empty.    "
	} else if len(req.Pincode) != 6 {
		errorMessage += "Pincode must be 6 digits.    "
	}
	if req.Street == "" {
		errorMessage += "Street cannot be empty.    "
	}
	if req.City == "" {
		errorMessage += "City cannot be empty.    "
	}
	if req.State == "" {
		errorMessage += "State cannot be empty.    "
	}

	if errorMessage != "" {
		return nil, errors.New(errorMessage)
	}

	return u.userRepository.UpdateAddress(pctx, userId, address_id, req)
}

func (u *userUsecase) DeleteAddress(pctx context.Context, userId, addressId string) error {
	return u.userRepository.DeleteAddress(pctx, userId, addressId)
}

func (u *userUsecase) UpdateUserTransaction(pctx context.Context, orderId, paymentId string) error {
	return u.userRepository.UpdateUserTransaction(pctx, orderId, paymentId)
}

func (u *userUsecase) FindOneUserOnEmail(pctx context.Context, email string) (*userPb.UserProfile, error) {
	result, err := u.userRepository.FindOneEmail(pctx, email)
	if err != nil {
		return nil, err
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

func (u *userUsecase) AddWalletAmount(pctx context.Context, userId string, amount float64) (*userPb.GetUserWalletAccountRes, error) {

	result, err := u.userRepository.AddWalletAmount(pctx, userId, amount)
	if err != nil {
		return nil, err
	}

	return &userPb.GetUserWalletAccountRes{
		Balance: result.Balance,
		UserId:  result.UserId,
	}, nil
}

func (u *userUsecase) DeductWalletAmount(pctx context.Context, userId string, amount float64) (*userPb.GetUserWalletAccountRes, error) {

	result, err := u.userRepository.DeductWalletAmount(pctx, userId, amount)
	if err != nil {
		return nil, err
	}

	return &userPb.GetUserWalletAccountRes{
		Balance: result.Balance,
		UserId:  result.UserId,
	}, nil

}

func (u *userUsecase) SalesReport(pctx context.Context) (any, error) {
	return u.userRepository.SalesReport(pctx)
}

func (u *userUsecase) ResetPassword(pctx context.Context, userId string, req *user.ResetPasswordReq) error {

	// hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 7)
	if err != nil {
		return errors.New("error: Failed to hash password")
	}

	if err := u.userRepository.ResetPassword(pctx, userId, req.OldPassword, string(hashedPassword)); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) UserPaymentReport(pctx context.Context, userId string) (any, error) {

	return u.userRepository.UserPaymentReport(pctx, userId)

}

func (u *userUsecase) SingleOrderPaymentReport(pctx context.Context, userId string) (any, error) {

	return u.userRepository.SingleOrderPaymentReport(pctx, userId)

}
