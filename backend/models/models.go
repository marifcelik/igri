package models

import (
	"time"

	"go-chat/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type M struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updated_at,omitempty"`
	DeletedAt time.Time          `json:"deletedAt,omitempty" bson:"deleted_at,omitempty"`
	Version   int                `json:"-" bson:"version,omitempty"`
}

type User struct {
	M        `bson:",inline"`
	Name     string    `json:"name,omitempty" bson:"name,omitempty"`
	Username string    `json:"username,omitempty" bson:"username,omitempty"`
	Password string    `json:"password,omitempty" bson:"password,omitempty"`
	LastSeen time.Time `json:"lastSeen,omitempty" bson:"last_seen,omitempty"`
}

type Conversation struct {
	M            `bson:",inline"`
	Type         enums.ConversationType `json:"type,omitempty" bson:"type,omitempty"`
	Participants []primitive.ObjectID   `json:"participants,omitempty" bson:"participants,omitempty"`
	LastMessage  *Message               `json:"lastMessageID,omitempty" bson:"last_message_id,omitempty"`
}

// TODO add ContentType field for handling text, image, video, etc.
// and create MessageContent struct for it
type Message struct {
	M              `bson:",inline"`
	ConversationID primitive.ObjectID `json:"conversationID,omitempty" bson:"conversation_id,omitempty"`
	SenderID       primitive.ObjectID `json:"senderID,omitempty" bson:"sender_id,omitempty"`
	Content        string             `json:"content,omitempty" bson:"content,omitempty"`
}

type UserConversation struct {
	M              `bson:",inline"`
	UserID         primitive.ObjectID `json:"userID,omitempty" bson:"user_id,omitempty"`
	ConversationID primitive.ObjectID `json:"conversationID,omitempty" bson:"conversation_id,omitempty"`
	UnreadCount    int                `json:"unreadCount,omitempty" bson:"unread_count,omitempty"`
	// LastReadMessageID primitive.ObjectID `json:"lastMessageID,omitempty" bson:"last_message_id,omitempty"`
}
