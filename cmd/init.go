package cmd

import (
	"fmt"
	"github.com/rosyrain/rgin/internal/generator"
	"github.com/spf13/cobra"
)

var (
	withDB      string
	withExample bool
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "初始化一个新的 Gin 项目",
	Long: `初始化一个新的 Gin 项目，包含推荐的项目结构和基本设置。

使用示例:
  1. 创建基础项目:
     rgin init myapp

  2. 使用 SQLite 数据库:
     rgin init myapp --with-db sqlite

  3. 包含示例代码:
     rgin init myapp --with-examples

  4. 完整示例:
     rgin init myapp --with-db sqlite --with-examples

项目结构:
  myapp/
  ├── cmd/          - 命令行入口
  ├── internal/     - 内部代码
  │   ├── controller/ - 控制器
  │   ├── model/    - 数据模型
  │   └── service/  - 业务服务
  ├── pkg/         - 公共包
  ├── api/         - API 文档
  └── config/      - 配置文件

注意事项:
  - 项目名称应该是有效的 Go 包名
  - 默认使用 SQLite 数据库（轻量级，无需配置）
  - 使用 --with-examples 可以生成带有注释的示例代码`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		fmt.Printf("Creating new project: %s\n", projectName)
		
		options := &generator.Options{
			ProjectName:  projectName,
			Database:     withDB,
			WithExamples: withExample,
		}
		
		if err := generator.InitProject(options); err != nil {
			cobra.CheckErr(err)
		}
		
		fmt.Printf("✨ Successfully created project: %s\n", projectName)
		fmt.Println("\nNext steps:")
		fmt.Printf("  cd %s\n", projectName)
		fmt.Println("  go mod tidy")
		fmt.Println("  go run main.go")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// 添加命令行参数
	initCmd.Flags().StringVar(&withDB, "with-db", "sqlite", "Database to use (sqlite/mysql/none)")
	initCmd.Flags().BoolVar(&withExample, "with-examples", false, "Include example code")
} 