package auth

import (
	"time"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user"
)

type (
	UserLoginReq struct {
		Email    string `json:"email" form:"email" validate:"required,email,max=255"`
		Password string `json:"password" form:"password" validate:"required,max=32"`
	}

	RefreshTokenReq struct {
		CredentialId string `json:"credential_id" form:"credential_id" validate:"required,max=64"`
		RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required,max=500"`
	}

	InsertUserRole struct {
		UserId   string `json:"user_id" validate:"required"`
		RoleCode []int  `json:"role_id" validate:"required"`
	}

	ProfileIntercepter struct {
		*user.UserProfile
		Credential *CredentialRes `json:"credential"`
	}

	CredentialRes struct {
		Id           string    `json:"_id"`
		UserId       string    `json:"user_id"`
		RoleCode     int       `json:"role_code"`
		AccessToken  string    `json:"access_token"`
		RefreshToken string    `json:"refresh_token"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	LogoutReq struct {
		CredentialId string `json:"credential_id" form:"credential_id" validate:"required,max=64"`
	}

	OtpRequestReq struct {
		Email string `json:"email" form:"email" validate:"required,max=255"`
	}

	OtpVerificationReq struct {
		Email string `json:"email" form:"email" validate:"required,max=255"`
		Otp   string `json:"otp" form:"otp" validate:"required,max=6"`
	}

	OtpModel struct {
		Email     string    `json:"email" bson:"email"`
		Otp       string    `json:"otp" bson:"otp"`
		ExpiresAt time.Time `json:"expires_at" bson:"expires_at"`
	}
)
