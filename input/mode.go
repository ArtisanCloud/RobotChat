package input

import (
	"regexp"
	"strings"
)

const Default = "default"

var prefixMatcher = regexp.MustCompile(`^/([a-zA-Z0-9:_]{1,32})`)

type ModeMessage struct {
	Mode    string
	Message string
}

func FilterModeAndMessage(message string) ModeMessage {
	mode := prefixMatcher.FindString(message)
	if mode == "" {
		mode = Default
	} else {
		message = strings.TrimPrefix(message, mode)
		message = strings.TrimLeft(message, " ")
		mode = strings.TrimLeft(mode, "/")
	}
	return ModeMessage{
		Mode:    mode,
		Message: message,
	}
}
