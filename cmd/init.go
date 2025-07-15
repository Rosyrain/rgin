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
	Short: "Initialize a new Gin project",
	Long: `Initialize a new Gin project with the recommended project structure and basic setup.
Examples:
  rgin init myapp
  rgin init myapp --with-db sqlite
  rgin init myapp --with-examples`,
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