package message

import (
	"context"
	"errors"
	"go-chat/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserMessageQuery struct {
	ReceiverId primitive.ObjectID
	Start, End time.Time
}

type messageRepo struct {
	userMessages, groupMessages *mongo.Collection
}

func NewMessageRepo(db *mongo.Database) *messageRepo {
	return &messageRepo{
		userMessages:  db.Collection("messages"),
		groupMessages: db.Collection("group_messages"),
	}
}

// GetUserMessages returns 25 messages for a user based on the query options
func (r *messageRepo) GetUserMessages(ctx context.Context, queryOptions UserMessageQuery) ([]models.UserMessage, error) {
	// TODO add additional query options like limit, offset, etc

	var cursor *mongo.Cursor
	var err error

	if _, err = primitive.ObjectIDFromHex(queryOptions.ReceiverId.Hex()); err != nil {
		return nil, err
	}
	if queryOptions.ReceiverId.IsZero() {
		return nil, errors.New("receiver_id is required")
	}

	query := bson.D{{Key: "receiver_id", Value: queryOptions.ReceiverId}}
	if !(queryOptions.Start.IsZero() || queryOptions.End.IsZero()) {
		query = append(query, bson.D{
			{Key: "created_at", Value: bson.D{
				{Key: "$gte", Value: queryOptions.Start},
				{Key: "$lte", Value: queryOptions.End},
			}},
		}...)
	}

	findOptions := options.Find()
	findOptions.SetLimit(25)
	cursor, err = r.userMessages.Find(ctx, query, findOptions)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []models.UserMessage
	if err = cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
