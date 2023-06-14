package model

import (
	"github.com/google/uuid"
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

func GenerateId() string {
	return uuid.New().String()
}

func NewConversation(userId string) *Conversation {
	return &Conversation{
		ID:        GenerateId(),
		UserID:    userId,
		Sessions:  []*Session{},
		StartTime: time.Now(),
		Status:    "active",
	}
}

func (mdl Conversation) IsActive() bool {
	return mdl.Status == ConversationStatusActive
}
