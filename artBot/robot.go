package artBot

import (
	"github.com/ArtisanCloud/RobotChat/artBot/contract"
	"github.com/ArtisanCloud/RobotChat/kernel/controller"
)

type ArtBot struct {
	Client contract.ClientInterface

	ConversationManager *controller.ConversationManager
}

func NewArtBot(client contract.ClientInterface) (*ArtBot, error) {

	return &ArtBot{
		Client:              client,
		ConversationManager: controller.NewConversationManager(),
	}, nil
}
