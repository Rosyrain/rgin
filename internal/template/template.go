package template

import (
	"os"
	"path/filepath"
	"text/template"
)

func LoadTemplate(name string) (*template.Template, error) {
	tmplPath := filepath.Join("template", name)
	data, err := os.ReadFile(tmplPath)
	if err != nil {
		return nil, err
	}

	return template.New(name).Parse(string(data))
}
