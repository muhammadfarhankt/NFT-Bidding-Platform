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
	// _ = grpcHandler
	_ = queueHandler

	user := s.app.Group("/user_v1")

	// Health Check
	// user.GET("", s.healthCheckService, s.middleware.JwtAuthorization)
	user.GET("", s.healthCheckService)

	// ----------------- USER ----------------- //
	user.POST("/user/register", httpHandler.InsertUser)

	user.GET("/user", httpHandler.FindOneUserProfile, s.middleware.JwtAuthorization)

	// reset password using old password
	user.PATCH("/user/reset-password", httpHandler.ResetPassword, s.middleware.JwtAuthorization)

	// ----------------- USER Payment & Wallet ----------------- //
	// user.POST("/user/add-wallet-money/:user_id", httpHandler.AddToWallet)
	user.POST("/user/add-wallet-money", httpHandler.AddToWallet, s.middleware.JwtAuthorization)

	user.GET("/user/payment", httpHandler.RazorPayLoad)
	user.POST("/user/payment/confirm", httpHandler.RazorPaymentConfirm)

	// user.GET("/user/wallet/:user_id", httpHandler.GetUserWalletAccount)
	user.GET("/user/wallet", httpHandler.GetUserWalletAccount, s.middleware.JwtAuthorization)

	// ----------------- Wish List ----------------- //
	user.POST("/user/wishlist/:nft_id", httpHandler.AddToWishList, s.middleware.JwtAuthorization)
	user.GET("/user/wishlist", httpHandler.GetWishList, s.middleware.JwtAuthorization)
	user.DELETE("/user/wishlist/:nft_id", httpHandler.RemoveFromWishList, s.middleware.JwtAuthorization)

	// ----------------- Address ----------------- //
	user.POST("/user/address/", httpHandler.AddAddress, s.middleware.JwtAuthorization)
	user.GET("/user/address/", httpHandler.GetAddress, s.middleware.JwtAuthorization)
	user.PATCH("/user/address/:address_id", httpHandler.UpdateAddress, s.middleware.JwtAuthorization)
	user.DELETE("/user/address/:address_id", httpHandler.DeleteAddress, s.middleware.JwtAuthorization)

	// ----------------- Payment Reports ----------------- //
	user.GET("/user/payment-report-pdf", httpHandler.UserPaymentReport, s.middleware.JwtAuthorization)
	// single order
	user.GET("/user/payment-report-pdf/:order_id", httpHandler.SingleOrderPaymentReport, s.middleware.JwtAuthorization)

	// ----------------- ADMIN ----------------- //

	// block or unblock user
	user.GET("/admin/:user_id/block-unblock", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(httpHandler.BlockOrUnblockUser, []int{1, 0})))

	// sales report
	user.GET("/admin/payment-report-pdf", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(httpHandler.SalesReport, []int{1})))
}
