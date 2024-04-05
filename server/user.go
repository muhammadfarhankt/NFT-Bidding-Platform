package server

import (
	"log"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userHandler"
	userPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userPb"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userRepository"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userUsecase"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/grpcConn"
)

func (s *server) userService() {
	repo := userRepository.NewUserRepository(s.db)
	usecase := userUsecase.NewUserUsecase(repo)
	httpHandler := userHandler.NewUserHttpHandler(s.cfg, usecase)
	grpcHandler := userHandler.NewUserGrpcHandler(usecase)
	queueHandler := userHandler.NewUserQueueHandler(s.cfg, usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpcConn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.UserUrl)

		userPb.RegisterUserGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("User gRPC server listening on %s", s.cfg.Grpc.UserUrl)
		grpcServer.Serve(lis)
	}()

	// _ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	user := s.app.Group("/user_v1")

	// Health Check
	// user.GET("", s.healthCheckService, s.middleware.JwtAuthorization)
	user.POST("", s.healthCheckService)
	user.POST("/user/register", httpHandler.InsertUser)
	user.GET("/user/:user_id", httpHandler.FindOneUserProfile)
	user.POST("/user/add-wallet-money/:user_id", httpHandler.AddToWallet)
	user.GET("/user/wallet/:user_id", httpHandler.GetUserWalletAccount)
}
