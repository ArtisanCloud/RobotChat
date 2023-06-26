package model

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"gorm.io/datatypes"
	"time"
)

type Conversation struct {
	ID              string
	UserID          string
	Action          string
	Sessions        []*Session
	ParentMessageID string
	Model           string
	StartTime       time.Time
	EndTime         time.Time
	Status          ConversationStatus
}

type ConversationStatus string

const (
	ConversationStatusActive   ConversationStatus = "active"
	ConversationStatusInactive ConversationStatus = "inactive"
	ConversationStatusClosed   ConversationStatus = "closed"
)

func NewConversation(userId string) *Conversation {
	return &Conversation{
		ID:        GenerateId(),
		UserID:    userId,
		Sessions:  []*Session{},
		StartTime: time.Now(),
		Status:    "active",
	}
}

func (c *Conversation) IsActive() bool {
	return c.Status == ConversationStatusActive
}

func (c *Conversation) AddMessage(sessionIndex int, msgType MessageType, authorRole Role, contentParts []string) (message *Message, index int) {
	message = NewMessage(msgType)

	message.Author = string(authorRole)
	content, _ := object.JsonEncode(Content{ContentType: "text", Parts: contentParts})
	message.Content = datatypes.JSON(content)
	meta, _ := object.JsonEncode(object.HashMap{"sessionId": sessionIndex})
	message.Metadata = datatypes.JSON(meta)

	c.Sessions[sessionIndex].Messages = append(c.Sessions[sessionIndex].Messages, message)
	index = len(c.Sessions) - 1 // 返回新 Session 的索引
	return
}
