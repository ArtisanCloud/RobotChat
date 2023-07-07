package model

type ProgressRequest struct {
	SkipCurrentImage bool `json:"skip_current_image"`
}

type State struct {
	Skipped             bool   `json:"skipped"`
	Interrupted         bool   `json:"interrupted"`
	Job                 string `json:"job"`
	JobCount            int    `json:"job_count"`
	JobNo               int    `json:"job_no"`
	CurrentStep         int    `json:"sampling_step"`
	TargetSamplingSteps int    `json:"sampling_steps"`
}

type ProgressResponse struct {
	Progress     float64 `json:"progress"`
	ETA          float64 `json:"eta_relative"`
	State        State   `json:"state"`
	CurrentImage string  `json:"current_image"`
}
