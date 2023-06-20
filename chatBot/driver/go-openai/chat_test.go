package go_openai

import (
	"github.com/ArtisanCloud/RobotChat/chatBot"
	fmt "github.com/ArtisanCloud/RobotChat/pkg/printx"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"testing"
)

var testConfig = &rcconfig.RCConfig{}

func init() {

	testConfig = rcconfig.LoadRCConfig()

}

func TestChatBot_TextResponse(t *testing.T) {

	driver := NewDriver(&testConfig.ChatBot)

	robot, err := chatBot.NewChatBot(driver)
	if err != nil {
		t.Error(err)
	}

	conf := robot.Client.GetConfig()
	fmt.Dump(conf)

}
