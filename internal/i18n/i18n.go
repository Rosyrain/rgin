package i18n

import (
	"fmt"
	"sync"
)

var (
	currentLang = "en"  // 默认英文
	mu          sync.RWMutex
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
	return GetText(key)
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