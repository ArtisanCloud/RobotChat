package trainerBot

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/pkg/dataformat"
	"github.com/ArtisanCloud/RobotChat/pkg/objectx"
	"mime/multipart"
)

type TrainerBot struct {
	glmEfficientTune *dataformat.GLMEfficientTune
}

func NewTrainerBot() *TrainerBot {
	return &TrainerBot{
		glmEfficientTune: dataformat.NewGLMEfficientTune(),
	}
}

func (bot *TrainerBot) ConvertExcelToSelfCognitionData(ctx context.Context, excelFileData *multipart.File, savedPath string) error {

	convertedDate, err := bot.glmEfficientTune.ConvertExcelToSelfCognitionData(ctx, excelFileData)
	if err != nil {
		return err
	}
	return objectx.SaveObjectToPath(convertedDate, savedPath)
}
