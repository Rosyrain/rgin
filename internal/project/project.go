package project

import (
	"os"
	"path/filepath"
)

type Project struct {
	Name    string
	RootDir string
}

var (
	// 创建子目录
	dirs = []string{
		"conf",
		"controller",
		"dao/mysql",
		"dao/redis",
		"logger",
		"logic",
		"middlewares",
		"models",
		"pkg/jwt",
		"pkg/snowflask",
		"router",
		"settings",
	}
)

func NewProject(name string) *Project {
	return &Project{
		Name:    name,
		RootDir: name,
	}
}

func (p *Project) Create() error {
	// 创建项目根目录
	if err := os.Mkdir(p.RootDir, os.ModePerm); err != nil {
		return err
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(p.RootDir, dir), os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
