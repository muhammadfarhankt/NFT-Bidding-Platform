package nftHandler

import (
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftUsecase"
)

type (
	NftHttpHandlerService interface{}

	nftHttpHandler struct {
		cfg        *config.Config
		nftUsecase nftUsecase.NftUsecaseService
	}
)

func NewNftHttpHandler(cfg *config.Config, nftUsecase nftUsecase.NftUsecaseService) NftHttpHandlerService {
	return &nftHttpHandler{
		cfg:        cfg,
		nftUsecase: nftUsecase,
	}
}
