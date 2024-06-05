package nftHandler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/request"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/response"
)

func (h *nftHttpHandler) FindManyBids(c echo.Context) error {

	ctx := context.Background()

	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")

	fmt.Println("user id: ", userId)

	res, err := h.nftUsecase.FindManyBids(ctx, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *nftHttpHandler) CreateBid(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(nft.CreateBidReq)

	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	fmt.Println("create bid req: ", req)
	// check nft_id is required
	if req.NftId == "" {
		return response.ErrResponse(c, http.StatusBadRequest, "nft_id is required")
	}

	// check bid expiry date is in future
	if req.ExpiryDate.Before(time.Now()) {
		return response.ErrResponse(c, http.StatusBadRequest, "expiry date must be in future")
	}

	// check price is greater than 0
	if req.Price <= 0 {
		return response.ErrResponse(c, http.StatusBadRequest, "price must be greater than 0")
	}

	res, err := h.nftUsecase.CreateBid(ctx, req, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *nftHttpHandler) DeleteBid(c echo.Context) error {

	ctx := context.Background()

	bidId := c.Param("bid_id")

	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")

	fmt.Println("bid id: ", bidId)

	if bidId == "" || bidId == ":bid_id" {
		return response.ErrResponse(c, http.StatusBadRequest, "bid_id is required")
	}

	// fmt.Println("bid id: ", bidId)
	// fmt.Println("user id: ", userId)

	err := h.nftUsecase.DeleteBid(ctx, bidId, userId)

	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Bid Deleted Successfully")
}

func (h *nftHttpHandler) EditBid(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(nft.CreateBidReq)

	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	fmt.Println("create bid req: ", req)

	bidId := c.Param("bid_id")

	// fmt.Println("bid id: ", bidId)

	if bidId == "" || bidId == ":bid_id" {
		return response.ErrResponse(c, http.StatusBadRequest, "bid_id is required")
	}

	res, err := h.nftUsecase.EditBid(ctx, bidId, userId, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

// ------------------- NFT Bidding User ------------------- //

func (h *nftHttpHandler) FindUserBids(c echo.Context) error {

	ctx := context.Background()

	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")
	// fmt.Println("user id: ", userId)

	res, err := h.nftUsecase.FindUserBids(ctx, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *nftHttpHandler) WithdrawBid(c echo.Context) error {
	ctx := context.Background()

	// get userId from context
	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")
	fmt.Println("user id: ", userId)

	bidId := c.Param("bid_id")
	if bidId == "" || bidId == ":bid_id" {
		return response.ErrResponse(c, http.StatusBadRequest, "bid_id is required")
	}

	// fmt.Println("bid id: ", bidId)

	err := h.nftUsecase.WithdrawBid(ctx, bidId, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Bid Withdrawn Successfully")
}

func (h *nftHttpHandler) BidNft(c echo.Context) error {

	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(nft.CreateUserBidReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	// get userId from context
	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")
	fmt.Println("user id: ", userId)

	// get nftId from request
	nftId := req.NftId
	if nftId == "" {
		return response.ErrResponse(c, http.StatusBadRequest, "nft_id is required")
	}
	// fmt.Println("nft id: ", nftId)

	// get price from request and convert it to string
	if req.Price <= 0 {
		return response.ErrResponse(c, http.StatusBadRequest, "price must be greater than 0")
	}
	price := fmt.Sprintf("%f", req.Price)
	// fmt.Println("price: ", price)

	result, err := h.nftUsecase.BidNft(ctx, nftId, userId, price)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, result)
}
