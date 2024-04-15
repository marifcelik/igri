package ws

import (
	"context"
	"go-chat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type wsRepo struct {
	users    *mongo.Collection
	messages *mongo.Collection
}

func NewWSRepo(db *mongo.Database) *wsRepo {
	return &wsRepo{
		users:    db.Collection("users"),
		messages: db.Collection("messages"),
	}
}

func (r *wsRepo) GetUserByUsername(username string, ctx context.Context) (models.User, error) {
	var user models.User
	result := r.users.FindOne(ctx, bson.M{"username": username})
	if result.Err() != nil {
		return user, result.Err()
	}
	err := result.Decode(&user)
	return user, err
}
