package user

import "time"

type (
	UserProfile struct {
		Id           string    `json:"_id"`
		Email        string    `json:"email"`
		Username     string    `json:"username"`
		WalletAmount float64   `json:"wallet_amount"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		IsBlocked    bool      `json:"is_blocked"`
	}

	UserClaims struct {
		Id       string `json:"id"`
		RoleCode int    `json:"role_code"`
	}

	CreateUserReq struct {
		Email    string `json:"email" form:"email" validate:"required,email,max=255"`
		Password string `json:"password" form:"password" validate:"required,max=32"`
		Username string `json:"username" form:"username" validate:"required,max=64"`
	}

	CreateUserTransactionReq struct {
		UserId    string  `json:"user_id" validate:"required,max=64"`
		Amount    float64 `json:"amount" validate:"required"`
		AddressId string  `json:"address_id" validate:"required"`
	}

	CreateUserAddressReq struct {
		Name    string `json:"name" form:"name" validate:"required,max=255"`
		Street  string `json:"street" form:"street" validate:"required,max=255"`
		City    string `json:"city" form:"city" validate:"required,max=255"`
		Phone   string `json:"phone" form:"phone" validate:"required,max=255"`
		Pincode string `json:"pincode" form:"pincode" validate:"required,max=255"`
		State   string `json:"state" form:"state" validate:"required,max=255"`
		Country string `json:"country" form:"country" validate:"required,max=255"`
	}

	ResetPasswordReq struct {
		OldPassword string `json:"old_password" form:"old_password" validate:"required,max=32"`
		NewPassword string `json:"new_password" form:"new_password" validate:"required,max=32"`
	}
)
