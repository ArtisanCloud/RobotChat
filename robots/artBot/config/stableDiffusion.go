package config

type StableDiffusion struct {
	Token     string `yaml:"Token" json:",optional"`
	BaseUrl   string `yaml:"BaseUrl" json:",optional"`
	PrefixUri string `yaml:"PrefixUri" json:",optional"`
	Version   string `yaml:"Version" json:",optional"`
	HttpDebug bool   `yaml:"HttpDebug" json:",optional"`
	ProxyUrl  string `yaml:"ProxyUrl" json:",optional"`
}

func (c *StableDiffusion) GetName() string {
	return "StableDiffusion"
}

func (c *StableDiffusion) Validate() error {
	return nil
}
