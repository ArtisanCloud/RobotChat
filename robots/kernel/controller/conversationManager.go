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

//func (cm *ConversationManager) AddMessage(sessionIndex int, msgType MessageType, authorRole Role, contentParts []string) (message *Message, index int) {
//	message = NewMessage(msgType)
//
//	message.Author = string(authorRole)
//	content, _ := object.JsonEncode(Content{ContentType: "text", Parts: contentParts})
//	message.Content = datatypes.JSON(content)
//	meta, _ := object.JsonEncode(object.HashMap{"sessionId": sessionIndex})
//	message.Metadata = datatypes.JSON(meta)
//
//	c.Sessions[sessionIndex].Messages = append(c.Sessions[sessionIndex].Messages, message)
//	index = len(c.Sessions) - 1 // 返回新 Session 的索引
//	return
//}
