package chatBot

import (
	"github.com/ArtisanCloud/RobotChat/chatBot/contract"
	"github.com/ArtisanCloud/RobotChat/kernel/controller"
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
