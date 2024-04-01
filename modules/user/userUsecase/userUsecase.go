package userUsecase

import "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userRepository"

type (
	UserUsecaseService interface{}

	userUsecase struct {
		userRepository userRepository.UserRepositoryService
	}
)

func NewUserUsecase(userRepository userRepository.UserRepositoryService) UserUsecaseService {
	return &userUsecase{userRepository: userRepository}
}
