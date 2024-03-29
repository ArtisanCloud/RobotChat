package service

import (
	"context"
	"fmt"
	request2 "github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/app/request/sd"
	"github.com/ArtisanCloud/RobotChat/pkg"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/artBot"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/driver/ArtisanCloud"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/driver/contract"
	model2 "github.com/ArtisanCloud/RobotChat/robots/artBot/model"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/model/controlNet"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/controller"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"github.com/artisancloud/httphelper"
)

// Robot Michelle for Stable Diffusion
var Michelle *ArtBotService

type ArtBotService struct {
	artBot              *artBot.ArtBot
	config              *rcconfig.RCConfig
	conversationManager *controller.ConversationManager
}

func NewArtBotService(config *rcconfig.RCConfig) (abs *ArtBotService) {

	var driver contract.ArtBotClientInterface
	if config.ArtBot.Channel == "" || pkg.Lower(config.ArtBot.Channel) == "stablediffusion" {
		// 使用 Meonako 作为 SD SDK驱动
		//driver = Meonako.NewDriver(&config.ArtBot)
		driver = ArtisanCloud.NewDriver(&config.ArtBot)
	}
	if driver == nil {
		return nil
	}

	robot, err := artBot.NewArtBot(driver)
	if err != nil {
		panic(err)
	}
	robot.NotifyUrl = config.ArtBot.Queue.NotifyUrl

	abs = &ArtBotService{
		artBot: robot,
		config: config,
	}
	return abs
}

func (srv *ArtBotService) IsAwaken(ctx context.Context) error {
	err := srv.artBot.IsAwaken(ctx)
	return err
}

func (srv *ArtBotService) Launch(ctx context.Context) error {

	// 预处理请求消息
	preProcess := func(ctx context.Context, job *model.Job) (*model.Job, error) {
		srv.artBot.Logger.Info(srv.artBot.Name, fmt.Sprintf("I get your message:%s,%v", job.Id, job.Payload.Metadata))
		return job, nil
	}
	srv.artBot.SetMessagePreHandler(preProcess)

	// 错误请求处理
	errHandle := func(errReply *model.ErrReply) {
		srv.artBot.Logger.Error("handle error:", errReply.Job.Id, errReply.Err.Error())
		if errReply.Err != nil {
			errReply.Job.Payload.Metadata.ErrMsg = errReply.Err.Error()
		} else {
			errReply.Job.Payload.Metadata.ErrMsg = "unknown Error from error handle"
		}

		httpClient, err := httphelper.NewRequestHelper(&httphelper.Config{
			BaseUrl: srv.artBot.NotifyUrl,
		})
		if err != nil {
			srv.artBot.Logger.Error(srv.artBot.Name, "handle error new client:", err.Error())
			return
		}

		srv.artBot.Logger.Error(srv.artBot.Name, "handle error post notify url:", srv.artBot.NotifyUrl)

		_, err = httpClient.Df().WithContext(ctx).Method("POST").Json(errReply.Job).Request()
		if err != nil {
			srv.artBot.Logger.Error(srv.artBot.Name, "handle error request webhook error:", err.Error())
			return
		}

		return

	}
	srv.artBot.SetErrorHandler(errHandle)

	// 队列回调，处理是否需要切换模型
	queuePostCheckModelHandle := srv.artBot.CheckSwitchModel

	// 队列回调请求
	queuePostJobHandle := func(ctx context.Context, job *model.Job) (*model.Job, error) {

		srv.artBot.Logger.Info("queue has process your request:", job.Id, job.Payload.Content)
		var (
			err     error
			message *model.Message
		)
		if job.Payload.MessageType == model.ImageMessage {
			message, err = srv.artBot.Client.Image2Image(ctx, job.Payload)
		} else {

			message, err = srv.artBot.Client.Text2Image(ctx, job.Payload)
		}

		if err != nil {
			return job, err
		}
		job.Payload = message

		return job, nil
	}

	queuePostWebhook := func(ctx context.Context, job *model.Job) (*model.Job, error) {

		httpClient, err := httphelper.NewRequestHelper(&httphelper.Config{
			BaseUrl: srv.artBot.NotifyUrl,
		})
		if err != nil {
			srv.artBot.Logger.Error(srv.artBot.Name, "webhook:", err.Error())
			return job, err
		}

		srv.artBot.Logger.Info(srv.artBot.Name, "post url:", srv.artBot.NotifyUrl)
		_, err = httpClient.Df().WithContext(ctx).Method("POST").Json(job).Request()
		if err != nil {
			srv.artBot.Logger.Error(srv.artBot.Name, "webhook:", err.Error())
			return job, err
		}

		return job, nil
	}

	srv.artBot.SetPostMessageHandler(queuePostCheckModelHandle, queuePostJobHandle, queuePostWebhook)

	// 启动机器人
	err := srv.artBot.Start(ctx)

	return err
}

func (srv *ArtBotService) Txt2Image(ctx context.Context, req *model2.Text2ImageRequest) (res *model2.Image2ImageResponse, err error) {

	message, err := srv.artBot.CreateTextMessage(req)
	if err != nil {
		return nil, err
	}

	res, err = srv.artBot.SendAndWait(ctx, message, srv.artBot.Client.Text2Image)

	return res, err
}

