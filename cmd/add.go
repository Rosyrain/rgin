package cmd

import (
	"fmt"
	"github.com/rosyrain/rgin/internal/generator"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [component] [name]",
	Short: "向项目添加新组件",
	Long: `向现有的 Gin 项目添加新组件。

可用组件类型:
  1. controller - 控制器
     用于处理 HTTP 请求和响应
     示例: rgin add controller user

  2. model - 数据模型
     用于定义数据结构和数据库交互
     示例: rgin add model product

  3. service - 业务服务
     用于实现业务逻辑
     示例: rgin add service auth

使用示例:
  1. 添加用户控制器:
     rgin add controller user
     # 生成 internal/controller/user.go
     # 包含: 基本的 CRUD 操作

  2. 添加产品模型:
     rgin add model product
     # 生成 internal/model/product.go
     # 包含: 结构体定义和基本的数据库操作

  3. 添加认证服务:
     rgin add service auth
     # 生成 internal/service/auth.go
     # 包含: 基本的认证逻辑框架

注意事项:
  - 确保在项目根目录下执行此命令
  - 组件名称应该是有效的 Go 标识符
  - 已存在的组件不会被覆盖（需要手动删除）`,
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