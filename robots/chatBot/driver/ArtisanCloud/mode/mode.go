package mode

import (
	"errors"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/driver/ArtisanCloud/chain"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/driver/ArtisanCloud/chain/rbtemplate"
	"regexp"

	"strings"
)

var prefixMatcher = regexp.MustCompile(`^/([a-zA-Z0-9:_]{1,32})`)

type TextModeManager struct {
	modeMap map[string]*chain.TextChain
}

func NewTextModeManager(model chain.TextChainAdapter) *TextModeManager {
	manager := TextModeManager{
		modeMap: make(map[string]*chain.TextChain),
	}
	manager.RegisterModeChain("default", chain.NewTextChain(model))
	manager.RegisterModeChain("image:prompt", chain.NewTextChain(model).SetTemplate(rbtemplate.NewTextImagePromptMode()))
	return &manager
}

func (t *TextModeManager) RegisterModeChain(mode string, c *chain.TextChain) {
	t.modeMap[mode] = c
}

func (t *TextModeManager) Response(message string) (string, error) {
	mode := prefixMatcher.FindString(message)
	if mode == "" {
		mode = "default"
	} else {
		message = strings.TrimPrefix(message, mode)
		message = strings.TrimLeft(message, " ")
		mode = strings.TrimLeft(mode, "/")
	}
	if t.modeMap[mode] == nil {
		return "", errors.New("mode not found")
	}
	return t.modeMap[mode].Response(message)
}
