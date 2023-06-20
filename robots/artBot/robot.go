package artBot

import (
	"github.com/ArtisanCloud/RobotChat/kernel/controller"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/contract"
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
