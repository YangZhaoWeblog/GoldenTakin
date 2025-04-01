package applog

import (
	"io"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
)

func TestNewAppLoggerWithKratos(t *testing.T) {

	fileOutput := NewFileOutput(&FileLogOption{
		Filename:   "./app.log", // 日志文件路径
		MaxSize:    100,         // 每个日志文件最大尺寸，单位为MB
		MaxBackups: 3,           // 保留的旧日志文件最大数量
		MaxAge:     7,           // 保留的旧日志文件最大寿命，单位为天
		Compress:   true,        // 压缩旧日志文件
		LocalTime:  true,        // 使用本地时间
	})

	opts := AppLoggerOptions{
		Component: "test-component",
		AppName:   "test-app",
		MinLevel:  InfoLevel,
		Outputs:   []io.Writer{fileOutput}, // 注入文件输出
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

	// 关闭日志输出
	err := logger.Close()
	if err != nil {
		return
	}
}
