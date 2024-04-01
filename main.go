package main

import (
	"context"
	"log"
	"os"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/database"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/server"
)

func main() {
	ctx := context.Background()
	_ = ctx

	// Initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	// print env config
	// log.Println(cfg)

	// Database connection
	db := database.DbConn(ctx, &cfg)
	// log.Println(db)
	defer db.Disconnect(ctx)

	// // Start Server
	server.Start(ctx, &cfg, db)
}
