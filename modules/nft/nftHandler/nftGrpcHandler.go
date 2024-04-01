package nftHandler

import "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftUsecase"

type (
	nftGrpcHandler struct {
		nftUsecase nftUsecase.NftUsecaseService
	}
)

func NewNftGrpcHandler(nftUsecase nftUsecase.NftUsecaseService) *nftGrpcHandler {
	return &nftGrpcHandler{
		nftUsecase: nftUsecase,
	}
}
