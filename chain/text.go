package chain

import (
	"github.com/artisancloud/robotchat/chain/rbtemplate"
)

type TextChainAdapter interface {
	HandleInput(input string) (string, error)
}

type TextFilter interface {
	Filter(input string) string
}

type TextChain struct {
	template rbtemplate.Template
	adapter  TextChainAdapter
	before   TextFilter
	after    TextFilter
}

func NewTextChain(adapter TextChainAdapter) *TextChain {
	return &TextChain{
		adapter: adapter,
	}
}

func (t *TextChain) SetTemplate(template rbtemplate.Template) *TextChain {
	t.template = template
	return t
}

func (t *TextChain) SetBefore(before TextFilter) *TextChain {
	t.before = before
	return t
}

func (t *TextChain) SetAfter(after TextFilter) *TextChain {
	t.after = after
	return t
}

func (t *TextChain) Response(message string) (string, error) {
	msg := message
	if t.before != nil {
		msg = t.before.Filter(msg)
	}
	if t.template != nil {
		msg = t.template.Execute(msg)
	}
	msg, err := t.adapter.HandleInput(msg)
	if err != nil {
		return "", err
	}
	if t.after != nil {
		msg = t.after.Filter(msg)
	}
	return msg, err
}
