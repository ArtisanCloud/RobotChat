package artBot

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/contract"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	queue2 "github.com/ArtisanCloud/RobotChat/robots/kernel/queue"
)

type ArtBot struct {
	Client contract.ClientInterface
	Queue  queue2.QueueInterface
}

func NewArtBot(client contract.ClientInterface) (*ArtBot, error) {
	conf := client.GetConfig()
	if conf == nil {
		return nil, errors.New("config file is nil")
	}

	q, err := queue2.LoadQueueDriver(&conf.Queue)
	if err != nil {
		return nil, err
	}
	err = q.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	return &ArtBot{
		Client: client,
		Queue:  q,
	}, nil
}

func (bot *ArtBot) IsAwaken(ctx context.Context) error {

	if bot.Client == nil {
		return errors.New("robot is not existed")
	}

	if !bot.Queue.IsConnected(ctx) {
		return errors.New("queue is not connected")
	}

	return nil
}

func (bot *ArtBot) WakeUp(ctx context.Context) error {

	err := bot.Queue.Connect(ctx)

	return err
}

func (bot *ArtBot) Send(ctx context.Context, message *model.Message, middlewares ...model.HandleMessageMiddleware) error {

	// 将消息传递给中间件处理
	for _, middleware := range middlewares {
		var err error
		// 执行中间件处理逻辑
		message, err = middleware(ctx, message)
		if err != nil {
			// 如果中间件返回错误，可以选择处理错误或直接返回
			return err
		}
	}

	job := model.Job{
		Id:      model.GenerateId(),
		Payload: message,
	}
	err := bot.Queue.ProduceMessage(ctx, job)
	if err != nil {
		return err
	}

	return nil
}
func (bot *ArtBot) Receive(ctx context.Context, middlewares ...model.HandleMessageMiddleware) (*model.Message, error) {
	return nil, nil
}
