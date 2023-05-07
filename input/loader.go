package input

import (
	"bytes"
	"embed"
	"encoding/json"
	"io/ioutil"
	"text/template"
)

type Template struct {
	Name     string
	Template *template.Template
}

func (t *Template) Execute(text string) (string, error) {
	var buf bytes.Buffer
	err := t.Template.Execute(&buf, text)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func LoadTemplate(jsonPath string) ([]*Template, error) {
	var templates []*Template
	jsonBytes, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}

	var data map[string]string
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return nil, err
	}

	for name, content := range data {
		tmpl, err := template.New(name).Parse(content)
		if err != nil {
			return nil, err
		}
		templates = append(templates, &Template{
			Name:     name,
			Template: tmpl,
		})
	}

	return templates, nil
}

func LoadTemplateForEmbed(embeddedFS embed.FS, name string) ([]*Template, error) {
	var templates []*Template

	file, err := embeddedFS.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	jsonBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(jsonBytes)
	if err != nil {
		return nil, err
	}

	var data map[string]string
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return nil, err
	}

	for name, content := range data {
		tmpl, err := template.New(name).Parse(content)
		if err != nil {
			return nil, err
		}
		templates = append(templates, &Template{
			Name:     name,
			Template: tmpl,
		})
	}

	return templates, nil
}
