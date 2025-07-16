package i18n

import (
	"fmt"
	"sync"
)

var (
	currentLang = "zh"
	mu          sync.RWMutex
	texts = map[string]map[string]string{
		"zh": {
			"project_created_success": "项目创建成功！",
			"tip_sqlite": "1. 基础项目依赖 SQLite，首次运行前请手动执行 models/create_table.sql 初始化表结构",
			"tip_sqlite_cmd": "   cd %s",
			"tip_sqlite_exec": "   sqlite3 ./data/app.db < models/create_table.sql",
			"tip_mysql_redis": "2. 已保留 MySQL/Redis 相关连接代码，可按需在 main.go 取消对应注释并在 config 填写相关信息",
			"tip_docs_example": "3. 更多文档和示例请参考 README.md 或参数 --with-example 生成的 example 目录",
			"tip_i18n": "4. 项目已内置中英文国际化支持，rgin init --lang en 可生成英文提示项目",
			"rgin_short": "Gin 脚手架工具",
			"rgin_long": "Rgin 是一个现代化的 Gin 脚手架工具，帮助你：\n  - 快速创建最佳实践的 Gin 项目\n  - 添加控制器、模型、服务等组件\n  - 支持数据库、示例代码等可选特性",
			"flag_lang": "语言(zh/en)",
			"init_short": "初始化 Gin 项目",
			"init_long": "初始化一个新的 Gin 项目，支持可选示例代码。",
			"flag_with_example": "生成包含示例代码的项目结构",
			"help_short": "获取关于 rgin 或特定命令的帮助信息",
			"help_long": "显示 rgin 工具或指定命令的详细帮助信息。",
		},
		"en": {
			"project_created_success": "Project created successfully!",
			"tip_sqlite": "1. The base project depends on SQLite. Please manually execute models/create_table.sql to initialize tables before first run.",
			"tip_sqlite_cmd": "   cd %s",
			"tip_sqlite_exec": "   sqlite3 ./data/app.db < models/create_table.sql",
			"tip_mysql_redis": "2. MySQL/Redis connection code is reserved. You can enable it in main.go and fill in config as needed.",
			"tip_docs_example": "3. For more docs and examples, see README.md or use --with-example to generate the example directory.",
			"tip_i18n": "4. Project supports i18n (zh/en). Use rgin init --lang en to generate an English project.",
			"rgin_short": "Gin scaffolding tool",
			"rgin_long": "Rgin is a modern scaffolding tool for the Gin framework that helps you:\n  - Create new Gin projects with best practices\n  - Add components (controllers, models, services) to your project\n  - Include optional features like database support\n  - Generate example code to help you get started",
			"flag_lang": "Language (zh/en)",
			"init_short": "Initialize a Gin project",
			"init_long": "Initialize a new Gin project with optional example code.",
			"flag_with_example": "Generate project structure with example code",
			"help_short": "Show help for rgin or a specific command",
			"help_long": "Display detailed help for the rgin tool or a specific command.",
		},
	}
)

// 支持的语言
const (
	LangEN = "en"
	LangZH = "zh"
)

// SetLanguage 设置当前语言
func SetLanguage(lang string) error {
	if lang != LangEN && lang != LangZH {
		return fmt.Errorf("unsupported language: %s", lang)
	}
	mu.Lock()
	currentLang = lang
	mu.Unlock()
	return nil
}

// GetText 获取指定key的当前语言文本
func GetText(key string) string {
	mu.RLock()
	lang := currentLang
	mu.RUnlock()
	
	if lang == LangZH {
		return zhTexts[key]
	}
	return enTexts[key]
}

// T 获取文本的简写方法
func T(key string) string {
	if v, ok := texts[currentLang][key]; ok {
		return v
	}
	return key
}

// IsEnglish 判断当前是否为英文
func IsEnglish() bool {
	mu.RLock()
	defer mu.RUnlock()
	return currentLang == LangEN
}

// IsChinese 判断当前是否为中文
func IsChinese() bool {
	mu.RLock()
	defer mu.RUnlock()
	return currentLang == LangZH
} 