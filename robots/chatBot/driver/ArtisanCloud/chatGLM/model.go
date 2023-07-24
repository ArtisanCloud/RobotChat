package chatGLM

type GLMResponse struct {
	Response string     `json:"response"`
	History  [][]string `json:"history"`
	Status   int        `json:"status"`
	Time     string     `json:"time"`
}
