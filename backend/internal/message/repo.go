package message

import (
	"context"
	"slices"
	"time"

	"go-chat/config"
	"go-chat/enums"
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
	users, messages, conversations *mongo.Collection
}

func NewMessageRepo(db *mongo.Database) *messageRepo {
	return &messageRepo{
		users:         db.Collection(config.C.DBKey.Users),
		messages:      db.Collection(config.C.DBKey.Messages),
		conversations: db.Collection(config.C.DBKey.Conversations),
	}
}

func (r *messageRepo) GetMessageByID(ctx context.Context, id primitive.ObjectID) (models.Message, error) {
	var message models.Message
	err := r.messages.FindOne(ctx, bson.M{"_id": id}).Decode(&message)
	return message, err
}

// TODO add pagination and query options
func (r *messageRepo) GetUserConversations(ctx context.Context, userID primitive.ObjectID) ([]models.Conversation, error) {
	opt := options.Find().SetSort(bson.M{"last_message.created_at": -1})

	cursor, err := r.conversations.Find(ctx, bson.M{"participants": userID}, opt)
	if err != nil {
		return nil, err
	}

	var conversations []models.Conversation
	for cursor.Next(ctx) {
		var conversation models.Conversation
		if err := cursor.Decode(&conversation); err != nil {
			return nil, err
		}

		if conversation.Type == enums.NormalConversation {
			otherParticipantID := getOtherParticipantID(conversation.Participants, userID)
			otherParticipantName, err := r.getParticipantNameByID(ctx, otherParticipantID)
			if err != nil {
				return nil, err
			}
			conversation.Name = otherParticipantName
		}

		conversations = append(conversations, conversation)
	}

	return conversations, nil
}

func (r *messageRepo) GetUserMessages(ctx context.Context, query MessageQuery) ([]models.UserConversation, error) {
	cursor, err := r.conversations.Find(ctx, bson.M{"userID": query.UserID})
	if err != nil {
		return nil, err
	}

	var messages []models.UserConversation
	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *messageRepo) GetConversationMessages(ctx context.Context, conversationID primitive.ObjectID) ([]models.Message, error) {
	opt := options.Find().SetSort(bson.M{"_id": -1}).SetLimit(15)
	cursor, err := r.messages.Find(ctx, bson.M{"conversation_id": conversationID}, opt)
	if err != nil {
		return nil, err
	}

	var messages []models.Message
	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	slices.Reverse(messages)

	return messages, nil
}

func (r *messageRepo) getParticipantNameByID(ctx context.Context, userID primitive.ObjectID) (string, error) {
	var user models.User
	err := r.users.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}

func getOtherParticipantID(participants []primitive.ObjectID, userID primitive.ObjectID) primitive.ObjectID {
	for _, p := range participants {
		if p != userID {
			return p
		}
	}
	return primitive.NilObjectID
}
