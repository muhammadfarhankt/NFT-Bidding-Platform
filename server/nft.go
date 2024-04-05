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

	// _ = httpHandler
	_ = grpcHandler

	nft := s.app.Group("/nft_v1")

	// Health Check
	nft.GET("", s.healthCheckService)

	// Create NFT
	nft.POST("/nft", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(httpHandler.CreateNft, []int{1, 0})))

	//Find one NFT
	nft.GET("/nft/:nft_id", httpHandler.FindOneNft)

	//Find many NFTs
	nft.GET("/nft", httpHandler.FindManyNfts)

	//Edit NFT
	nft.PATCH("/nft/:nft_id", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(httpHandler.EditNft, []int{1, 0})))

	//Block or Unblock NFT
	nft.PATCH("/nft/:nft_id/block-unblock", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(httpHandler.BlockOrUnblockNft, []int{1, 0})))
}
