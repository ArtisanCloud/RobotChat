package contract

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
)

// RobotInterface 定义了机器人对话的接口
type RobotInterface interface {
	Send(ctx context.Context, message *model.Message, middlewares ...model.HandleMessageMiddleware) error
	Receive(ctx context.Context, middlewares ...model.HandleMessageMiddleware) (*model.Message, error)
}
