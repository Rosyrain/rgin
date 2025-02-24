/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rgin",
	Short: "Generate the development scaffolding of the gin framework",
	Long: `Here's how to use it:
	go install https://github.com/rosyrain/rgin
	rgin create -n projectname`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("==========\nWelcome to rgin,use `rgin create -n projectName` to create project.\n==========")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	//err := rootCmd.Execute()
	//if err != nil {
	//	os.Exit(1)
	//}

	cobra.CheckErr(rootCmd.Execute())

}

func init() {
	//cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $GOPATH/bin/rgin/conf/config.yaml.tmpl)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		cfgFile = "./conf/config.yaml.tmpl"
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

}
