package database

import (
	"context"
	"log"
	"time"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// func DbConn(cfg *config.Config) *mongo.Client {
// 	return nil
// }

// pctx - parent context
func DbConn(pctx context.Context, cfg *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.Url))
	if err != nil {
		log.Fatalf("Error: Conntect to database error: %s", err.Error())
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Error: Pinging to database error: %s", err.Error())
	}

	// log.Println("Connected to MongoDB!")
	return client
}
