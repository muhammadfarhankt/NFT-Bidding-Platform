package nftHandler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftUsecase"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/request"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/response"
)

type (
	NftHttpHandlerService interface {
		CreateNft(c echo.Context) error
		FindOneNft(c echo.Context) error
		FindManyNfts(c echo.Context) error
		EditNft(c echo.Context) error
		BlockOrUnblockNft(c echo.Context) error
		DeleteNft(c echo.Context) error

		FindTopWishlistNfts(c echo.Context) error
		FindTopBiddingNfts(c echo.Context) error

		// -------------------- NFT Image -------------------- //
		UploadToGCP(c echo.Context) error
		DeleteFromGCP(c echo.Context) error

		// -------------------- Category -------------------- //
		CreateCategory(c echo.Context) error
		FindOneCategory(c echo.Context) error
		FindManyCategories(c echo.Context) error
		EditCategory(c echo.Context) error
		BlockOrUnblockCategory(c echo.Context) error
		DeleteCategory(c echo.Context) error

		// -------------------- NFT Bidding Owner -------------------- //
		FindManyBids(c echo.Context) error
		CreateBid(c echo.Context) error
		EditBid(c echo.Context) error
		DeleteBid(c echo.Context) error

		// -------------------- NFT Bidding User -------------------- //
		FindUserBids(c echo.Context) error
		BidNft(c echo.Context) error
		WithdrawBid(c echo.Context) error

		// -------------------- NFT Bidding Admin -------------------- //
		ExecuteBids(c echo.Context) error
	}

	nftHttpHandler struct {
		cfg        *config.Config
		nftUsecase nftUsecase.NftUsecaseService
	}
)

func NewNftHttpHandler(cfg *config.Config, nftUsecase nftUsecase.NftUsecaseService) NftHttpHandlerService {
	return &nftHttpHandler{
		cfg:        cfg,
		nftUsecase: nftUsecase,
	}
}

func (h *nftHttpHandler) CreateNft(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(nft.CreateNftReq)

	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.nftUsecase.CreateNft(ctx, req, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *nftHttpHandler) FindOneNft(c echo.Context) error {
	ctx := context.Background()

	nftId := strings.TrimPrefix(c.Param("nft_id"), "nft:")

	res, err := h.nftUsecase.FindOneNft(ctx, nftId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *nftHttpHandler) FindManyNfts(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(nft.NftSearchReq)

	//fmt.Println("Findmany nfts req: ", req.Category, req.Title, req.Start, req.Limit)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.nftUsecase.FindManyNfts(ctx, h.cfg.Paginate.NftNextPageBasedUrl, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *nftHttpHandler) EditNft(c echo.Context) error {
	ctx := context.Background()

	nftId := string(c.Param("nft_id"))
	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")
	wrapper := request.ContextWrapper(c)

	req := new(nft.NftUpdateReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.nftUsecase.EditNft(ctx, nftId, userId, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *nftHttpHandler) BlockOrUnblockNft(c echo.Context) error {
	ctx := context.Background()

	nftId := strings.TrimPrefix(c.Param("nft_id"), "nft:")

	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")

	res, err := h.nftUsecase.BlockOrUnblockNft(ctx, nftId, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if res {
		return response.SuccessResponse(c, http.StatusOK, map[string]any{
			"message": fmt.Sprintf("nft_id: %s is successfully blocked", nftId),
		})
	}

	return response.SuccessResponse(c, http.StatusOK, map[string]any{
		"message": fmt.Sprintf("nft_id: %s is successfully unblocked", nftId),
	})
}

func (h *nftHttpHandler) DeleteNft(c echo.Context) error {
	ctx := context.Background()

	nftId := strings.TrimPrefix(c.Param("nft_id"), "nft:")

	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")

	_, err := h.nftUsecase.DeleteNft(ctx, nftId, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, map[string]any{
		"message": fmt.Sprintf("nft_id: %s is successfully deleted (Soft Delete)", nftId),
	})
}

func (h *nftHttpHandler) FindTopWishlistNfts(c echo.Context) error {
	ctx := context.Background()

	res, err := h.nftUsecase.FindTopWishlistNfts(ctx)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *nftHttpHandler) FindTopBiddingNfts(c echo.Context) error {
	ctx := context.Background()

	res, err := h.nftUsecase.FindTopBiddingNfts(ctx)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
