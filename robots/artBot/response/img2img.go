package response

import "github.com/ArtisanCloud/RobotChat/robots/artBot/request"

type Image2Image struct {
	Images        []string `json:"images"`
	DecodedImages [][]byte
	Parameters    request.Image2Image `json:"parameters"`
	Info          string              `json:"info"`
}
