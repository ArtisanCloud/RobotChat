package model

type ChatResponse struct {
	Error  string `json:"error,omitempty"`
	Detail string `json:"detail,omitempty"`
	Errors string `json:"errors,omitempty"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
