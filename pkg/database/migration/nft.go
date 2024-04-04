package migration

import (
	"context"
	"log"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/database"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func nftDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("nft_db")
}

func NftMigrate(pctx context.Context, cfg *config.Config) {
	db := nftDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("nfts")
	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"title", 1}}},
	})
	for _, index := range indexs {
		log.Printf("Index: %s", index)
	}

	documents := func() []any {
		roles := []*nft.Nft{
			{
				Title:       "Pixelated Punk",
				Price:       500,
				Description: "A classic generative art piece from the early days of NFTs.",
				ImageUrl:    "https://i.imgur.com/x2jB3zT.jpg",
				UsageStatus: true,
				AuthorId:    "660a424713a303985211df5b",
				OwnerId:     "660a424713a303985211df5b",
				Category:    "Generative Art",
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
				ListingType: "Fixed Price",
			},
			{
				Title:       "The Abstract Bloom",
				Price:       250,
				Description: "A swirling blend of colors and shapes, evoking a sense of cosmic energy.",
				ImageUrl:    "https://i.imgur.com/yWZQm1a.jpg",
				UsageStatus: false, // Example of an NFT not currently in use
				AuthorId:    "660a424713a303985211df5b",
				OwnerId:     "660a424713a303985211df5b",
				Category:    "Abstract",
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
				ListingType: "Fixed Price",
			},
			{
				Title:       "Cyber Kitty",
				Price:       1200,
				Description: "A rare collectible with unique attributes and a playful personality.",
				ImageUrl:    "https://i.imgur.com/a1o2xRj.png",
				UsageStatus: true,
				AuthorId:    "660a424713a303985211df5c",
				OwnerId:     "660a424713a303985211df5c",
				Category:    "Collectible",
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
				ListingType: "Fixed Price",
			},
			{
				Title:       "MusicVerse Pass",
				Price:       300,
				Description: "Grants exclusive access to a music festival in the metaverse, plus behind-the-scenes content.",
				ImageUrl:    "https://i.imgur.com/h83bC5P.jpg",
				UsageStatus: true,
				AuthorId:    "660a424713a303985211df5d",
				OwnerId:     "660a424713a303985211df5d",
				Category:    "Utility",
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
				ListingType: "Fixed Price",
			},
		}

		docs := make([]any, 0)
		for _, r := range roles {
			docs = append(docs, r)
		}
		return docs
	}()

	results, err := col.InsertMany(pctx, documents, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate nft completed: ", results)
}
