package chatgpt

import (
	"github.com/artisancloud/robotchat/input"
)

type TextAdapter struct {
	model  ChatGPT
	mode   map[string]*input.Template
	before func(input string) (string, error)
	after  func(input string) (string, error)
}

func NewTextAdapter(model ChatGPT) *TextAdapter {
	textAdapter := &TextAdapter{
		model: model,
		mode:  make(map[string]*input.Template),
	}
	return textAdapter
}

func (t *TextAdapter) Input(in string) (out string, err error) {
	// mode build input
	modeMessage := input.FilterModeAndMessage(in)
	if modeMessage.Mode != input.Default {
		tmpl, ok := t.mode[modeMessage.Mode]
		if ok {
			in, err = tmpl.Execute(in)
			if err != nil {
				return "", err
			}
		} else {
			in = modeMessage.Message
		}
	}

	// before
	if t.before != nil {
		in, err = t.before(in)
		if err != nil {
			return "", err
		}
	}

	// response
	oMessages := []Message{
		{
			Role:    "user",
			Content: in,
		},
	}
	choices, err := t.model.Response(oMessages)
	if err != nil {
		return "", err
	}
	out = choices[0].Message.Content

	// after
	if t.after != nil {
		out, err = t.after(out)
		if err != nil {
			return "", err
		}
	}
	return out, nil
}

func (t *TextAdapter) RegisterMode(tmpl *input.Template) {
	t.mode[tmpl.Name] = tmpl
}

func (t *TextAdapter) SetBefore(before func(input string) (string, error)) {
	t.before = before
}

func (t *TextAdapter) SetAfter(after func(input string) (string, error)) {
	t.after = after
}
