package cmd

import (
	"fmt"
	"github.com/rosyrain/rgin/internal/generator"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [component] [name]",
	Short: "Add a component to the project",
	Long: `Add a new component to your Gin project.
Available components:
  - controller: Add a new controller
  - model: Add a new model
  - service: Add a new service

Examples:
  rgin add controller user
  rgin add model product
  rgin add service auth`,
	Args: cobra.ExactArgs(2),
	ValidArgs: []string{"controller", "model", "service"},
	Run: func(cmd *cobra.Command, args []string) {
		componentType := args[0]
		componentName := args[1]
		
		fmt.Printf("Adding new %s: %s\n", componentType, componentName)
		
		options := &generator.ComponentOptions{
			Type: componentType,
			Name: componentName,
		}
		
		if err := generator.AddComponent(options); err != nil {
			cobra.CheckErr(err)
		}
		
		fmt.Printf("✨ Successfully added %s: %s\n", componentType, componentName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
} 