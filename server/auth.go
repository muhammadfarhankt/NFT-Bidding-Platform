package server

import (
	"log"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/auth/authHandler"
	authPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/auth/authPb"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/auth/authRepository"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/auth/authUsecase"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/grpcConn"
)

func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpcConn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Auth gRPC server listening on %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	// Health Check
	auth.GET("", s.healthCheckService)
}
