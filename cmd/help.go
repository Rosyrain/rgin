package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/rosyrain/rgin/internal/i18n"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help [command]",
	Short: i18n.T("help_short"),
	Long:  i18n.T("help_long"),
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