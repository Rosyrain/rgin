/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"
	"github.com/rosyrain/rgin/cmd"
	"github.com/rosyrain/rgin/internal/template"
)

func main() {
	template.ListEmbedFiles() // 调试用，输出 embed 的所有文件
	// 优先处理全局语言切换
	if len(os.Args) == 3 && (os.Args[1] == "-l" || os.Args[1] == "--lang") {
		lang := os.Args[2]
		cmd.SaveLangConfig(lang) // 假设 SaveLangConfig 是可导出的
		if lang == "en" {
			println("[rgin] Language switched to English (en). All subsequent commands will use English.")
		} else {
			println("[rgin] 语言已切换为中文（zh），后续所有命令将使用中文提示。")
		}
		return
	}
	// 只有 help 命令或 --help 参数时才输出完整帮助
	if len(os.Args) > 1 && (os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-h") {
		cmd.Execute()
		return
	}
	// 其他命令正常执行
	cmd.Execute()
}
