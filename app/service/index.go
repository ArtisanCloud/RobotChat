package service

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
)

func InitService(config *rcconfig.RCConfig) error {
	// 生成启动ArtBot
	SrvArtBot = NewArtBotService(config)
	if SrvArtBot == nil {
		return errors.New("init ArtBot Service failed")
	}

	err := SrvArtBot.Start(context.Background())
	if err != nil {
		return err
	}

	// 生成启动ChatBot
	SrvChatBot = NewChatBotService(config)
	if SrvChatBot == nil {
		return errors.New("init ChatBot Service failed")
	}
	err = SrvChatBot.Start(context.Background())
	if err != nil {
		return err
	}

	return nil
}
