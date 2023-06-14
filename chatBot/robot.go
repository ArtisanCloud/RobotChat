package chatRobot

import (
	"github.com/ArtisanCloud/RobotChat/chatRobot/contract"
)

type ChatRobot struct {
	Client contract.ClientInterface
}

func NewChatRobot(client contract.ClientInterface) (*ChatRobot, error) {

	return &ChatRobot{
		Client: client,
	}, nil
}
