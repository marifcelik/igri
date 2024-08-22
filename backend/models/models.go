package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type M struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updated_at,omitempty"`
	DeletedAt time.Time          `json:"deletedAt,omitempty" bson:"deleted_at,omitempty"`
	Version   int                `json:"version,omitempty" bson:"version,omitempty"`
}

type User struct {
	M        `bson:",inline"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

type UserMessage struct {
	M          `bson:",inline"`
	Message    string             `json:"text,omitempty" bson:"text,omitempty"`
	SenderID   primitive.ObjectID `json:"senderId,omitempty" bson:"sender_id,omitempty"`
	ReceiverID primitive.ObjectID `json:"receiverId,omitempty" bson:"receiver_id,omitempty"`
}

type Group struct {
	M     `bson:",inline"`
	Name  string               `json:"name,omitempty" bson:"name,omitempty"`
	Users []primitive.ObjectID `json:"users,omitempty" bson:"users,omitempty"`
}

type GroupMessage struct {
	M        `bson:",inline"`
	Message  string             `json:"text,omitempty" bson:"text,omitempty"`
	GroupID  primitive.ObjectID `json:"groupId,omitempty" bson:"group_id,omitempty"`
	SenderID primitive.ObjectID `json:"senderId,omitempty" bson:"sender_id,omitempty"`
	// SeenByIDs []string           `json:"seenBy,omitempty" bson:"seenBy,omitempty"`
}
