package template

import (
	"embed"
	"fmt"
	"io/fs"
	"text/template"
)

//go:embed templates/*
var templateFS embed.FS

func init() {
	// 遍历检查嵌入的模板
	err := fs.WalkDir(templateFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("Embedded file:", path)
		return nil
	})
	if err != nil {
		fmt.Println("embed files failed,err: ", err)
	}
}

func LoadTemplate(name string) (*template.Template, error) {

	//tmplPath := filepath.Join("templates", name) // 使用此方法会导致错误
	tmplPath := "templates" + name
	fmt.Println("path: ", tmplPath)
	data, err := templateFS.ReadFile(tmplPath)
	if err != nil {
		return nil, err
	}

	return template.New(name).Parse(string(data))
}
