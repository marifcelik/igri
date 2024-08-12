package ws

import (
	"context"
	"go-chat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type wsRepo struct {
	messages *mongo.Collection
}

func NewWSRepo(db *mongo.Database) *wsRepo {
	return &wsRepo{messages: db.Collection("messages")}
}

/* TODO research the use of generic methods to write/read user and group messages
without seperate methods */

func (r *wsRepo) GetMessages(ctx context.Context, startEnd ...int) ([]models.UserMessage, error) {
	var messages []models.UserMessage
	cursor, err := r.messages.Find(ctx, bson.D{})
	if err != nil {
		return messages, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &messages)
	return messages, err
}

func (r *wsRepo) SaveMessage(message models.UserMessage, ctx context.Context) error {
	_, err := r.messages.InsertOne(ctx, message)
	return err
}

func (r *wsRepo) SaveGroupMessage(message models.GroupMessage, ctx context.Context) error {
	_, err := r.messages.InsertOne(ctx, message)
	return err
}
