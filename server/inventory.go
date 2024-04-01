package server

import (
	"log"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/inventory/inventoryHandler"
	inventoryPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/inventory/inventoryPb"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/inventory/inventoryRepository"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/inventory/inventoryUsecase"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/grpcConn"
)

func (s *server) inventoryService() {
	repo := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpcConn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.InventUrl)

		inventoryPb.RegisterInventoryGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Inventory gRPC server listening on %s", s.cfg.Grpc.InventUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	// Health Check
	inventory.GET("", s.healthCheckService)
}
