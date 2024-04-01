package paymentHandler

import (
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/payment/paymentUsecase"
)

type (
	PaymentQueueHandlerService interface{}

	paymentQueueHandler struct {
		cfg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentQueueHandlerService {
	return &paymentQueueHandler{
		cfg:            cfg,
		paymentUsecase: paymentUsecase,
	}
}
