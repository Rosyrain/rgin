package project

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type Project struct {
	Name       string
	RootDir    string
	ModulePath string // 新增字段，记录 go.mod 的 module 路径
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
		"example", // 添加示例代码目录
		"example/conf",
		"example/controller",
		"example/dao/mysql",
		"example/dao/redis",
		"example/logger",
		"example/logic",
		"example/middlewares",
		"example/models",
		"example/pkg/jwt",
		"example/pkg/snowflask",
		"example/router",
		"example/settings",
	}
)

func NewProject(name string) *Project {
	proj := &Project{
		Name:    name,
		RootDir: name,
	}
	// 读取 go.mod.tmpl 获取 module 路径
	modTmplPath := filepath.Join("internal/template/templates/go.mod.tmpl")
	file, err := os.Open(modTmplPath)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(line, "module ") {
				proj.ModulePath = strings.TrimSpace(strings.TrimPrefix(line, "module "))
				break
			}
		}
		file.Close()
	}
	if proj.ModulePath == "" {
		proj.ModulePath = name // 兜底
	}
	return proj
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
