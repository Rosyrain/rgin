package generator

import (
	"fmt"
	"github.com/rosyrain/rgin/internal/project"
	"github.com/rosyrain/rgin/internal/template"
	"os"
	"path/filepath"
)

var (
	// 定义一个文件生成的配置结构
	fileTemplates = []struct {
		OutputPath string // 输出文件路径
		Template   string // 模板文件路径
	}{
		{"conf/config.yaml", "/conf/config.yaml.tmpl"},

		{"controller/code.go", "/controller/code.go.tmpl"},
		{"controller/request.go", "/controller/request.go.tmpl"},
		{"controller/response.go", "/controller/response.go.tmpl"},
		{"controller/validator.go", "/controller/validator.go.tmpl"},
		{"controller/user.go", "/controller/user.go.tmpl"},

		{"dao/mysql/mysql.go", "/dao/mysql/mysql.go.tmpl"},
		{"dao/mysql/user.go", "/dao/mysql/user.go.tmpl"},
		{"dao/mysql/error_code.go", "/dao/mysql/error_code.go.tmpl"},

		{"dao/redis/redis.go", "/dao/redis/redis.go.tmpl"},
		{"dao/redis/key.go", "/dao/redis/key.go.tmpl"},
		{"dao/redis/user.go", "/dao/redis/user.go.tmpl"},

		{"logger/logger.go", "/logger/logger.go.tmpl"},

		{"logic/user.go", "/logic/user.go.tmpl"},
		{"logic/request.go", "/logic/request.go.tmpl"},

		{"middlewares/auth.go", "/middlewares/auth.go.tmpl"},
		{"middlewares/ratelimit.go", "/middlewares/ratelimit.go.tmpl"},

		{"models/create_table.sql", "/models/create_table.sql.tmpl"},
		{"models/params.go", "/models/params.go.tmpl"},
		{"models/user.go", "/models/user.go.tmpl"},

		{"pkg/jwt/jwt.go", "/pkg/jwt/jwt.go.tmpl"},
		{"pkg/snowflask/snowflask.go", "/pkg/snowflask/snowflask.go.tmpl"},

		{"router/route.go", "/router/route.go.tmpl"},

		{"settings/settings.go", "/settings/settings.go.tmpl"},

		{"main.go", "/main.go.tmpl"},
		{"go.mod", "/go.mod.tmpl"},
		{"go.sum", "/go.sum.tmpl"},
		{"Dockerfile", "/Dockerfile.tmpl"},
		{"wait-for.sh", "/wait-for.sh.tmpl"},
	}
)

// create project template
func GenerateProject(projectName string) (err error) {
	// create project struct
	// 创建项目结构
	proj := project.NewProject(projectName)
	if err = proj.Create(); err != nil {
		return fmt.Errorf("create project struct failed: %v", err)
	}

	// 生成文件
	if err = generateFiles(proj); err != nil {
		return fmt.Errorf("生成文件失败: %v", err)
	}

	return err
}

func generateFiles(proj *project.Project) error {
	for _, fileConfig := range fileTemplates {
		// 从配置中提取模板路径和输出路径
		tmplPath := fileConfig.Template
		outputPath := fileConfig.OutputPath

		// 调用 generateFromTemplate 函数生成文件
		if err := generateFromTemplate(proj, tmplPath, outputPath); err != nil {
			return fmt.Errorf("failed to generate file %s: %w", outputPath, err)
		}
	}
	return nil
}

func generateFromTemplate(proj *project.Project, tmplPath, outputPath string) error {
	tmpl, err := template.LoadTemplate(tmplPath)
	if err != nil {
		return err
	}

	outputFile := filepath.Join(proj.RootDir, outputPath)
	if err := os.MkdirAll(filepath.Dir(outputFile), os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(outputFile)
	fmt.Println("create file:", outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, proj)
}
