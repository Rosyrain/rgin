/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"gopkg.in/yaml.v2"
	"github.com/rosyrain/rgin/internal/i18n"
)

var rootCmd = &cobra.Command{
	Use:   "rgin",
	Short: i18n.T("rgin_short"),
	Long:  i18n.T("rgin_long"),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

var lang string
var configPath = filepath.Join(os.Getenv("HOME"), ".rgin.yaml")

type RginConfig struct {
	Lang string `yaml:"lang"`
}

func loadLangConfig() string {
	cfg := RginConfig{Lang: "zh"}
	f, err := os.Open(configPath)
	if err == nil {
		yaml.NewDecoder(f).Decode(&cfg)
		f.Close()
	}
	return cfg.Lang
}

// 将 saveLangConfig 导出
func SaveLangConfig(lang string) {
	cfg := RginConfig{Lang: lang}
	f, err := os.Create(configPath)
	if err == nil {
		yaml.NewEncoder(f).Encode(cfg)
		f.Close()
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&lang, "lang", "l", "", i18n.T("flag_lang"))
	// 这里不再需要配置文件相关的代码
}

// 在 main 或各命令执行前
// i18n.SetLang(lang)

func Execute() {
	// 1. 读取全局配置
	globalLang := loadLangConfig()
	// 2. 如果命令行参数指定了 -l/--lang，则覆盖全局配置
	if lang != "" {
		globalLang = lang
		SaveLangConfig(lang)
		if len(os.Args) == 2 && (os.Args[1] == "-l" || os.Args[1] == "--lang") {
			println("[rgin] 请输入语言参数，如 rgin -l zh 或 rgin -l en")
			return
		}
		if len(os.Args) == 3 && (os.Args[1] == "-l" || os.Args[1] == "--lang") {
			if lang == "en" {
				println("[rgin] Language switched to English (en). All subsequent commands will use English.")
			} else {
				println("[rgin] 语言已切换为中文（zh），后续所有命令将使用中文提示。")
			}
			return // 只切换语言时不再执行 rootCmd.Execute()
		}
	}
	i18n.SetLanguage(globalLang)
	rootCmd.Execute()
}
