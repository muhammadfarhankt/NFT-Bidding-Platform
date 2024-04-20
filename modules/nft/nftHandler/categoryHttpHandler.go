package nftHandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/request"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/response"
)

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

func (h *nftHttpHandler) DeleteCategory(c echo.Context) error {
	ctx := context.Background()

	categoryId := c.Param("category_id")

	res, err := h.nftUsecase.DeleteCategory(ctx, categoryId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	req := new(nft.NftCategory)
	category, _ := h.nftUsecase.EditCategory(ctx, categoryId, req)

	return response.SuccessResponse(c, http.StatusOK, map[string]any{
		"message":  fmt.Sprintf("CategoryId : %s is successfully soft deleted (Blocked): %v", categoryId, res),
		"category": category,
	})
}
