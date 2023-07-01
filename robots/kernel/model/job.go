package model

type Job struct {
	Id      string   `json:"id"`
	Payload *Message `json:"payload"`
	// 其他任务相关的字段
}
