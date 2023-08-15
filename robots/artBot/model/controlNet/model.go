package controlNet

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/model"
)

type ControlNetVersion struct {
	Version string `json:"vesion"`
}

type ControlNetModel struct {
	ModelList []string `json:"model_list"`
}

type Modules struct {
	ModuleList   []string        `json:"module_list"`
	ModuleDetail *object.HashMap `json:"module_detail"`
}

type ControlNetSettings struct {
	ControlNetMaxModelsNum int `json:"control_net_max_models_num"`
}

type DetectInfo struct {
	ControlNetModule       string        `json:"controlnet_module"`
	ControlNetInputImages  []interface{} `json:"controlnet_input_images"`
	ControlNetProcessorRes int           `json:"controlnet_processor_res"`
	ControlNetThresholdA   int           `json:"controlnet_threshold_a"`
	ControlNetThresholdB   int           `json:"controlnet_threshold_b"`
}

type ArtBotControlNetModelResponse struct {
	model.SDResponse

	*ControlNetModel
}

type ArtBotControlNetModuleResponse struct {
	model.SDResponse

	*Modules
}

type ArtBotControlNetSettingsResponse struct {
	model.SDResponse

	*ControlNetSettings
}

type ArtBotControlNetVersionResponse struct {
	model.SDResponse

	*ControlNetVersion
}

type ArtBotControlNetDetectResponse struct {
	model.SDResponse

	Res interface{}
}
