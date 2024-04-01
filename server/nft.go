package server

import (
	"log"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftHandler"
	nftPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftPb"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftRepository"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftUsecase"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/grpcConn"
)

func (s *server) nftService() {
	repo := nftRepository.NewNftRepository(s.db)
	usecase := nftUsecase.NewNftUsecase(repo)
	httpHandler := nftHandler.NewNftHttpHandler(s.cfg, usecase)
	grpcHandler := nftHandler.NewNftGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpcConn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.NftUrl)

		nftPb.RegisterNftGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Nft gRPC server listening on %s", s.cfg.Grpc.NftUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	nft := s.app.Group("/nft_v1")

	// Health Check
	nft.GET("", s.healthCheckService)
}
