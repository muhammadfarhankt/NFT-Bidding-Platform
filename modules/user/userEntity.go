package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	User struct {
		Id           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Email        string             `json:"email" bson:"email"`
		Password     string             `json:"password" bson:"password"`
		Username     string             `json:"username" bson:"username"`
		WalletAmount float64            `json:"wallet_amount" bson:"wallet_amount"`
		OTP          string             `json:"otp" bson:"otp"`
		OtpExpiredAt time.Time          `json:"otp_expired_at" bson:"otp_expired_at"`
		CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
		UserRoles    []UserRole         `bson:"user_roles"`
		IsBlocked    bool               `json:"is_blocked" bson:"is_blocked"`
		WishList     []string           `json:"wishlist" bson:"wishlist"`
		Addressess   []AddressModel     `json:"addressess" bson:"addressess"`
	}

	UserRole struct {
		RoleTitle string `json:"role_title" bson:"role_title"`
		RoleCode  int    `json:"role_code" bson:"role_code"`
	}

	UserProfileBson struct {
		Id           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Email        string             `json:"email" bson:"email"`
		Username     string             `json:"username" bson:"username"`
		WalletAmount float64            `json:"wallet_amount" bson:"wallet_amount"`
		CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
		IsBlocked    bool               `json:"is_blocked" bson:"is_blocked"`
	}

	UserWalletAccount struct {
		UserId  string  `json:"user_id" bson:"user_id"`
		Balance float64 `json:"balance" bson:"balance"`
	}

	UserTransaction struct {
		Id           primitive.ObjectID `bson:"_id,omitempty"`
		UserId       string             `bson:"user_id"`
		Amount       float64            `bson:"amount"`
		CreatedAt    time.Time          `bson:"created_at"`
		AddressId    string             `bson:"address_id"`
		OrderId      string             `bson:"order_id"`
		OrderSuccess string             `bson:"order_success"`
		UpdatedAt    time.Time          `bson:"updated_at"`
	}

	UserWishList struct {
		UserId   string   `json:"user_id" bson:"user_id"`
		WishList []string `json:"wishlist" bson:"wishlist"`
	}

	AddressModel struct {
		Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Name      string             `json:"name" bson:"name"`
		Phone     string             `json:"phone" bson:"phone"`
		Street    string             `json:"street" bson:"street"`
		City      string             `json:"city" bson:"city"`
		State     string             `json:"state" bson:"state"`
		Pincode   string             `json:"pincode" bson:"pincode"`
		Country   string             `json:"country" bson:"country"`
		CreatedAt time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	}
)
