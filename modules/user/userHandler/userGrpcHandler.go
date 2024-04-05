package userHandler

import (
	"context"

	userPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userPb"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userUsecase"
)

type (
	userGrpcHandler struct {
		userUsecase userUsecase.UserUsecaseService
		userPb.UnimplementedUserGrpcServiceServer
	}
)

func NewUserGrpcHandler(userUsecase userUsecase.UserUsecaseService) *userGrpcHandler {
	return &userGrpcHandler{
		userUsecase: userUsecase,
	}
}

func (g *userGrpcHandler) CredentialSearch(ctx context.Context, req *userPb.CredentialSearchReq) (*userPb.UserProfile, error) {
	return g.userUsecase.FindOneEmail(ctx, req.Password, req.Email)
}

func (g *userGrpcHandler) FindOneUserProfileToRefresh(ctx context.Context, req *userPb.FindOneUserProfileToRefreshReq) (*userPb.UserProfile, error) {
	return g.userUsecase.FindOneUserProfileToRefresh(ctx, req.UserId)
}

func (g *userGrpcHandler) GetUserWalletAccount(ctx context.Context, req *userPb.GetUserWalletAccountReq) (*userPb.GetUserWalletAccountRes, error) {
	return nil, nil
}
