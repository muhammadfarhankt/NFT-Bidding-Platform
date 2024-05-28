package nftHandler

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/request"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/response"
)

func (h *nftHttpHandler) CreateBid(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(nft.CreateBidReq)

	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

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

func (h *nftHttpHandler) EditBid(c echo.Context) error {
	return nil
}

func (h *nftHttpHandler) DeleteBid(c echo.Context) error {
	return nil
}

func (h *nftHttpHandler) BidNft(c echo.Context) error {
	return nil
}

func (h *nftHttpHandler) WithdrawBid(c echo.Context) error {
	return nil
}
