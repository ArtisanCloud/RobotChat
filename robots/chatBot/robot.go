package chatBot

import (
	"github.com/ArtisanCloud/RobotChat/kernel/controller"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/contract"
)

type ChatBot struct {
	Client              contract.ClientInterface
	ConversationManager *controller.ConversationManager
}

func NewChatBot(client contract.ClientInterface) (*ChatBot, error) {

	return &ChatBot{
		Client:              client,
		ConversationManager: controller.NewConversationManager(),
	}, nil
}
