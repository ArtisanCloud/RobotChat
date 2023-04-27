package rbconfig

type RobotConfig struct {
	ChatGPT *ChatGPTConfig
}

func (r RobotConfig) GetName() string {
	return "Robot"
}

func (r RobotConfig) Validate() error {
	if r.ChatGPT == nil {
		return nil
	}
	if err := r.ChatGPT.Validate(); err != nil {
		return err
	}
	return nil
}
