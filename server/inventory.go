package server

import (
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/inventory/inventoryHandler"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/inventory/inventoryRepository"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	repo := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	// Health Check
	inventory.GET("", s.healthCheckService)
}
