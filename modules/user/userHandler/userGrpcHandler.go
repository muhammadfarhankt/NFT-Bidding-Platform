package userHandler

import "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userUsecase"

type (
	userGrpcHandlerService struct {
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserGrpcHandler(userUsecase userUsecase.UserUsecaseService) *userGrpcHandlerService {
	return &userGrpcHandlerService{userUsecase}
}
