package ws

import (
	"context"
	"go-chat/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type wsRepo struct {
	userMessages, groupMessages *mongo.Collection
}

func NewWSRepo(db *mongo.Database) *wsRepo {
	return &wsRepo{
		userMessages:  db.Collection("messages"),
		groupMessages: db.Collection("group_messages"),
	}
}

/* TODO research the use of generic methods to write/read user and group messages
without seperate methods */

func (r *wsRepo) SaveMessage(message models.UserMessage, ctx context.Context) error {
	_, err := r.userMessages.InsertOne(ctx, message)
	return err
}

func (r *wsRepo) SaveGroupMessage(message models.GroupMessage, ctx context.Context) error {
	_, err := r.groupMessages.InsertOne(ctx, message)
	return err
}
