package userHandler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userUsecase"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/request"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/response"
)

type (
	UserHttpHandlerService interface {
		InsertUser(c echo.Context) error
		FindOneUserProfile(c echo.Context) error
		AddToWallet(c echo.Context) error
		GetUserWalletAccount(c echo.Context) error
	}

	userHttpHandler struct {
		cfg         *config.Config
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserHttpHandler(cfg *config.Config, userUsecase userUsecase.UserUsecaseService) UserHttpHandlerService {
	return &userHttpHandler{userUsecase: userUsecase}
}

func (h *userHttpHandler) InsertUser(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(user.CreateUserReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.userUsecase.InsertUser(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *userHttpHandler) FindOneUserProfile(c echo.Context) error {
	ctx := context.Background()

	userId := strings.TrimPrefix(c.Param("user_id"), "user:")

	fmt.Println("userId", userId)

	if userId == ":user_id" || userId == "" {
		return response.ErrResponse(c, http.StatusBadRequest, "userId cannot be empty")
	}

	res, err := h.userUsecase.FindOneUserProfile(ctx, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusBadRequest, res)
}

func (h *userHttpHandler) AddToWallet(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(user.CreateUserTransactionReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.userUsecase.AddToWallet(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *userHttpHandler) GetUserWalletAccount(c echo.Context) error {
	ctx := context.Background()

	userId := c.Param("user_id")

	if userId == ":user_id" || userId == "" {
		return response.ErrResponse(c, http.StatusBadRequest, "userId cannot be empty")
	}

	res, err := h.userUsecase.GetUserWalletAccount(ctx, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusBadRequest, res)
}
