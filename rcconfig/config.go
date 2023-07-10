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
	Channel         string                  `yaml:"Channel" json:",optional"`
	StableDiffusion config2.StableDiffusion `yaml:"StableDiffusion" json:",optional"`
	Queue           Queue                   `yaml:"Queue" json:",optional"`
	Log             Log                     `yaml:"Log" json:",optional"`
}

type ChatBot struct {
	Channel  string          `yaml:"Channel" json:",optional"`
	ChatGPT  config.ChatGPT  `yaml:"ChatGPT" json:",optional"`
	XFYun    config.XFYun    `yaml:"XFYun" json:",optional"`
	THUDMGLM config.THUDMGLM `yaml:"THUDM_GLM" json:",optional"`
	Queue    Queue           `yaml:"Queue" json:",optional"`
	Log      Log             `yaml:"Log" json:",optional"`
}

type Redis struct {
	Addr       string `yaml:"Addr" json:",optional"`
	ClientName string `yaml:"ClientName" json:",optional"`
	Username   string `yaml:"Username" json:",optional"`
	Password   string `yaml:"Password" json:",optional"`
	DB         int    `yaml:"DB" json:",optional"`
	MaxRetries int    `yaml:"MaxRetries" json:",optional"`
}

type Queue struct {
	Driver    string `yaml:"Driver" json:",optional"`
	NotifyUrl string `yaml:"NotifyUrl" json:",optional"`
	Redis     Redis  `yaml:"Redis" json:",optional"`
}

type Log struct {
	Driver   string `yaml:"Driver" json:",optional"`
	Env      string `yaml:"Env" json:",optional"`
	InfoLog  string `yaml:"InfoLog" json:",optional"`
	ErrorLog string `yaml:"ErrorLog" json:",optional"`
}

type RCConfig struct {
	Database Database `yaml:"Database" json:",optional"`
	Auth     Auth     `yaml:"Auth" json:",optional"`
	ArtBot   ArtBot   `yaml:"ArtBot" json:",optional"`
	ChatBot  ChatBot  `yaml:"ChatBot" json:",optional"`
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
