package server

import (
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftHandler"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftRepository"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftUsecase"
)

func (s *server) nftService() {
	repo := nftRepository.NewNftRepository(s.db)
	usecase := nftUsecase.NewNftUsecase(repo)
	httpHandler := nftHandler.NewNftHttpHandler(s.cfg, usecase)
	grpcHandler := nftHandler.NewNftGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	nft := s.app.Group("/nft_v1")

	// Health Check
	nft.GET("", s.healthCheckService)
}
