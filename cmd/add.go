package cmd

import (
	"fmt"
	"github.com/rosyrain/rgin/internal/generator"
	"github.com/rosyrain/rgin/internal/i18n"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [component] [name]",
	Short: "Add a component to the project", // 将由 i18n 系统更新
	Long:  "Loading...", // 将由 i18n 系统更新
	Args: cobra.ExactArgs(2),
	ValidArgs: []string{"controller", "model", "service"},
	Run: func(cmd *cobra.Command, args []string) {
		componentType := args[0]
		componentName := args[1]
		
		if i18n.IsChinese() {
			fmt.Printf("正在添加%s：%s\n", getComponentTypeCN(componentType), componentName)
		} else {
			fmt.Printf("Adding new %s: %s\n", componentType, componentName)
		}
		
		options := &generator.ComponentOptions{
			Type: componentType,
			Name: componentName,
		}
		
		if err := generator.AddComponent(options); err != nil {
			cobra.CheckErr(err)
		}
		
		if i18n.IsChinese() {
			fmt.Printf("✨ %s添加成功：%s\n", getComponentTypeCN(componentType), componentName)
		} else {
			fmt.Printf("✨ Successfully added %s: %s\n", componentType, componentName)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

// getComponentTypeCN 获取组件类型的中文名称
func getComponentTypeCN(componentType string) string {
	switch componentType {
	case "controller":
		return "控制器"
	case "model":
		return "模型"
	case "service":
		return "服务"
	default:
		return componentType
	}
} 