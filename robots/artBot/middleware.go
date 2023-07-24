package artBot

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/RobotChat/pkg/objectx"
	model2 "github.com/ArtisanCloud/RobotChat/robots/artBot/model"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
)

func (bot *ArtBot) CheckSwitchModel(ctx context.Context, job *model.Job) (*model.Job, error) {

	reqModelHash := ""
	var err error

	// check if current engine is available and will hang out with
	//resProgress, err := bot.Client.Progress(ctx)
	//if err != nil {
	//	return job, err
	//}
	//if resProgress.State

	// get request message model hash
	if job.Payload.MessageType == model.ImageMessage {
		msgImg2Img := &model2.Image2Image{}
		err = objectx.TransformData(job.Payload.Content, msgImg2Img)
		if err != nil {
			return job, err
		}
		reqModelHash = msgImg2Img.SdModelHash
	} else {
		msgTxt2Img := &model2.Text2Image{}
		err = objectx.TransformData(job.Payload.Content, msgTxt2Img)
		if err != nil {
			return job, err
		}
		reqModelHash = msgTxt2Img.SdModelHash
	}

	if reqModelHash == "" {
		return job, errors.New("no hashed model request")
	}

	// 获取当前SD引擎模型Hash
	res, err := bot.Client.GetOptions(ctx)
	if err != nil {
		return job, err
	}
	if res.Options.SdModelCheckpoint == "" {
		return job, errors.New("current  model hash value invalid")
	}

	//fmt.Dump(res.Options.SdCheckpointHash, reqModelHash)
	shortHash := res.Options.SdCheckpointHash[0:10]
	//fmt.Dump(shortHash, reqModelHash)
	if shortHash != reqModelHash {
		// switch model
		sdModels, err := bot.Client.GetModels(ctx)
		if err != nil {
			return job, err
		}
		sdModel := GetModelNameFromHash(reqModelHash, sdModels)

		reqOptions := &model2.OptionsRequest{
			Options: &model2.Options{
				SdModelCheckpoint: sdModel.ModelName,
				//SdCheckpointHash: reqModelHash,
			},
		}
		//fmt.Dump(reqOptions)
		err = bot.Client.SetOptions(ctx, reqOptions)
		if err != nil {
			return job, err
		}
	}

	return job, nil
}

func GetModelNameFromHash(hash string, models []*model2.ArtBotModel) *model2.ArtBotModel {
	for _, model := range models {
		if hash == model.Hash {
			return model
		}
	}
	return nil
}
