package nftHandler

import (
	"context"

	nftPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftPb"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftUsecase"
)

type (
	nftGrpcHandler struct {
		nftUsecase nftUsecase.NftUsecaseService
		nftPb.UnimplementedNftGrpcServiceServer
	}
)

func NewNftGrpcHandler(nftUsecase nftUsecase.NftUsecaseService) *nftGrpcHandler {
	return &nftGrpcHandler{
		nftUsecase: nftUsecase,
	}
}

func (g *nftGrpcHandler) FindNftsInIds(ctx context.Context, req *nftPb.FindNftsInIdsReq) (*nftPb.FindNftsInIdsRes, error) {
	return g.nftUsecase.FindNftsInIds(ctx, req)
}

func (g *nftGrpcHandler) AddNftWishlist(ctx context.Context, req *nftPb.AddNftWishlistReq) (*nftPb.AddNftWishlistRes, error) {
	return g.nftUsecase.AddNftWishlist(ctx, req)
}
