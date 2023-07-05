package contract

import "github.com/ArtisanCloud/RobotChat/robots/kernel/model"

// RobotInterface 定义接口
type RobotInterface interface {
	StartConversation(conversationManager ConversationManagerInterface)
	EndConversation() error
	Introduce() *model.RobotAttributes
	SendMessage(message *model.Message) *model.Job
	ReceiveMessage(job model.Job) *model.Message
}
