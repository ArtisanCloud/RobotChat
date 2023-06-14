package model

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"time"
)

type Session struct {
	ID         string
	ctx        context.Context
	Messages   []Message
	Metadata   *object.HashMap
	StartTime  time.Time
	EndTime    time.Time
	IsComplete bool
}

func (c *Conversation) AddSession(ctx context.Context) (session *Session, index int) {
	session = &Session{
		ID:         GenerateId(),
		Messages:   []Message{},
		Metadata:   &object.HashMap{},
		StartTime:  time.Now(),
		IsComplete: false,
	}
	c.Sessions = append(c.Sessions, session)
	index = len(c.Sessions) - 1 // 返回新 Session 的索引
	return
}

func (mdl *Session) SetContext(ctx context.Context) {
	mdl.ctx = ctx
}

func (mdl *Session) GetContext() context.Context {
	return mdl.ctx
}