func (srv *ArtBotService) Image2Image(ctx context.Context, req *model2.Image2ImageRequest) (res *model2.Image2ImageResponse, err error) {

	message, err := srv.artBot.CreateImageMessage(req)
	if err != nil {
		return nil, err
	}

	res, err = srv.artBot.SendAndWait(ctx, message, srv.artBot.Client.Image2Image)

	return res, err
}

func (srv *ArtBotService) ChatTxt2Image(ctx context.Context, req *sd.ParaText2Image) (job *model.Job, err error) {

	//conversation := srv.conversationManager.GetConversationByID(req.ConversationId)
	//conversation.GetSessionById[req.SessionId]

	message, err := srv.artBot.CreateTextMessage(req)
	if err != nil {
		return nil, err
	}

	job, err = srv.artBot.Send(ctx, message)

	return job, err
}

func (srv *ArtBotService) ChatImage2Image(ctx context.Context, req *sd.ParaImage2Image) (job *model.Job, err error) {

	//conversation := srv.conversationManager.GetConversationByID(req.ConversationId)
	//conversation.GetSessionById[req.SessionId]

	message, err := srv.artBot.CreateImageMessage(req)
	if err != nil {
		return nil, err
	}

	job, err = srv.artBot.Send(ctx, message)

	return job, err
}

func (srv *ArtBotService) GetModels(ctx context.Context) (res *model2.ArtBotModelsResponse, err error) {
	models, err := srv.artBot.Client.GetModels(ctx)
	if err != nil {
		return nil, err
	}

	return &model2.ArtBotModelsResponse{
		Models: models,
	}, nil

}

func (srv *ArtBotService) GetSamplers(ctx context.Context) (res *model2.ArtBotSamplersResponse, err error) {
	samplers, err := srv.artBot.Client.GetSamplers(ctx)
	if err != nil {
		return nil, err
	}

	return &model2.ArtBotSamplersResponse{
		Samplers: samplers,
	}, nil

}

func (srv *ArtBotService) GetLoras(ctx context.Context) (res *model2.ArtBotLorasResponse, err error) {
	loras, err := srv.artBot.Client.GetLoras(ctx)
	if err != nil {
		return nil, err
	}

	return &model2.ArtBotLorasResponse{
		Loras: loras,
	}, nil
}

func (srv *ArtBotService) RefreshLoras(ctx context.Context) (err error) {
	return srv.artBot.Client.RefreshLoras(ctx)
}

func (srv *ArtBotService) Progress(ctx context.Context) (res *model2.ProgressResponse, err error) {
	return srv.artBot.Client.Progress(ctx)

}

func (srv *ArtBotService) GetOptions(ctx context.Context) (res *model2.OptionsResponse, err error) {
	return srv.artBot.Client.GetOptions(ctx)

}

func (srv *ArtBotService) SetOptions(ctx context.Context, options *model2.OptionsRequest) (err error) {
	return srv.artBot.Client.SetOptions(ctx, options)

}

func (srv *ArtBotService) WebhookText(ctx context.Context, notify *request2.ParaQueueNotify) {
	srv.artBot.Logger.Info("test notify:", "info key", notify)
}

func (srv *ArtBotService) GetControlNetModelList(ctx context.Context) (res *controlNet.ArtBotControlNetModelResponse, err error) {
	models, err := srv.artBot.Client.GetControlNetModelList(ctx)
	if err != nil {
		return nil, err
	}

	return &controlNet.ArtBotControlNetModelResponse{
		ControlNetModels: models,
	}, nil
}

func (srv *ArtBotService) GetControlNetModuleList(ctx context.Context) (res *controlNet.ArtBotControlNetModuleResponse, err error) {
	modules, err := srv.artBot.Client.GetControlNetModuleList(ctx)
	if err != nil {
		return nil, err
	}

	return &controlNet.ArtBotControlNetModuleResponse{
		Modules: modules,
	}, nil
}

func (srv *ArtBotService) GetControlNetControlTypesList(ctx context.Context) (res *controlNet.ArtBotControlNetControlTypeResponse, err error) {
	controlTypes, err := srv.artBot.Client.GetControlNetControlTypesList(ctx)
	if err != nil {
		return nil, err
	}

	return &controlNet.ArtBotControlNetControlTypeResponse{
		ControlNetTypes: controlTypes,
	}, nil
}

func (srv *ArtBotService) GetControlNetSettings(ctx context.Context) (res *controlNet.ArtBotControlNetSettingsResponse, err error) {
	settings, err := srv.artBot.Client.GetControlNetSettings(ctx)
	if err != nil {
		return nil, err
	}

	return &controlNet.ArtBotControlNetSettingsResponse{
		ControlNetSettings: settings,
	}, nil
}

func (srv *ArtBotService) GetControlNetVersion(ctx context.Context) (res *controlNet.ArtBotControlNetVersionResponse, err error) {
	version, err := srv.artBot.Client.GetControlNetVersion(ctx)
	if err != nil {
		return nil, err
	}

	return &controlNet.ArtBotControlNetVersionResponse{
		ControlNetVersion: version,
	}, nil
}

func (srv *ArtBotService) DetectControlNet(ctx context.Context, req *sd.ParaControlNetDetectInfo) (res *controlNet.ArtBotControlNetDetectResponse, err error) {
	result, err := srv.artBot.Client.DetectControlNet(ctx, req.DetectInfo)
	if err != nil {
		return nil, err
	}

	return &controlNet.ArtBotControlNetDetectResponse{
		Res: result,
	}, nil
}
