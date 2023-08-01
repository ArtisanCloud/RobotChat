package service

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/trainerBot"
	"mime/multipart"
)

// Bot Trainer Michael
var Michael *TrainerBotService

type TrainerBotService struct {
	trainerBot *trainerBot.TrainerBot
	config     *rcconfig.RCConfig
}

func NewTrainerBotService(config *rcconfig.RCConfig) (bts *TrainerBotService) {
	return &TrainerBotService{
		trainerBot: trainerBot.NewTrainerBot(),
	}
}

func (srv *TrainerBotService) ConvertExcelToSelfCognitionData(ctx context.Context, excelFileData *multipart.File, savePath string) (err error) {
	return srv.trainerBot.ConvertExcelToSelfCognitionData(ctx, excelFileData, savePath)
}
