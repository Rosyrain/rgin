package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help [command]",
	Short: "获取关于 rgin 或特定命令的帮助信息",
	Long: `Rgin 是一个现代化的 Gin 框架脚手架工具。

基本使用方法:
  1. 创建新项目:
     rgin init myapp              # 创建基础项目
     rgin init myapp --with-db sqlite    # 使用 SQLite 数据库
     rgin init myapp --with-examples     # 包含示例代码

  2. 添加组件:
     rgin add controller user     # 添加用户控制器
     rgin add model product       # 添加产品模型
     rgin add service auth        # 添加认证服务

项目结构说明:
  cmd/          - 命令行入口
  internal/     - 内部代码
    controller/ - 控制器
    model/      - 数据模型
    service/    - 业务服务
  pkg/          - 公共包
  api/          - API 文档
  config/       - 配置文件

常见问题:
  Q: 如何选择数据库？
  A: 使用 --with-db 参数：
     - sqlite (默认，开箱即用)
     - mysql  (需要额外配置)
     - none   (无数据库)

  Q: 在哪里找到示例代码？
  A: 使用 --with-examples 参数生成示例代码，
     示例代码将保存在 examples/ 目录下

  Q: 如何自定义模板？
  A: 模板文件位于 templates/ 目录，
     可以根据需要修改或添加新模板

更多信息:
  文档: https://github.com/rosyrain/rgin/blob/main/README.md
  问题反馈: https://github.com/rosyrain/rgin/issues`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cmd.Long)
			return
		}
		// 如果指定了具体命令，显示该命令的帮助
		if c, _, err := cmd.Root().Find(args); err == nil {
			c.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
} 