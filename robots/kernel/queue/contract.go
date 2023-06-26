package queue

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
)

type QueueInterface interface {

	// Connect 建立连接
	Connect(ctx context.Context) error

	// IsConnected 确认连接
	IsConnected(ctx context.Context) bool

	// ProduceMessage 生产消息
	ProduceMessage(ctx context.Context, job model.Job) error

	// ConsumeMessage 消费消息
	ConsumeMessage(ctx context.Context) (interface{}, error)

	// QueueLength 获取队列长度
	QueueLength(ctx context.Context) (int, error)

	// Close 关闭队列连接
	Close(ctx context.Context) error
}
