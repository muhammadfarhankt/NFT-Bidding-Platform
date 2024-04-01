package server

import (
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userHandler"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userRepository"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userUsecase"
)

func (s *server) userService() {
	repo := userRepository.NewUserRepository(s.db)
	usecase := userUsecase.NewUserUsecase(repo)
	httpHandler := userHandler.NewUserHttpHandler(s.cfg, usecase)
	grpcHandler := userHandler.NewUserGrpcHandler(usecase)
	queueHandler := userHandler.NewUserQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	user := s.app.Group("/user_v1")

	// Health Check
	user.GET("", s.healthCheckService)
}
