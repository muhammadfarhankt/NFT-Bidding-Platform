package userRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	UserRepositoryService interface{}

	userRepository struct {
		db *mongo.Client
	}
)

func NewUserRepository(db *mongo.Client) UserRepositoryService {
	return &userRepository{db: db}
}

func (r *userRepository) userDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("user_db")
}
