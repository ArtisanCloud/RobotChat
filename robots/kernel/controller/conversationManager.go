package controller

import (
	model2 "github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"time"
)

type ConversationManager struct {
	Conversations []*model2.Conversation
}

func NewConversationManager() *ConversationManager {
	return &ConversationManager{
		Conversations: []*model2.Conversation{},
	}
}

func (cm *ConversationManager) CreateConversation(userId string) *model2.Conversation {
	conversation := &model2.Conversation{
		ID:        model2.GenerateId(),
		UserID:    userId,
		Sessions:  []*model2.Session{},
		StartTime: time.Now(),
		Status:    "active",
	}
	cm.Conversations = append(cm.Conversations, conversation)
	return conversation
}

func (cm *ConversationManager) GetConversationByID(id string) *model2.Conversation {
	for _, conv := range cm.Conversations {
		if conv.ID == id {
			return conv
		}
	}
	return nil
}

func (cm *ConversationManager) GetActiveConversations() []*model2.Conversation {
	activeConversations := []*model2.Conversation{}
	for _, conv := range cm.Conversations {
		if conv.IsActive() {
			activeConversations = append(activeConversations, conv)
		}
	}
	return activeConversations
}
