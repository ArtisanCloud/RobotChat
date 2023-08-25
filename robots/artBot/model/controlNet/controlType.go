package controlNet

type ControlType struct {
	ModuleList    []string `json:"module_list"`
	ModelList     []string `json:"model_list"`
	DefaultOption string   `json:"default_option"`
	DefaultModel  string   `json:"default_model"`
}

type ControlNetTypes struct {
	ControlTypes struct {
		All       *ControlType `json:"All"`
		Canny     *ControlType `json:"Canny"`
		Depth     *ControlType `json:"Depth"`
		Normal    *ControlType `json:"Normal"`
		OpenPose  *ControlType `json:"OpenPose"`
		MLSD      *ControlType `json:"MLSD"`
		Lineart   *ControlType `json:"Lineart"`
		SoftEdge  *ControlType `json:"SoftEdge"`
		Scribble  *ControlType `json:"Scribble"`
		Seg       *ControlType `json:"Seg"`
		Shuffle   *ControlType `json:"Shuffle"`
		Tile      *ControlType `json:"Tile"`
		Inpaint   *ControlType `json:"Inpaint"`
		IP2P      *ControlType `json:"IP2P"`
		Reference *ControlType `json:"Reference"`
		T2IA      *ControlType `json:"T2IA"`
	} `json:"control_types"`
}
