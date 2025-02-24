package template

import (
	"embed"
	"path/filepath"
	"text/template"
)

var templateFS embed.FS

func LoadTemplate(name string) (*template.Template, error) {
	tmplPath := filepath.Join("template", name)
	data, err := templateFS.ReadFile(tmplPath)
	if err != nil {
		return nil, err
	}

	return template.New(name).Parse(string(data))
}
