/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "rgin",
	Version: "v0.1.0",
	Short:   "A modern scaffolding tool for Gin framework",
	Long: `Rgin is a modern scaffolding tool for the Gin framework that helps you:
  - Create new Gin projects with best practices
  - Add components (controllers, models, services) to your project
  - Include optional features like database support
  - Generate example code to help you get started

To get started, run:
  rgin init myapp

For more examples and documentation, visit:
  https://github.com/rosyrain/rgin`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// 这里不再需要配置文件相关的代码
}
