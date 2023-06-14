package model

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"time"
)

type Message struct {
	ID        string
	Author    Author
	Content   Content
	Metadata  *object.HashMap
	Timestamp time.Time
}

type Content struct {
	ContentType string
	Parts       []string
}

type Author struct {
	Role Role
}

type Role string

const (
	UserRole   Role = "user"
	AdminRole  Role = "admin"
	SystemRole Role = "system"
)

func (c *Conversation) AddMessage(sessionIndex int, authorRole Role, contentParts []string) (message Message, index int) {
	message = Message{
		ID:        GenerateId(),
		Author:    Author{Role: authorRole},
		Content:   Content{ContentType: "text", Parts: contentParts},
		Metadata:  &object.HashMap{},
		Timestamp: time.Now(),
	}
	c.Sessions[sessionIndex].Messages = append(c.Sessions[sessionIndex].Messages, message)
	index = len(c.Sessions) - 1 // 返回新 Session 的索引
	return
}
