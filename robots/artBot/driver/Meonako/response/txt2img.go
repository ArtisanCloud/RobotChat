package response

import api "github.com/Meonako/webui-api"

type Text2Image struct {
	Images        []string      `json:"images"`     // Base64-encoded Images Data
	DecodedImages [][]byte      `json:"-"`          // Base64-decoded Images Data store here after "DecodeAllImages()" called
	Parameters    api.Txt2Image `json:"parameters"` // Generation Parameters. Should be the same value as the one you pass to generate.
	Info          string        `json:"info"`       // Info field contains generation parameters like "parameters" field but in long string instead.
}
