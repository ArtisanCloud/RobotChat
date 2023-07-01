package config

type StableDiffusionConfig struct {
	Token     string `yaml:"Token"`
	BaseUrl   string `yaml:"BaseUrl"`
	PrefixUri string `yaml:"PrefixUri"`
	Version   string `yaml:"Version"`
	HttpDebug bool   `yaml:"HttpDebug"`
	ProxyUrl  string `yaml:"ProxyUrl"`
}

func (c *StableDiffusionConfig) GetName() string {
	return "StableDiffusion"
}

func (c *StableDiffusionConfig) Validate() error {
	return nil
}
