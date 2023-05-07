package model

import (
	"github.com/artisancloud/robotchat/input"
)

type TextModel interface {
	RegisterMode(tmpl *input.Template)
	Input(input string) (string, error)
	SetBefore(before func(input string) (string, error))
	SetAfter(after func(input string) (string, error))
}
