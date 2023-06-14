package config

type StableDiffusionConfig struct {
	Token     string `yaml:"Token"`
	BaseUrl   string `yaml:"BaseUrl"`
	PrefixUri string `yaml:"PrefixUri"`
	Version   string `yaml:"Version"`
	HttpDebug bool   `yaml:"HttpDebug"`
	ProxyURL  string `yaml:"ProxyURL"`
}

func (c *StableDiffusionConfig) GetName() string {
	return "StableDiffusion"
}

func (c *StableDiffusionConfig) Validate() error {
	return nil
}
