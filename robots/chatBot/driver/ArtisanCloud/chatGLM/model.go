package chatGLM

type GLMRequest struct {
	Prompt  string        `json:"prompt"`
	History []interface{} `json:"history"`
}

type GLMResponse struct {
	Response string     `json:"response"`
	History  [][]string `json:"history"`
	Status   int        `json:"status"`
	Time     string     `json:"time"`
}
