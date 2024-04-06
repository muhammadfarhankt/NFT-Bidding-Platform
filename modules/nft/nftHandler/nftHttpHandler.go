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
		CreateCategory(c echo.Context) error
		FindOneCategory(c echo.Context) error
		FindManyCategories(c echo.Context) error
		EditCategory(c echo.Context) error
		BlockOrUnblockCategory(c echo.Context) error
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

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.nftUsecase.CreateNft(ctx, req)
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

	nftId := strings.TrimPrefix(c.Param("nft_id"), "nft:")

	wrapper := request.ContextWrapper(c)

	req := new(nft.NftUpdateReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.nftUsecase.EditNft(ctx, nftId, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *nftHttpHandler) BlockOrUnblockNft(c echo.Context) error {
	ctx := context.Background()

	nftId := strings.TrimPrefix(c.Param("nft_id"), "nft:")

	res, err := h.nftUsecase.BlockOrUnblockNft(ctx, nftId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, map[string]any{
		"message": fmt.Sprintf("nft_id: %s is successfully changed to: %v", nftId, res),
	})
}

func (h *nftHttpHandler) CreateCategory(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(nft.NftCategoryReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.nftUsecase.CreateCategory(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *nftHttpHandler) FindOneCategory(c echo.Context) error {
	ctx := context.Background()

	// categoryId := strings.TrimPrefix(c.Param("nft_id"), "nft:")
	categoryId := c.Param("category_id")

	res, err := h.nftUsecase.FindOneCategory(ctx, categoryId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *nftHttpHandler) FindManyCategories(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(nft.NftSearchReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.nftUsecase.FindManyCategories(ctx, h.cfg.Paginate.NftNextPageBasedUrl, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *nftHttpHandler) EditCategory(c echo.Context) error {
	ctx := context.Background()

	categoryId := c.Param("category_id")

	wrapper := request.ContextWrapper(c)

	req := new(nft.NftCategory)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.nftUsecase.EditCategory(ctx, categoryId, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *nftHttpHandler) BlockOrUnblockCategory(c echo.Context) error {
	ctx := context.Background()

	categoryId := c.Param("category_id")

	res, err := h.nftUsecase.BlockOrUnblockCategory(ctx, categoryId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	req := new(nft.NftCategory)
	category, _ := h.nftUsecase.EditCategory(ctx, categoryId, req)

	return response.SuccessResponse(c, http.StatusOK, map[string]any{
		"message":  fmt.Sprintf("CategoryId : %s is successfully changed to: %v", categoryId, res),
		"category": category,
	})
}
