package grpcConn

import (
	"errors"
	"log"
	"net"

	authPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/auth/authPb"
	inventoryPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/inventory/inventoryPb"
	nftPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftPb"
	userPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userPb"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	GrpcClientFactoryHandler interface {
		Auth() authPb.AuthGrpcServiceClient
		User() userPb.UserGrpcServiceClient
		Nft() nftPb.NftGrpcServiceClient
		Inventory() inventoryPb.InventoryGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}

	grpcAuth struct {
	}
)

func (g *grpcClientFactory) Auth() authPb.AuthGrpcServiceClient {
	return authPb.NewAuthGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) User() userPb.UserGrpcServiceClient {
	return userPb.NewUserGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Nft() nftPb.NftGrpcServiceClient {
	return nftPb.NewNftGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Inventory() inventoryPb.InventoryGrpcServiceClient {
	return inventoryPb.NewInventoryGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrpcClientFactoryHandler, error) {
	opts := make([]grpc.DialOption, 0)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	clientConn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Printf("Error: Grpc client connection failed: %s", err.Error())
		return nil, errors.New("error: grpc client connection failed")
	}

	return &grpcClientFactory{
		client: clientConn,
	}, nil
}

func NewGrpcServer(cfg *config.Jwt, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)

	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Error: Failed to listen: %v", err)
	}

	return grpcServer, lis
}
