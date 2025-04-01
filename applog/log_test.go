package applog

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
)

func TestNewAppLoggerWithKratos(t *testing.T) {
	// 创建 AppLoggerOptions
	opts := AppLoggerOptions{
		Component: "test-component",
		AppName:   "test-app",
	}

	// 创建 AppLogger
	logger := NewAppLoggerWithKratos(opts)

	// 测试基本日志输出
	logger.Debug("这是一条Debug日志")
	logger.Info("这是一条Info日志", "key1", "value1")
	logger.Warn("这是一条Warn日志")
	logger.Error("这是一条Error日志")

	// 测试奇数个参数
	logger.Info("奇数参数测试", "key1", "value1", "key2")

	// 测试 Kratos logger
	kratosLogger := log.GetLogger()
	kratosLogger.Log(log.LevelInfo, "kratos消息", "key", "value")
}
