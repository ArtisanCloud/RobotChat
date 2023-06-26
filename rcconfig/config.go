package rcconfig

import (
	config2 "github.com/ArtisanCloud/RobotChat/robots/artBot/config"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/config"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type ConfigInterface interface {
	GetName() string
	Validate() error
}

type Database struct {
	Driver string `yaml:"Driver"`
	DSN    string `yaml:"DSN"`
}

type Auth struct {
	Account  string `yaml:"Account"`
	Password string `yaml:"Password"`
}

type ArtBot struct {
	config2.StableDiffusionConfig `yaml:"StableDiffusion"`
	Queue                         `yaml:"Queue"`
}

type ChatBot struct {
	config.ChatGPTConfig `yaml:"ChatGPT"`
	Queue                `yaml:"Queue"`
}

type Redis struct {
	Addr       string `yaml:"Addr"`
	ClientName string `yaml:"ClientName"`
	Username   string `yaml:"Username"`
	Password   string `yaml:"Password"`
	DB         int    `yaml:"DB"`
	MaxRetries int    `yaml:"MaxRetries"`
}

type Queue struct {
	Driver string `yaml:"Driver"`
	Redis  `yaml:"Redis"`
}

type RCConfig struct {
	Database `yaml:"Database"`
	Auth     `yaml:"Auth"`
	ArtBot   `yaml:"ArtBot"`
	ChatBot  `yaml:"ChatBot"`
}

func LoadRCConfig() *RCConfig {

	exePath, err := os.Executable()
	if err != nil {
		// 处理错误
		panic(err)
	}

	// 获取配置文件所在的目录路径
	configDir := filepath.Dir(exePath)

	configPath := filepath.Join(configDir, "../config.yml")

	// Read the file
	data, err := os.ReadFile(configPath)
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
