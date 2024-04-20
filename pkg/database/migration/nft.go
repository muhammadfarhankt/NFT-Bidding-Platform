package migration

import (
	"context"
	"log"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/nft"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/database"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	id1, _ := primitive.ObjectIDFromHex("662336e3abeca0cd8ca871d3")
	category1, _ := primitive.ObjectIDFromHex("6623401dc5a2de40a6dae06e")

	id2, _ := primitive.ObjectIDFromHex("662336e3abeca0cd8ca871d4")
	category2, _ := primitive.ObjectIDFromHex("6623404ec5a2de40a6dae06f")

	id3, _ := primitive.ObjectIDFromHex("662336e3abeca0cd8ca871d5")
	category3, _ := primitive.ObjectIDFromHex("66234059c5a2de40a6dae070")

	id4, _ := primitive.ObjectIDFromHex("662336e3abeca0cd8ca871d6")

	documents := func() []any {
		roles := []*nft.Nft{
			{
				Title:       "Pixelated Punk",
				Price:       500,
				Description: "A classic generative art piece from the early days of NFTs.",
				ImageUrl:    "https://i.imgur.com/x2jB3zT.jpg",
				UsageStatus: true,
				AuthorId:    id1,
				OwnerId:     id1,
				Category:    category1,
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
				AuthorId:    id2,
				OwnerId:     id2,
				Category:    category2,
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
				AuthorId:    id3,
				OwnerId:     id3,
				Category:    category3,
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
				AuthorId:    id4,
				OwnerId:     id4,
				Category:    category3,
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
