package contract

import "github.com/ArtisanCloud/RobotChat/robots/kernel/model"

// ConversationManagerInterface 定义接口
type ConversationManagerInterface interface {
	StartConversation(robot RobotInterface)
	EndConversation()
	SendMessage(message *model.Message) *model.Job
	ReceiveMessage(job model.Job) *model.Message
}
