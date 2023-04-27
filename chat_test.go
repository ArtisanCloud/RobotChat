package robotchat

import (
	"robotchat/rbconfig"
	"testing"
)

func TestChatRobot_TextResponse(t *testing.T) {
	robot, err := NewChatRobot(rbconfig.RobotConfig{
		ChatGPT: &rbconfig.ChatGPTConfig{
			OpenAPIKey:   "key",
			Organization: "org",
			HttpDebug:    true,
			ProxyURL:     "http://127.0.0.1:33210",
		},
	})
	if err != nil {
		t.Error(err)
	}
	response, err := robot.TextResponse("你好")
	if err != nil {
		t.Error(err)
	}
	t.Log(response)
	response, err = robot.TextResponse("/image:prompt 彩色的小鸟")
	if err != nil {
		t.Error(err)
	}
	t.Log(response)
}
