package userHandler

import (
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userUsecase"
)

type (
	UserQueueHandlerService interface{}

	userQueueHandler struct {
		cfg         *config.Config
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserQueueHandler(cfg *config.Config, userUsecase userUsecase.UserUsecaseService) UserQueueHandlerService {
	return &userQueueHandler{userUsecase: userUsecase}
}
