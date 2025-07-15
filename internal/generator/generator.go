package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

// Options 定义项目初始化选项
type Options struct {
	ProjectName  string
	Database     string // sqlite, mysql, none
	WithExamples bool
}

// ComponentOptions 定义组件生成选项
type ComponentOptions struct {
	Type string // controller, model, service
	Name string
}

// InitProject 初始化新项目
func InitProject(opts *Options) error {
	// 创建项目目录
	if err := os.MkdirAll(opts.ProjectName, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// 生成基础结构
	if err := generateBaseStructure(opts); err != nil {
		return fmt.Errorf("failed to generate base structure: %w", err)
	}

	// 生成示例代码
	if opts.WithExamples {
		if err := generateExamples(opts); err != nil {
			return fmt.Errorf("failed to generate examples: %w", err)
		}
	}

	return nil
}

// AddComponent 添加新组件
func AddComponent(opts *ComponentOptions) error {
	// 验证当前目录是否是有效的项目
	if !isValidProject() {
		return fmt.Errorf("not a valid rgin project directory")
	}

	// 根据组件类型生成代码
	switch opts.Type {
	case "controller":
		return generateController(opts.Name)
	case "model":
		return generateModel(opts.Name)
	case "service":
		return generateService(opts.Name)
	default:
		return fmt.Errorf("unknown component type: %s", opts.Type)
	}
}

// generateBaseStructure 生成基础项目结构
func generateBaseStructure(opts *Options) error {
	// 创建基础目录结构
	dirs := []string{
		"cmd",
		"internal/controller",
		"internal/model",
		"internal/service",
		"internal/middleware",
		"pkg",
		"api",
		"config",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(opts.ProjectName, dir), 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// TODO: 生成基础文件
	// 1. main.go
	// 2. go.mod
	// 3. config files
	// 4. README.md

	return nil
}

// generateExamples 生成示例代码
func generateExamples(opts *Options) error {
	// TODO: 实现示例代码生成
	return nil
}

// isValidProject 检查当前目录是否是有效的项目
func isValidProject() bool {
	// TODO: 实现项目验证逻辑
	return true
}

// generateController 生成控制器代码
func generateController(name string) error {
	// TODO: 实现控制器生成
	return nil
}

// generateModel 生成模型代码
func generateModel(name string) error {
	// TODO: 实现模型生成
	return nil
}

// generateService 生成服务代码
func generateService(name string) error {
	// TODO: 实现服务生成
	return nil
}
