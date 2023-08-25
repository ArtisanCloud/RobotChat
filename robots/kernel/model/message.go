package model

import (
	"context"
	"gorm.io/datatypes"
)

// MessageType 是消息类型的枚举定义
type MessageType int

const (
	TextMessage MessageType = iota
	ImageMessage
	AudioMessage
	// 可以添加更多消息类型...
)

type MetaData struct {
	ErrMsg       string           `json:"errMsg"`
	ErrCode      int              `json:"errorCode"`
	Robot        *RobotAttributes `json:"robot"`
	Conversation *Conversation    `json:"conversation"`
	CustomerId   int64            `json:"customerId"`
	// 用来保存请问的原始数据，但是不用包含输入的图片
	RequestData datatypes.JSON `json:"requestData"`
}

// Message 是消息的结构定义
type Message struct {
	RobotModel

	ModelType   string         `gorm:"comment:模型类型" json:"modelType"`
	MessageType MessageType    `gorm:"comment:消息类型" json:"messageType"`
	Author      string         `gorm:"comment:作者" json:"author"`
	Content     datatypes.JSON `gorm:"comment:内容" json:"content"`
	Metadata    MetaData       `gorm:"comment:meta" json:"metadata"`
}

type Content struct {
	ContentType string
	Parts       []string
}

type Author struct {
	Role Role
}

type Role string

const (
	SystemRole    Role = "system"
	UserRole      Role = "user"
	AssistantRole Role = "assistant"
)

// Middleware 是中间件函数的类型定义
type HandlePreSend func(ctx context.Context, job *Job) (*Job, error)
type HandlePostReply func(ctx context.Context, job *Job) (*Job, error)

type ErrReply struct {
	Ctx context.Context
	Job *Job
	Err error
}
type HandleError func(reply *ErrReply)

func NewMessage(msgType MessageType) *Message {

	msg := &Message{
		RobotModel:  *NewRobotModel(),
		MessageType: msgType,
	}

	return msg
}

func CopyMessage(source *Message) *Message {
	msg := &Message{
		RobotModel: *NewRobotModel(),
	}
	msg.ModelType = source.ModelType
	msg.MessageType = source.MessageType
	msg.Author = source.Author
	msg.Content = source.Content
	msg.Metadata = source.Metadata

	return msg
}
