package rcconfig

import (
	"flag"
	config2 "github.com/ArtisanCloud/RobotChat/robots/artBot/config"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/config"
	"gopkg.in/yaml.v2"
	"os"
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
	Channel                       string `yaml:"Channel"`
	config2.StableDiffusionConfig `yaml:"StableDiffusion"`
	Queue                         `yaml:"Queue"`
	Log                           `yaml:"Log"`
}

type ChatBot struct {
	Channel               string `yaml:"Channel"`
	config.ChatGPTConfig  `yaml:"ChatGPT"`
	config.XFYunConfig    `yaml:"XFYun"`
	config.THUDMGLMConfig `yaml:"THUDM_GLM"`
	Queue                 `yaml:"Queue"`
	Log                   `yaml:"Log"`
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
	Driver    string `yaml:"Driver"`
	NotifyUrl string `yaml:"NotifyUrl"`
	Redis     `yaml:"Redis"`
}

type Log struct {
	Driver   string `yaml:"Driver"`
	Env      string `yaml:"Env"`
	InfoLog  string `yaml:"InfoLog"`
	ErrorLog string `yaml:"ErrorLog"`
}

type RCConfig struct {
	Database `yaml:"Database"`
	Auth     `yaml:"Auth"`
	ArtBot   `yaml:"ArtBot"`
	ChatBot  `yaml:"ChatBot"`
}

func LoadRCConfigByPath(configPath string) *RCConfig {

	// Read the file
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
		return nil
	}

	// Create a struct to hold the YAML data
	var conf = &RCConfig{}

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		panic(err)
		return nil
	}

	// Print the data
	return conf
}

func LoadRCConfig() *RCConfig {

	// 获取配置文件所在的目录路径
	configFile := flag.String("f", "config.yaml", "the config file")
	flag.Parse()
	//fmt.Dump(configFile)

	return LoadRCConfigByPath(*configFile)

}
