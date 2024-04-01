package nftUsecase

import "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftRepository"

type (
	NftUsecaseService interface{}

	nftUsecase struct {
		nftRepository nftRepository.NftRepositoryService
	}
)

func NewNftUsecase(nftRepository nftRepository.NftRepositoryService) NftUsecaseService {
	return &nftUsecase{nftRepository}
}
