package template

import (
	"embed"
	"fmt"
	"io/fs"
	"text/template"
)

//go:embed templates/*
var templateFS embed.FS

//go:embed templates/example/*
var ExampleFS embed.FS

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

	// 遍历检查嵌入的示例代码
	err = fs.WalkDir(ExampleFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("Embedded example file:", path)
		return nil
	})
	if err != nil {
		fmt.Println("embed example files failed,err: ", err)
	}
}

func LoadTemplate(name string) (*template.Template, error) {
	tmplPath := "templates" + name
	fmt.Println("path: ", tmplPath)
	data, err := templateFS.ReadFile(tmplPath)
	if err != nil {
		return nil, err
	}

	return template.New(name).Parse(string(data))
}
