package payment

type (
	NftServiceReq struct {
		Nfts []*NftServiceReqDatum `json:"nfts" validate:"required"`
	}

	NftServiceReqDatum struct {
		NftId string  `json:"nft_id" validate:"required,max=64"`
		Price float64 `json:"price"`
	}
)
