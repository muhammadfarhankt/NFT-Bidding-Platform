package authHandler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/auth"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/auth/authUsecase"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/request"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/response"
)

type (
	AuthHttpHandlerService interface {
		Login(c echo.Context) error
		RefreshToken(c echo.Context) error
		Logout(c echo.Context) error
		OtpRequest(c echo.Context) error
		OtpVerification(c echo.Context) error
	}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authUsecase.AuthUsecaseService) AuthHttpHandlerService {
	return &authHttpHandler{cfg, authUsecase}
}

func (h *authHttpHandler) Login(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(auth.UserLoginReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.authUsecase.Login(ctx, h.cfg, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *authHttpHandler) RefreshToken(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(auth.RefreshTokenReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	log.Println("RefreshTokenReq: ", req)
	res, err := h.authUsecase.RefreshToken(ctx, h.cfg, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *authHttpHandler) Logout(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(auth.LogoutReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.authUsecase.Logout(ctx, req.CredentialId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, &response.MsgResponse{
		Message: fmt.Sprintf("Logged Out Successfully : %d", res),
	})
}

func (h *authHttpHandler) OtpRequest(c echo.Context) error {

	ctx := context.Background()
	// wrapper := request.ContextWrapper(c)

	// sending email as path variables in request
	email := c.Param("email")
	// fmt.Println("email: ", email)

	// req := new(auth.OtpRequestReq)

	// fmt.Println("req: ", req)

	if email == "" {
		return response.ErrResponse(c, http.StatusBadRequest, "Email is required")
	}

	// email validation using govalidator
	// if !request.ValidateEmail(req.Email) {
	// 	return response.ErrResponse(c, http.StatusBadRequest, "Invalid Email")
	// }

	// if err := wrapper.Bind(email); err != nil {
	// 	return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	// }

	// fmt.Println("username : ", h.cfg.Email.Username)

	err := h.authUsecase.OtpRequest(ctx, email, h.cfg)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "OTP Sent Successfully")
}

func (h *authHttpHandler) OtpVerification(c echo.Context) error {

	ctx := context.Background()
	wrapper := request.ContextWrapper(c)
	req := new(auth.OtpVerificationReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	// email := c.QueryParam("email")

	fmt.Println("otp login req: ", req)

	res, err := h.authUsecase.OtpVerification(ctx, h.cfg, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	// if res == nil {
	// 	return response.ErrResponse(c, http.StatusBadRequest, "OTP Verification Success")
	// }

	return response.SuccessResponse(c, http.StatusOK, res)
}
