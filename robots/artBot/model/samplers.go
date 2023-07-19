package model

type ArtBotSamplersResponse struct {
	SDResponse
	Samplers []*Sampler
}

type Option struct {
	UsesEnSD               string `json:"uses_ensd,omitempty"`
	SecondOrder            string `json:"second_order,omitempty"`
	BrownianNoise          string `json:"brownian_noise,omitempty"`
	Scheduler              string `json:"scheduler,omitempty"`
	DiscardNextToLastSigma string `json:"discard_next_to_last_sigma,omitempty"`
}

type Sampler struct {
	Name    string   `json:"name"`
	Aliases []string `json:"aliases"`
	Options *Option  `json:"options"`
}
