package chatRobot

import (
	go_openai "github.com/ArtisanCloud/RobotChat/chatRobot/driver/go-openai"
	fmt "github.com/ArtisanCloud/RobotChat/pkg/printx"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"testing"
)

var testConfig = &rcconfig.RCConfig{}

func init() {

	testConfig = rcconfig.LoadRCConfig()

}

func TestChatRobot_TextResponse(t *testing.T) {

	driver := go_openai.NewDriver(&testConfig.ChatRobot)

	robot, err := NewChatRobot(driver)
	if err != nil {
		t.Error(err)
	}

	conf := robot.Client.GetConfig()
	fmt.Dump(conf)

}
