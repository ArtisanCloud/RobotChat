package chatBot

import (
	"github.com/ArtisanCloud/RobotChat/chatBot/contract"
)

type ChatBot struct {
	Client contract.ClientInterface
}

func NewChatBot(client contract.ClientInterface) (*ChatBot, error) {

	return &ChatBot{
		Client: client,
	}, nil
}
