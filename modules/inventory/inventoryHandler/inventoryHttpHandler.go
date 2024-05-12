package inventoryHandler

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/inventory"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/inventory/inventoryUsecase"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/request"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/response"
)

type (
	InventoryHttpHandlerService interface {
		FindUserNfts(c echo.Context) error
	}

	inventoryHttpHandler struct {
		cfg              *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryHttpHandler(cfg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryHttpHandlerService {
	return &inventoryHttpHandler{
		cfg:              cfg,
		inventoryUsecase: inventoryUsecase,
	}
}

func (h *inventoryHttpHandler) FindUserNfts(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(inventory.InventorySearchReq)
	userId := c.Param("user_id")
	log.Println("User ID: ", userId)
	log.Println("Request: ", req)
	if err := wrapper.Bind(req); err != nil {
		log.Println("Error: Bind data failed: ", err.Error())
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.inventoryUsecase.FindUserNfts(ctx, h.cfg, userId, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
