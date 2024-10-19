package ws

import (
	"context"
	"go-chat/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type wsRepo struct {
	messages, conversations *mongo.Collection
}

func NewWSRepo(db *mongo.Database) *wsRepo {
	return &wsRepo{
		messages:      db.Collection("messages"),
		conversations: db.Collection("conversations"),
	}
}

func (r *wsRepo) SaveConversation(conversation models.Conversation, ctx context.Context) (primitive.ObjectID, error) {
	result, err := r.conversations.InsertOne(ctx, conversation)
	return result.InsertedID.(primitive.ObjectID), err
}

func (r *wsRepo) SaveMessage(message models.Message, ctx context.Context) error {
	_, err := r.messages.InsertOne(ctx, message)
	return err
}
