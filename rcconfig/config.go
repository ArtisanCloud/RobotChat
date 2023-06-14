package rcconfig

import (
	config2 "github.com/ArtisanCloud/RobotChat/artRobot/config"
	"github.com/ArtisanCloud/RobotChat/chatRobot/config"
	"gopkg.in/yaml.v2"
	"os"
)

type ConfigInterface interface {
	GetName() string
	Validate() error
}

type ArtRobot struct {
	config2.StableDiffusionConfig `yaml:"StableDiffusion"`
}

type ChatRobot struct {
	config.ChatGPTConfig `yaml:"ChatGPT"`
}

type RCConfig struct {
	ArtRobot  `yaml:"ArtRobot"`
	ChatRobot `yaml:"ChatRobot"`
}

func LoadRCConfig() *RCConfig {
	// Read the file
	data, err := os.ReadFile("../config.yml")
	if err != nil {
		panic(err)
		return nil
	}

	// Create a struct to hold the YAML data
	var config = &RCConfig{}

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, config)
	if err != nil {
		panic(err)
		return nil
	}

	// Print the data
	return config
}
