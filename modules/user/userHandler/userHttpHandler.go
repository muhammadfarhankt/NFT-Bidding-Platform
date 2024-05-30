package userHandler

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/razorpay/razorpay-go"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userUsecase"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/request"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/response"
)

type (
	UserHttpHandlerService interface {
		InsertUser(c echo.Context) error
		FindOneUserProfile(c echo.Context) error

		// --- Wallet ---
		AddToWallet(c echo.Context) error
		GetUserWalletAccount(c echo.Context) error

		// --- Payment ---
		RazorPayLoad(c echo.Context) error
		RazorPaymentConfirm(c echo.Context) error

		// --- Wish List ---
		AddToWishList(c echo.Context) error
		GetWishList(c echo.Context) error
		RemoveFromWishList(c echo.Context) error

		// --- Address ---
		AddAddress(c echo.Context) error
		GetAddress(c echo.Context) error
		UpdateAddress(c echo.Context) error
		DeleteAddress(c echo.Context) error

		// --- Admin ---
		BlockOrUnblockUser(c echo.Context) error
		SalesReport(c echo.Context) error
	}

	userHttpHandler struct {
		cfg         *config.Config
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserHttpHandler(cfg *config.Config, userUsecase userUsecase.UserUsecaseService) UserHttpHandlerService {
	return &userHttpHandler{cfg, userUsecase}
}

func (h *userHttpHandler) InsertUser(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(user.CreateUserReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.userUsecase.InsertUser(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *userHttpHandler) FindOneUserProfile(c echo.Context) error {
	ctx := context.Background()

	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")

	//fmt.Println("userId", userId)

	if userId == ":user_id" || userId == "" {
		return response.ErrResponse(c, http.StatusBadRequest, "userId cannot be empty")
	}

	res, err := h.userUsecase.FindOneUserProfile(ctx, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusBadRequest, res)
}

func (h *userHttpHandler) AddToWallet(c echo.Context) error {

	fmt.Println("AddToWallet http handler")

	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(user.CreateUserTransactionReq)
	req.UserId = c.Get("user_id").(string)
	fmt.Println("req.UserId : ", req.UserId)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	addressId := req.AddressId

	fmt.Println("addressId : ", addressId)
	if addressId == "" {
		return response.ErrResponse(c, http.StatusBadRequest, errors.New("address_id cannot be empty").Error())
	}

	amount := int(req.Amount)

	if amount <= 0 {
		return response.ErrResponse(c, http.StatusBadRequest, errors.New(" Amount cannot be less than or equal to zero").Error())
	}

	fmt.Println("Amount : ", amount)

	// RazorPay Payment

	// var cfg *config.Razorpay

	razorPayKey := h.cfg.Razorpay.Key
	razorPaySecret := h.cfg.Razorpay.Secret

	// razorPayKey := "key"
	// razorPaySecret := "secret"

	// fmt.Println("key : ", razorPayKey)
	// fmt.Println("secret : ", razorPaySecret)

	orderId, err := PaymentHandler(amount, req.UserId, razorPayKey, razorPaySecret)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "orderId creation failed. "+err.Error())
	}

	res, err := h.userUsecase.AddToWallet(ctx, req, orderId)

	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	fmt.Println("add to wallet res : ", res)

	return response.SuccessResponse(c, http.StatusCreated, orderId)
}

func (h *userHttpHandler) GetUserWalletAccount(c echo.Context) error {
	ctx := context.Background()

	// userId := c.Param("user_id")
	userId := c.Get("user_id").(string)

	if userId == ":user_id" || userId == "" {
		return response.ErrResponse(c, http.StatusBadRequest, "userId cannot be empty")
	}

	res, err := h.userUsecase.GetUserWalletAccount(ctx, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusBadRequest, res)
}

func (h *userHttpHandler) BlockOrUnblockUser(c echo.Context) error {
	log.Println("BlockOrUnblockUser http handler")

	ctx := context.Background()

	// userId := c.Get("user_id").(string)

	userId := c.Param("user_id")
	log.Println("userId", userId)

	res, err := h.userUsecase.BlockOrUnblockUser(ctx, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	user, _ := h.userUsecase.FindOneUserProfile(ctx, userId)

	return response.SuccessResponse(c, http.StatusOK, map[string]any{
		"message": fmt.Sprintf("user_id: %s is successfully changed to: %v", userId, res),
		"user":    user,
	})
}

// --------------- Wish List --------------- //
func (h *userHttpHandler) AddToWishList(c echo.Context) error {
	ctx := context.Background()

	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")
	nftId := c.Param("nft_id")

	res, err := h.userUsecase.AddToWishList(ctx, userId, nftId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *userHttpHandler) GetWishList(c echo.Context) error {
	ctx := context.Background()

	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")

	res, err := h.userUsecase.GetWishList(ctx, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusBadRequest, res)
}

func (h *userHttpHandler) RemoveFromWishList(c echo.Context) error {
	ctx := context.Background()

	userId := strings.TrimPrefix(c.Get("user_id").(string), "user:")
	nftId := c.Param("nft_id")

	err := h.userUsecase.RemoveFromWishList(ctx, userId, nftId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, string("NFT removed from wish list successfully"))
}

// --------------- Address --------------- //
func (h *userHttpHandler) AddAddress(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(user.CreateUserAddressReq)
	UserId := strings.TrimPrefix(c.Get("user_id").(string), "user:")

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.userUsecase.AddAddress(ctx, UserId, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *userHttpHandler) GetAddress(c echo.Context) error {
	ctx := context.Background()

	UserId := strings.TrimPrefix(c.Get("user_id").(string), "user:")

	res, err := h.userUsecase.GetAddress(ctx, UserId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusBadRequest, res)
}

func (h *userHttpHandler) UpdateAddress(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	AddressId := c.Param("address_id")

	if AddressId == "" || AddressId == ":address_id" {
		return response.ErrResponse(c, http.StatusBadRequest, errors.New("address_id cannot be empty").Error())
	}

	if len(AddressId) != 24 {
		return response.ErrResponse(c, http.StatusBadRequest, errors.New("address_id is invalid").Error())
	}

	req := new(user.CreateUserAddressReq)
	UserId := strings.TrimPrefix(c.Get("user_id").(string), "user:")

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.userUsecase.UpdateAddress(ctx, UserId, AddressId, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *userHttpHandler) DeleteAddress(c echo.Context) error {
	ctx := context.Background()

	UserId := strings.TrimPrefix(c.Get("user_id").(string), "user:")
	addressId := c.Param("address_id")

	if addressId == "" || addressId == ":id" {
		return response.ErrResponse(c, http.StatusBadRequest, errors.New("address_id cannot be empty").Error())
	}

	if len(addressId) != 24 {
		return response.ErrResponse(c, http.StatusBadRequest, errors.New("address_id is invalid").Error())
	}

	err := h.userUsecase.DeleteAddress(ctx, UserId, addressId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Address deleted successfully")
}

func (h *userHttpHandler) RazorPayLoad(c echo.Context) error {
	// Load razor pay html page to browser from /modules/templates/payment.html
	tmpl, err := template.ParseFiles("././modules/templates/payment.html")
	if err != nil {
		return err
	}

	// Set Content-Type header
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	// Define data to be passed to the template
	data := struct {
		Name string
	}{
		Name: "RazorPay Payment",
	}

	// Render the template with the data in the browser
	err = tmpl.Execute(c.Response().Writer, data)
	if err != nil {
		return err
	}

	return nil
}

func PaymentHandler(amount int, orderId, razorPayKey, razorPaySecret string) (string, error) {
	client := razorpay.NewClient(razorPayKey, razorPaySecret)

	// fmt.Println("key : ", razorPayKey)
	// fmt.Println("secret : ", razorPaySecret)

	orderParams := map[string]interface{}{
		"amount":   amount * 100,
		"currency": "INR",
		"receipt":  orderId,
	}
	order, err := client.Order.Create(orderParams, nil)
	if err != nil {
		return "", errors.New("PAYMENT NOT INITIATED")
	}

	razorId, _ := order["id"].(string)
	return razorId, nil
}

func (h *userHttpHandler) RazorPaymentConfirm(c echo.Context) error {

	// Get the payment details from the frontend
	// Retrieve the data sent to the backend
	type PaymentDetails struct {
		OrderID   string `json:"order_id"`
		PaymentID string `json:"payment_id"`
		Signature string `json:"signature"`
	}

	var paymentDetails PaymentDetails
	if err := c.Bind(&paymentDetails); err != nil {
		return err
	}

	// Access the data
	orderID := paymentDetails.OrderID
	paymentID := paymentDetails.PaymentID
	signature := paymentDetails.Signature

	// fmt.Println("orderID: ", orderID)
	// fmt.Println("paymentID: ", paymentID)

	secret := h.cfg.Razorpay.Secret

	//============== verify the signature ================
	err := RazorPaymentVerification(signature, orderID, paymentID, secret)
	if err != nil {
		return errors.New("PAYMENT FAILED")
	}

	// update users_transaction document with order_success = paymentId
	ctx := context.Background()
	err = h.userUsecase.UpdateUserTransaction(ctx, orderID, paymentID)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	// fmt.Println("PAYMENT SUCCESS")
	return response.SuccessResponse(c, http.StatusOK, "PAYMENT SUCCESS")
}

func RazorPaymentVerification(sign, orderId, paymentId, secret string) error {
	signature := sign
	data := orderId + "|" + paymentId
	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(signature)) != 1 {
		return errors.New("PAYMENT FAILED")
	} else {
		return nil
	}
}

func (h *userHttpHandler) SalesReport(c echo.Context) error {
	ctx := context.Background()

	res, err := h.userUsecase.SalesReport(ctx)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
