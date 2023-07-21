package botTrainer

import (
	"context"
	"github.com/ArtisanCloud/RobotChat/pkg/dataformat"
	"github.com/ArtisanCloud/RobotChat/pkg/objectx"
	"mime/multipart"
)

type BotTrainer struct {
	glmEfficientTune *dataformat.GLMEfficientTune
}

func NewBotTrainer() *BotTrainer {
	return &BotTrainer{
		glmEfficientTune: dataformat.NewGLMEfficientTune(),
	}
}

func (bot *BotTrainer) ConvertExcelToSelfCognitionData(ctx context.Context, excelFileData *multipart.File, savedPath string) error {

	convertedDate, err := bot.glmEfficientTune.ConvertExcelToSelfCognitionData(ctx, excelFileData)
	if err != nil {
		return err
	}
	return objectx.SaveObjectToPath(convertedDate, savedPath)
}
