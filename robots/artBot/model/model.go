package model

type ArtBotModel struct {
	Title     string `json:"title"`
	ModelName string `json:"model_name"`
	Hash      string `json:"hash"`
	FileName  string `json:"filename"`
	Config    string `json:"config"`
}

type ArtBotModelsResponse struct {
	Models []*ArtBotModel
}
