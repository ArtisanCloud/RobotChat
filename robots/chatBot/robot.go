package chatBot

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/contract"
	queue2 "github.com/ArtisanCloud/RobotChat/robots/kernel/queue"
)

type ChatBot struct {
	Client contract.ClientInterface
	Queue  queue2.QueueInterface
}

func NewChatBot(client contract.ClientInterface) (*ChatBot, error) {
	conf := client.GetConfig()
	if conf == nil {
		return nil, errors.New("config file is nil")
	}

	q, err := queue2.LoadQueueDriver(&conf.Queue)
	if err != nil {
		return nil, err
	}
	isConnected := q.IsConnected(context.Background())
	if !isConnected {
		return nil, errors.New("cannot connect queue driver")
	}

	return &ChatBot{
		Client: client,
		Queue:  q,
	}, nil
}
