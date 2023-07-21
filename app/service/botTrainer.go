package service

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/botTrainer"
	"mime/multipart"
)

// Bot Trainer Michael
var Michael *BotTrainerService

type BotTrainerService struct {
	botTrainer *botTrainer.BotTrainer
	config     *rcconfig.RCConfig
}

func NewBotTrainerService(config *rcconfig.RCConfig) (bts *BotTrainerService) {
	return &BotTrainerService{
		botTrainer: botTrainer.NewBotTrainer(),
	}
}

func (srv *BotTrainerService) ConvertExcelToSelfCognitionData(ctx context.Context, excelFileData *multipart.File, savePath string) (err error) {
	return srv.botTrainer.ConvertExcelToSelfCognitionData(ctx, excelFileData, savePath)
}
