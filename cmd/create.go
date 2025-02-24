/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/rosyrain/rgin/internal/generator"
	"github.com/spf13/cobra"
)

var projectName string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create template(by gin)",
	Long:  `rgin create -n projectname`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		if err := generator.GenerateProject(projectName); err != nil {
			cobra.CheckErr(err)
		}
		fmt.Println("Successfully created,projectName:", projectName)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&projectName, "name", "n", "rginDemo", "The name of the create project(default is rginDemo)")
	createCmd.MarkFlagRequired("name")

}
