package controller

import (
	"github.com/ArtisanCloud/RobotChat/kernel/model"
	"time"
)

type ConversationManager struct {
	Conversations []*model.Conversation
}

func NewConversationManager() *ConversationManager {
	return &ConversationManager{
		Conversations: []*model.Conversation{},
	}
}

func (cm *ConversationManager) CreateConversation(userId string) *model.Conversation {
	conversation := &model.Conversation{
		ID:        model.GenerateId(),
		UserID:    userId,
		Sessions:  []*model.Session{},
		StartTime: time.Now(),
		Status:    "active",
	}
	cm.Conversations = append(cm.Conversations, conversation)
	return conversation
}

func (cm *ConversationManager) GetConversationByID(id string) *model.Conversation {
	for _, conv := range cm.Conversations {
		if conv.ID == id {
			return conv
		}
	}
	return nil
}

func (cm *ConversationManager) GetActiveConversations() []*model.Conversation {
	activeConversations := []*model.Conversation{}
	for _, conv := range cm.Conversations {
		if conv.IsActive() {
			activeConversations = append(activeConversations, conv)
		}
	}
	return activeConversations
}
