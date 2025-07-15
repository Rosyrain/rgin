package cmd

import (
	"fmt"
	"github.com/rosyrain/rgin/internal/generator"
	"github.com/rosyrain/rgin/internal/i18n"
	"github.com/spf13/cobra"
)

var (
	withDB      string
	withExample bool
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new Gin project", // 将由 i18n 系统更新
	Long:  "Loading...", // 将由 i18n 系统更新
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		
		if i18n.IsChinese() {
			fmt.Printf("正在创建项目：%s\n", projectName)
		} else {
			fmt.Printf("Creating project: %s\n", projectName)
		}
		
		options := &generator.Options{
			ProjectName:  projectName,
			Database:     withDB,
			WithExamples: withExample,
		}
		
		if err := generator.InitProject(options); err != nil {
			cobra.CheckErr(err)
		}
		
		if i18n.IsChinese() {
			fmt.Printf("✨ 项目创建成功：%s\n", projectName)
			fmt.Println("\n后续步骤：")
			fmt.Printf("  cd %s\n", projectName)
			fmt.Println("  go mod tidy")
			fmt.Println("  go run main.go")
		} else {
			fmt.Printf("✨ Successfully created project: %s\n", projectName)
			fmt.Println("\nNext steps:")
			fmt.Printf("  cd %s\n", projectName)
			fmt.Println("  go mod tidy")
			fmt.Println("  go run main.go")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	
	if i18n.IsChinese() {
		initCmd.Flags().StringVar(&withDB, "with-db", "sqlite", "选择数据库 (sqlite/mysql/none)")
		initCmd.Flags().BoolVar(&withExample, "with-examples", false, "包含示例代码")
	} else {
		initCmd.Flags().StringVar(&withDB, "with-db", "sqlite", "Database to use (sqlite/mysql/none)")
		initCmd.Flags().BoolVar(&withExample, "with-examples", false, "Include example code")
	}
} 