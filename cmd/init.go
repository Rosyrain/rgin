package cmd

import (
	"fmt"
	"github.com/rosyrain/rgin/internal/generator"
	"github.com/rosyrain/rgin/internal/i18n"
	"github.com/spf13/cobra"
)

var (
	withExample bool // 是否生成示例代码
)

var initCmd = &cobra.Command{
	Use:   "init [project_name]",
	Short: i18n.T("init_short"),
	Long:  i18n.T("init_long"),
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]

		opts := &generator.Options{
			ProjectName: projectName,
			WithExample: withExample,
		}

		if err := generator.InitProject(opts); err != nil {
			return fmt.Errorf("failed to initialize project: %v", err)
		}

		fmt.Printf("Successfully created project %s\n", projectName)
		if withExample {
			fmt.Println("Example code has been generated in the project")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVar(&withExample, "with-example", false, i18n.T("flag_with_example"))
} 