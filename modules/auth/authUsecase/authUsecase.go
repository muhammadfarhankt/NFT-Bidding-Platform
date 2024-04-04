package authUsecase

import (
	"context"
	"time"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/auth"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/auth/authRepository"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/jwtAuth"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/utils"

	userPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userPb"
)

type (
	AuthUsecaseService interface {
		Login(pctx context.Context, cfg *config.Config, req *auth.UserLoginReq) (*auth.ProfileIntercepter, error)
	}

	authUsecase struct {
		authRepository authRepository.AuthRepositoryService
	}
)

func NewAuthUsecase(authRepository authRepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{authRepository}
}

func (u *authUsecase) Login(pctx context.Context, cfg *config.Config, req *auth.UserLoginReq) (*auth.ProfileIntercepter, error) {
	profile, err := u.authRepository.CredentialSearch(pctx, cfg.Grpc.UserUrl, &userPb.CredentialSearchReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	profile.Id = "user:" + profile.Id

	accessToken := jwtAuth.NewAccessToken(cfg.Jwt.AccessSecretKey, cfg.Jwt.AccessDuration, &jwtAuth.Claims{
		UserId:   profile.Id,
		RoleCode: int(profile.RoleCode),
	}).SignToken()

	refreshToken := jwtAuth.NewRefreshToken(cfg.Jwt.RefreshSecretKey, cfg.Jwt.RefreshDuration, &jwtAuth.Claims{
		UserId:   profile.Id,
		RoleCode: int(profile.RoleCode),
	}).SignToken()

	credentialId, err := u.authRepository.InsertOneUserCredential(pctx, &auth.Credential{
		UserId:       profile.Id,
		RoleCode:     int(profile.RoleCode),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		CreatedAt:    utils.LocalTime(),
		UpdatedAt:    utils.LocalTime(),
	})

	credential, err := u.authRepository.FindOneUserCredential(pctx, credentialId.Hex())
	if err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("Asia/Calcutta")

	return &auth.ProfileIntercepter{
		UserProfile: &user.UserProfile{
			Id:        profile.Id,
			Email:     profile.Email,
			Username:  profile.Username,
			CreatedAt: utils.ConvertStringTimeToTime(profile.CreatedAt).In(loc),
			UpdatedAt: utils.ConvertStringTimeToTime(profile.UpdatedAt).In(loc),
		},
		Credential: &auth.CredentialRes{
			Id:           credential.Id.Hex(),
			UserId:       credential.UserId,
			RoleCode:     int(credential.RoleCode),
			AccessToken:  credential.AccessToken,
			RefreshToken: credential.RefreshToken,
			CreatedAt:    credential.CreatedAt.In(loc),
			UpdatedAt:    credential.UpdatedAt.In(loc),
		},
	}, nil
}
