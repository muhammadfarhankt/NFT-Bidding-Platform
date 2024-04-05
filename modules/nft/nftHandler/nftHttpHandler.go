package nftHandler

import (
	"context"
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
