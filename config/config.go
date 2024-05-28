package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App      App
		Db       Db
		Jwt      Jwt
		Kafka    Kafka
		Grpc     Grpc
		Paginate Paginate
		Email    Email
		Razorpay Razorpay
	}

	App struct {
		Name  string
		Url   string
		Stage string
	}

	Db struct {
		Url string
	}

	Jwt struct {
		AccessSecretKey  string
		RefreshSecretKey string
		ApiSecretKey     string
		AccessDuration   int64
		RefreshDuration  int64
	}

	Kafka struct {
		Url    string
		ApiKey string
		Secret string
	}

	Grpc struct {
		AuthUrl    string
		InventUrl  string
		NftUrl     string
		PaymentUrl string
		UserUrl    string
	}

	Paginate struct {
		NftNextPageBasedUrl       string
		InventoryNextPageBasedUrl string
	}

	Email struct {
		Username string
		Password string
	}

	Razorpay struct {
		Key    string
		Secret string
	}
)

func LoadConfig(path string) Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return Config{
		App: App{
			Name:  os.Getenv("APP_NAME"),
			Url:   os.Getenv("APP_URL"),
			Stage: os.Getenv("APP_STAGE"),
		},
		Db: Db{
			Url: os.Getenv("DB_URL"),
		},
		Jwt: Jwt{
			AccessSecretKey:  os.Getenv("JWT_ACCESS_SECRET_KEY"),
			RefreshSecretKey: os.Getenv("JWT_REFRESH_SECRET_KEY"),
			ApiSecretKey:     os.Getenv("JWT_API_SECRET_KEY"),
			AccessDuration: func() int64 {
				accessDuration, err := strconv.ParseInt(os.Getenv("JWT_ACCESS_DURATION"), 10, 64)
				if err != nil {
					log.Fatalf("Error loading ACCESS_DURATION: %v", err)
				}
				return accessDuration
			}(),
			RefreshDuration: func() int64 {
				refreshDuration, err := strconv.ParseInt(os.Getenv("JWT_REFRESH_DURATION"), 10, 64)
				if err != nil {
					log.Fatalf("Error loading REFRESH_DURATION: %v", err)
				}
				return refreshDuration
			}(),
		},
		Kafka: Kafka{
			Url:    os.Getenv("KAFKA_URL"),
			ApiKey: os.Getenv("KAFKA_API_KEY"),
			Secret: os.Getenv("KAFKA_SECRET"),
		},
		Grpc: Grpc{
			AuthUrl:    os.Getenv("GRPC_AUTH_URL"),
			InventUrl:  os.Getenv("GRPC_INVENTORY_URL"),
			NftUrl:     os.Getenv("GRPC_NFT_URL"),
			PaymentUrl: os.Getenv("GRPC_PAYMENT_URL"),
			UserUrl:    os.Getenv("GRPC_USER_URL"),
		},
		Paginate: Paginate{
			NftNextPageBasedUrl:       os.Getenv("PAGINATE_NFT_NEXT_PAGE_BASED_URL"),
			InventoryNextPageBasedUrl: os.Getenv("PAGINATE_INVENTORY_NEXT_PAGE_BASED_URL"),
		},
		Email: Email{
			Username: os.Getenv("APP_EMAIL"),
			Password: os.Getenv("APP_PASSWORD"),
		},
		Razorpay: Razorpay{
			Key:    os.Getenv("RAZOR_PAY_KEY"),
			Secret: os.Getenv("RAZOR_PAY_SECRET"),
		},
	}
}
