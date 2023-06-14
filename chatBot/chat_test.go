package chatBot

import (
	go_openai "github.com/ArtisanCloud/RobotChat/chatBot/driver/go-openai"
	fmt "github.com/ArtisanCloud/RobotChat/pkg/printx"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"testing"
)

var testConfig = &rcconfig.RCConfig{}

func init() {

	testConfig = rcconfig.LoadRCConfig()

}

func TestChatBot_TextResponse(t *testing.T) {

	driver := go_openai.NewDriver(&testConfig.ChatBot)

	robot, err := NewChatBot(driver)
	if err != nil {
		t.Error(err)
	}

	conf := robot.Client.GetConfig()
	fmt.Dump(conf)

}
