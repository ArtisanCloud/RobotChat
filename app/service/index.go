package service

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
)

func InitService(config *rcconfig.RCConfig) error {
	// 生成启动ArtBot
	Michelle = NewArtBotService(config)
	if Michelle == nil {
		return errors.New("init ArtBot Service failed")
	}

	err := Michelle.Launch(context.Background())
	if err != nil {
		return err
	}

	// 生成启动ChatBot
	Joy = NewChatBotService(config)
	if Joy == nil {
		return errors.New("init ChatBot Service failed")
	}
	err = Joy.Launch(context.Background())
	if err != nil {
		return err
	}

	return nil
}
