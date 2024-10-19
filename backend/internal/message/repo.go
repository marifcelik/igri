package message

import (
	"context"
	"time"

	"go-chat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageQuery struct {
	UserID     primitive.ObjectID
	Start, End time.Time
	Limit      int64 // options.Find().SetLimit requires int64
}

type messageRepo struct {
	messages, userConversations *mongo.Collection
}

func NewMessageRepo(db *mongo.Database) *messageRepo {
	return &messageRepo{
		messages:          db.Collection("messages"),
		userConversations: db.Collection("userConversations"),
	}
}

func (r *messageRepo) GetMessageByID(ctx context.Context, id primitive.ObjectID) (models.Message, error) {
	var message models.Message
	err := r.userConversations.FindOne(ctx, bson.M{"_id": id}).Decode(&message)
	return message, err
}

func (r *messageRepo) GetUserConversations(ctx context.Context, query MessageQuery) ([]models.UserConversation, error) {
	options := options.Find()
	options.SetLimit(10)
	if query.Limit >= 0 {
		options.SetLimit(query.Limit)
	}
	options.SetSort(bson.M{"updatedAt": -1})

	cursor, err := r.userConversations.Find(ctx, bson.M{"userID": query.UserID}, options)
	if err != nil {
		return nil, err
	}

	var messages []models.UserConversation
	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *messageRepo) GetUserMessages(ctx context.Context, query MessageQuery) ([]models.UserConversation, error) {
	cursor, err := r.userConversations.Find(ctx, bson.M{"userID": query.UserID})
	if err != nil {
		return nil, err
	}

	var messages []models.UserConversation
	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
