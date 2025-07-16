package generator

import (
	"fmt"
	"github.com/rosyrain/rgin/internal/project"
	"github.com/rosyrain/rgin/internal/template"
	"io"
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
		{"controller/post.go", "/controller/post.go.tmpl"},

		{"dao/mysql/mysql.go", "/dao/mysql/mysql.go.tmpl"},
		{"dao/mysql/error_code.go", "/dao/mysql/error_code.go.tmpl"},

		{"dao/sqlite/sqlite.go", "/dao/sqlite/sqlite.go.tmpl"},
		{"dao/sqlite/error_code.go", "/dao/sqlite/error_code.go.tmpl"},
		{"dao/sqlite/user.go", "/dao/sqlite/user.go.tmpl"},
		{"dao/sqlite/post.go", "/dao/sqlite/post.go.tmpl"},

		{"dao/redis/redis.go", "/dao/redis/redis.go.tmpl"},
		{"dao/redis/key.go", "/dao/redis/key.go.tmpl"},

		{"logger/logger.go", "/logger/logger.go.tmpl"},

		{"logic/error_code.go", "/logic/error_code.go.tmpl"},
		{"logic/user.go", "/logic/user.go.tmpl"},
		{"logic/request.go", "/logic/request.go.tmpl"},
		{"logic/post.go", "/logic/post.go.tmpl"},

		{"middlewares/auth.go", "/middlewares/auth.go.tmpl"},
		{"middlewares/ratelimit.go", "/middlewares/ratelimit.go.tmpl"},
		{"middlewares/cors.go", "/middlewares/cors.go.tmpl"},

		{"models/create_table.sql", "/models/create_table.sql.tmpl"},
		{"models/create_table.sql.example", "/models/create_table.sql.example.tmpl"},
		{"models/params.go.example", "/models/params.go.example.tmpl"},
		{"models/params.go", "/models/params.go.tmpl"},
		{"models/user.go", "/models/user.go.tmpl"},
		{"models/post.go", "/models/post.go.tmpl"},

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

	// 示例代码文件列表
	exampleFiles = []string{
		"example/conf/config.yaml",
		"example/controller/code.go",
		"example/controller/request.go",
		"example/controller/response.go",
		"example/controller/validator.go",
		"example/controller/user.go",
		"example/controller/post.go",
		"example/controller/community.go",
		"example/controller/vote.go",

		"example/dao/mysql/mysql.go",
		"example/dao/mysql/error_code.go",
		"example/dao/mysql/user.go",
		"example/dao/mysql/post.go",
		"example/dao/mysql/community.go",

		"example/dao/redis/redis.go",
		"example/dao/redis/key.go",
		"example/dao/redis/post.go",
		"example/dao/redis/vote.go",

		"example/logger/logger.go",

		"example/logic/user.go",
		"example/logic/post.go",
		"example/logic/community.go",
		"example/logic/vote.go",

		"example/middlewares/auth.go",
		"example/middlewares/ratelimit.go",
		"example/middlewares/cors.go",

		"example/models/create_table.sql",
		"example/models/params.go",
		"example/models/user.go",
		"example/models/post.go",
		"example/models/community.go",

		"example/pkg/jwt/jwt.go",
		"example/pkg/snowflask/snowflask.go",

		"example/router/route.go",

		"example/settings/settings.go",

		"example/main.go",
		"example/go.mod",
		"example/go.sum",
		"example/Dockerfile",
		"example/wait-for.sh",
		"example/README.md",
	}
)

// Options 定义项目初始化选项
type Options struct {
	ProjectName string
	WithExample bool // 是否生成示例代码
}

// InitProject 初始化新项目
func InitProject(opts *Options) error {
	// 创建项目结构
	proj := project.NewProject(opts.ProjectName)
	if err := proj.Create(); err != nil {
		return fmt.Errorf("create project struct failed: %v", err)
	}

	// 生成基础文件
	if err := generateFiles(proj); err != nil {
		return fmt.Errorf("generate files failed: %v", err)
	}

	// 如果需要生成示例代码
	if opts.WithExample {
		// 创建 example 目录
		if err := os.MkdirAll(filepath.Join(proj.RootDir, "example"), os.ModePerm); err != nil {
			return fmt.Errorf("create example directory failed: %v", err)
		}
		
		if err := copyExampleFiles(proj); err != nil {
			return fmt.Errorf("copy example files failed: %v", err)
		}
	}

	// 项目创建成功后输出引导信息
	fmt.Println("\n项目创建成功！")
	fmt.Println("----------------------------------------")
	fmt.Println("1. 基础项目依赖 SQLite，首次运行前请手动执行 models/create_table.sql 初始化表结构")
	fmt.Println("   cd", opts.ProjectName)
	fmt.Println("   sqlite3 ./data/app.db < models/create_table.sql")
	fmt.Println("2. 已保留 MySQL/Redis 相关连接代码，可按需在 main.go取消对应注释 config填写相关信息")
	fmt.Println("3. 更多文档和示例请参考 README.md 或 参数 --with-example 生成的 example 目录")
	fmt.Println("----------------------------------------\n")

	return nil
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
	if err != nil {
		return err
	}
	defer file.Close()

	// 传递 map，包含 Name、RootDir、ModulePath
	data := map[string]interface{}{
		"Name":       proj.Name,
		"RootDir":    proj.RootDir,
		"ModulePath": proj.ModulePath,
	}
	if err := tmpl.Execute(file, data); err != nil {
		return err
	}
	return nil
}

// copyExampleFiles 复制示例代码文件
func copyExampleFiles(proj *project.Project) error {
	exampleRoot := "internal/template/templates/example/bluebell"
	for _, filePath := range exampleFiles {
		srcPath := filepath.Join(exampleRoot, filePath[len("example/"):])
		srcFile, err := os.Open(srcPath)
		if err != nil {
			return fmt.Errorf("failed to open example file %s: %v", srcPath, err)
		}
		defer srcFile.Close()

		dstPath := filepath.Join(proj.RootDir, filePath)
		if err := os.MkdirAll(filepath.Dir(dstPath), os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory for %s: %v", dstPath, err)
		}

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %v", dstPath, err)
		}
		defer dstFile.Close()

		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return fmt.Errorf("failed to copy file content to %s: %v", dstPath, err)
		}
		// 只在 debug 或需要时打印详细信息，否则静默
	}
	return nil
}
