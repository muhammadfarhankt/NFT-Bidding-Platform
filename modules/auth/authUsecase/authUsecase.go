package authUsecase

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/auth"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/auth/authRepository"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/jwtAuth"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/utils"

	authPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/auth/authPb"
	userPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userPb"
)

type (
	AuthUsecaseService interface {
		Login(pctx context.Context, cfg *config.Config, req *auth.UserLoginReq) (*auth.ProfileIntercepter, error)
		RefreshToken(pctx context.Context, cfg *config.Config, req *auth.RefreshTokenReq) (*auth.ProfileIntercepter, error)
		Logout(pctx context.Context, credentialId string) (int64, error)
		AccessTokenSearch(pctx context.Context, accessToken string) (*authPb.AccessTokenSearchRes, error)
		RolesCount(pctx context.Context) (*authPb.RolesCountRes, error)
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
			RoleCode:     credential.RoleCode,
			AccessToken:  credential.AccessToken,
			RefreshToken: credential.RefreshToken,
			CreatedAt:    credential.CreatedAt.In(loc),
			UpdatedAt:    credential.UpdatedAt.In(loc),
		},
	}, nil
}

func (u *authUsecase) RefreshToken(pctx context.Context, cfg *config.Config, req *auth.RefreshTokenReq) (*auth.ProfileIntercepter, error) {
	claims, err := jwtAuth.ParseToken(cfg.Jwt.RefreshSecretKey, req.RefreshToken)
	if err != nil {
		log.Printf("Error: RefreshToken: %s", err.Error())
		return nil, errors.New(err.Error())
	}

	profile, err := u.authRepository.FindOneUserProfileToRefresh(pctx, cfg.Grpc.UserUrl, &userPb.FindOneUserProfileToRefreshReq{
		UserId: strings.TrimPrefix(claims.UserId, "user:"),
	})
	if err != nil {
		return nil, err
	}

	accessToken := jwtAuth.NewAccessToken(cfg.Jwt.AccessSecretKey, cfg.Jwt.AccessDuration, &jwtAuth.Claims{
		UserId:   profile.Id,
		RoleCode: int(profile.RoleCode),
	}).SignToken()

	refreshToken := jwtAuth.ReloadToken(cfg.Jwt.RefreshSecretKey, claims.ExpiresAt.Unix(), &jwtAuth.Claims{
		UserId:   profile.Id,
		RoleCode: int(profile.RoleCode),
	})

	if err := u.authRepository.UpdateOneUserCredential(pctx, req.CredentialId, &auth.UpdateRefreshTokenReq{
		UserId:       profile.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UpdatedAt:    utils.LocalTime(),
	}); err != nil {
		return nil, err
	}

	credential, err := u.authRepository.FindOneUserCredential(pctx, req.CredentialId)
	if err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("Asia/Calcutta")

	return &auth.ProfileIntercepter{
		UserProfile: &user.UserProfile{
			Id:        "user:" + profile.Id,
			Email:     profile.Email,
			Username:  profile.Username,
			CreatedAt: utils.ConvertStringTimeToTime(profile.CreatedAt),
			UpdatedAt: utils.ConvertStringTimeToTime(profile.UpdatedAt),
		},
		Credential: &auth.CredentialRes{
			Id:           credential.Id.Hex(),
			UserId:       credential.UserId,
			RoleCode:     credential.RoleCode,
			AccessToken:  credential.AccessToken,
			RefreshToken: credential.RefreshToken,
			CreatedAt:    credential.CreatedAt.In(loc),
			UpdatedAt:    credential.UpdatedAt.In(loc),
		},
	}, nil
}

func (u *authUsecase) Logout(pctx context.Context, credentialId string) (int64, error) {
	return u.authRepository.DeleteOneUserCredential(pctx, credentialId)
}

func (u *authUsecase) AccessTokenSearch(pctx context.Context, accessToken string) (*authPb.AccessTokenSearchRes, error) {
	credential, err := u.authRepository.FindOneAccessToken(pctx, accessToken)
	if err != nil {
		return &authPb.AccessTokenSearchRes{
			IsValid: false,
		}, err
	}

	if credential == nil {
		return &authPb.AccessTokenSearchRes{
			IsValid: false,
		}, errors.New("error: access token is invalid")
	}

	return &authPb.AccessTokenSearchRes{
		IsValid: true,
	}, nil
}

func (u *authUsecase) RolesCount(pctx context.Context) (*authPb.RolesCountRes, error) {
	result, err := u.authRepository.RolesCount(pctx)
	if err != nil {
		return nil, err
	}

	return &authPb.RolesCountRes{
		Count: result,
	}, nil
}
