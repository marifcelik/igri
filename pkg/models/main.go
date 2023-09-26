package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string  `gorm:"primaryKey;type:varchar(36) not null" json:"id"`
	Name     string  `gorm:"type:varchar(255)" json:"name"`
	Username string  `gorm:"type:varchar(255);unique" json:"username"`
	Password string  `json:"password"`
	Groups   []Group `gorm:"many2many:user_groups"`
}

type UserMessage struct {
	gorm.Model
	ID         string `gorm:"primaryKey;type:varchar(36) not null"`
	Text       string `gorm:"type:text not null"`
	Seen       bool   `gorm:"default:false"`
	SenderID   string `gorm:"type:varchar(36) not null"`
	Sender     User
	ReceiverID string `gorm:"type:varchar(36) not null"`
	Receiver   User
}

type Group struct {
	gorm.Model
	ID    string `gorm:"primaryKey;type:varchar(36) not null"`
	Name  string `gorm:"type:varchar(255)"`
	Users []User `gorm:"many2many:user_groups"`
}

type GroupMessage struct {
	gorm.Model
	ID       string `gorm:"primaryKey;type:varchar(36) not null"`
	Text     string `gorm:"type:text not null"`
	SenderID string `gorm:"type:varchar(36) not null"`
	Sender   User
	GroupID  string `gorm:"type:varchar(36) not null"`
	Group    Group
}

type GroupMessageSeenBy struct {
	gorm.Model
	ID             string `gorm:"primaryKey;type:varchar(36) not null"`
	GroupMessageID string `gorm:"type:varchar(36) not null index"`
	GroupMessage   GroupMessage
	UserID         string `gorm:"type:varchar(36) not null index"`
	User           User
	SeenAt         time.Time
}
