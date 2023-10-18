package model

type Job struct {
	Id      string   `json:"id"`
	Payload *Message `json:"payload"`
	Type    int      `json:"jobType"`
	// 其他任务相关的字段
}
