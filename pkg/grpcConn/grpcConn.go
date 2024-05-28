package grpcConn

import (
	"context"
	"errors"
	"log"
	"net"

	authPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/auth/authPb"
	inventoryPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/inventory/inventoryPb"
	nftPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft/nftPb"
	userPb "github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/user/userPb"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/jwtAuth"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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
		secretKey string
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

	grpcAuth := &grpcAuth{
		secretKey: cfg.ApiSecretKey,
	}

	opts = append(opts, grpc.UnaryInterceptor(grpcAuth.unaryAuthorization))

	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Error: Failed to listen: %v", err)
	}

	return grpcServer, lis
}

func (g *grpcAuth) unaryAuthorization(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("Error: Metadata not found")
		return nil, errors.New("error: metadata not found")
	}

	authHeader, ok := md["auth"]
	if !ok {
		log.Printf("Error: Metadata not found")
		return nil, errors.New("error: metadata not found")
	}

	if len(authHeader) == 0 {
		log.Printf("Error: Metadata not found")
		return nil, errors.New("error: metadata not found")
	}

	_, err := jwtAuth.ParseToken(g.secretKey, string(authHeader[0]))
	if err != nil {
		log.Printf("Error: Parse token failed: %s", err.Error())
		return nil, errors.New("error: token is invalid")
	}
	// log.Printf("claims: %v", claims)

	return handler(ctx, req)
}
