package rbtemplate

import (
	"strings"
	"text/template"
)

const (
	TextImagePrompt         = "text:image-prompt"
	textImagePromptTemplate = `我需要你帮助我生成用于图像生成的提示, 给你提供几个例子:
1. annihilation movie style, Astral realm ruins, shiny white female robot full-body Porcelain and hammered matt silver showing cracked inner working, tiny white flowers growing from within, light from within all white,, reflecting in the mirror, flowers growing on the robot parts, , swamp, cinematic lighting, amazing composition , 3d octane render, unreal engine, hyper realistic, soft illumination, trending artstation, environmental concept art, all in grey, 4k trending on ArtStation
2. A young female walking toward the camera. She is wearing a V neck shirt and short beige pants. Shot from above
3. photo realistic, female fashion model, space station corridor, bokeh, hyper realistic, helmet, white, Canon R5 ISO100 F1.4

请为"{{.}}"这句话生成提示词, 按照如下格式输出:
"
1. prompt1... \n
(中文阐述提示词的含义)
2. prompt2... \n
(中文阐述提示词的含义)
3. prompt3... \n
(中文阐述提示词的含义)
"`
)

type TextImagePromptTemplate struct {
	name     string
	template *template.Template
}

func (t TextImagePromptTemplate) GetName() string {
	return t.name
}

func (t TextImagePromptTemplate) Execute(input string) string {
	var buf strings.Builder
	err := t.template.Execute(&buf, input)
	// must not be error
	if err != nil {
		return input
	}
	return buf.String()
}

func NewTextImagePromptMode() *TextImagePromptTemplate {
	temp := template.Must(template.New("text-image-prompt").Parse(textImagePromptTemplate))
	return &TextImagePromptTemplate{
		name:     TextImagePrompt,
		template: temp,
	}
}
