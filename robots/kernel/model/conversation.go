package model

import (
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
