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

	// -------------------------------- NFT Public End Points -------------------------------- //
	//Find one NFT
	nft.GET("/nft/:nft_id", httpHandler.FindOneNft)

	//Find many NFTs
	nft.GET("/nft", httpHandler.FindManyNfts)

	//Find one Category
	nft.GET("/nft/category/:category_id", httpHandler.FindOneCategory)

	//Find many Categories
	nft.GET("/nft/category", httpHandler.FindManyCategories)

	// Find Top 10 Wishlist NFTs
	nft.GET("/nft/top-wishlist", httpHandler.FindTopWishlistNfts)

	// Find Top 10 Bidding NFTs
	nft.GET("/nft/top-bidding", httpHandler.FindTopBiddingNfts)

	// -------------------------------- NFT Owner End Points -------------------------------- //

	// Create NFT
	nft.POST("/nft", httpHandler.CreateNft, s.middleware.JwtAuthorization)

	//Edit NFT
	nft.PATCH("/nft/:nft_id", httpHandler.EditNft, s.middleware.JwtAuthorization)

	//Block or Unblock NFT
	nft.PATCH("/nft/:nft_id/block-unblock", httpHandler.BlockOrUnblockNft, s.middleware.JwtAuthorization)

	//Delete NFT
	nft.DELETE("/nft/:nft_id", httpHandler.DeleteNft, s.middleware.JwtAuthorization)

	// -------------------------------- NFT Image -------------------------------- //

	// Upload NFT Image
	nft.POST("/image", httpHandler.UploadToGCP, s.middleware.JwtAuthorization)

	// Delete NFT Image
	nft.PATCH("/image", httpHandler.DeleteFromGCP, s.middleware.JwtAuthorization)

	// -------------------------------- NFT Admin End Points -------------------------------- //

	//Create Category
	nft.POST("/nft/category", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(httpHandler.CreateCategory, []int{1, 0})))

	//Edit Category
	nft.PATCH("/nft/category/:category_id", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(httpHandler.EditCategory, []int{1, 0})))

	//Block or Unblock Category
	nft.GET("/nft/category/:category_id/block-unblock", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(httpHandler.BlockOrUnblockCategory, []int{1, 0})))

	//Delete Category
	nft.DELETE("/nft/category/:category_id", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(httpHandler.DeleteCategory, []int{1, 0})))

	// -------------------------------- Bidding NFT Owner End Points -------------------------------- //
	// show all bids
	nft.GET("/nft/bid", httpHandler.FindManyBids, s.middleware.JwtAuthorization)

	// create bid
	nft.POST("/nft/bid", httpHandler.CreateBid, s.middleware.JwtAuthorization)

	// edit bid
	nft.PATCH("/nft/bid/:bid_id", httpHandler.EditBid, s.middleware.JwtAuthorization)

	// delete bid
	nft.DELETE("/nft/bid/:bid_id", httpHandler.DeleteBid, s.middleware.JwtAuthorization)

	// -------------------------------- Bidding User End Points -------------------------------- //

	// Find all bids of a user
	nft.GET("/nft/bidding", httpHandler.FindUserBids, s.middleware.JwtAuthorization)

	// user bidding a NFT
	nft.POST("/nft/bidding/", httpHandler.BidNft, s.middleware.JwtAuthorization)

	// user withdraw a bid
	nft.PATCH("/nft/bidding/:bid_id/withdraw", httpHandler.WithdrawBid, s.middleware.JwtAuthorization)

	// -------------------------------- Bidding Admin End Points -------------------------------- //
	nft.GET("/nft/bidding/execute-bids", httpHandler.ExecuteBids)

}
