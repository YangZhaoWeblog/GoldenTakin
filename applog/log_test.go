package applog

import (
	"testing"
)

// 测试 NewAppLoggerWithKratos 功能
func TestNewAppLoggerWithKratos(t *testing.T) {
	// 创建 AppLoggerOptions
	opts := AppLoggerOptions{
		Component: "test-component",
		AppName:   "test-app",
		MinLevel:  DebugLevel,
	}

	// 创建 AppLogger
	logger := NewAppLoggerWithKratos(opts)

	// 输出不同级别的日志
	logger.Debug("这是一条Debug日志")
	logger.Info("这是一条Info日志")
	logger.Warn("这是一条Warn日志")
	logger.Error("这是一条Error日志")
}
