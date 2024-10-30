package ws

import (
	"context"
	"time"

	"go-chat/config"
	"go-chat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type wsRepo struct {
	messages, conversations, userConversations *mongo.Collection
}

func NewWSRepo(db *mongo.Database) *wsRepo {
	return &wsRepo{
		messages:          db.Collection(config.C.DBKey.Messages),
		conversations:     db.Collection(config.C.DBKey.Conversations),
		userConversations: db.Collection(config.C.DBKey.UserConversations),
	}
}

func (r *wsRepo) CreateMessage(message models.Message, ctx context.Context) (primitive.ObjectID, error) {
	result, err := r.messages.InsertOne(ctx, message)
	return result.InsertedID.(primitive.ObjectID), err
}

func (r *wsRepo) CreateConversation(conversation models.Conversation, ctx context.Context) (primitive.ObjectID, error) {
	result, err := r.conversations.InsertOne(ctx, conversation)
	return result.InsertedID.(primitive.ObjectID), err
}

func (r *wsRepo) UpdateLastMessageOfConversation(message models.Message, ctx context.Context) error {
	_, err := r.conversations.UpdateOne(ctx,
		bson.M{"_id": message.ConversationID},
		bson.M{
			"$set": bson.M{"last_message": message, "updated_at": time.Now()},
		},
	)
	return err
}

func (r *wsRepo) GetConversationsByUserID(userID primitive.ObjectID, ctx context.Context) ([]models.Conversation, error) {
	opt := options.Find().SetSort(bson.M{"last_message.created_at": -1})

	cursor, err := r.conversations.Find(ctx, bson.M{"participants": bson.M{"$in": []primitive.ObjectID{userID}}}, opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var conversations []models.Conversation
	if err = cursor.All(ctx, &conversations); err != nil {
		return nil, err
	}
	return conversations, nil
}

func (r *wsRepo) GetRecipientIDByConversationID(conversationID primitive.ObjectID, userID primitive.ObjectID, ctx context.Context) (primitive.ObjectID, error) {
	log.Info("get recipient id by conversation id", "conversation_id", conversationID, "user_id", userID)
	var conversation models.Conversation
	err := r.conversations.FindOne(ctx, bson.M{"_id": conversationID, "participants": userID}).Decode(&conversation)
	if err != nil {
		return primitive.NilObjectID, err
	}

	recipiendID := conversation.Participants[0]
	if recipiendID.Hex() == userID.Hex() {
		recipiendID = conversation.Participants[1]
	}
	return recipiendID, nil
}
