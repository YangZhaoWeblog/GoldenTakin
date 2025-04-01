package takin_log

import (
	"io"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
)

func TestNewAppLoggerWithKratos(t *testing.T) {

	fileOutput, err := NewFileOutput(&FileLogOption{
		FilePath:   "./app.log", // 日志文件路径
		MaxSize:    100,         // 每个日志文件最大尺寸，单位为MB
		MaxBackups: 3,           // 保留的旧日志文件最大数量
		MaxAge:     7,           // 保留的旧日志文件最大寿命，单位为天
		Compress:   true,        // 压缩旧日志文件
		LocalTime:  true,        // 使用本地时间
	})
	if err != nil {
		t.Fatalf("创建文件日志输出失败: %v", err)
	}

	opts := AppLoggerOptions{
		Component: "test-component",
		AppName:   "test-app",
		//MinLevel:  InfoLevel,
		MinLevel: DebugLevel,
		Outputs:  []io.Writer{fileOutput}, // 注入文件输出
	}

	// 创建 TakinLogger
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
	err = logger.Close()
	if err != nil {
		t.Error(err)
	}
}

// 测试使用内置FileLogOption的开箱即用方式
func TestNewAppLoggerWithFileOption(t *testing.T) {
	// 使用开箱即用方式配置文件日志
	opts := AppLoggerOptions{
		Component: "test-component",
		AppName:   "test-app-builtin",
		MinLevel:  DebugLevel,
		// 直接使用FileLogOption，无需手动创建fileOutput
		FileLogOption: &FileLogOption{
			FilePath:   "./app_builtin.log",
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     7,
			Compress:   true,
			LocalTime:  true,
		},
	}

	// 创建 TakinLogger
	logger := NewAppLoggerWithKratos(opts)

	// 测试日志输出
	logger.Debug("内置FileLogOption方式 - Debug日志")
	logger.Info("内置FileLogOption方式 - Info日志", "key1", "value1")
	logger.Warn("内置FileLogOption方式 - Warn日志")
	logger.Error("内置FileLogOption方式 - Error日志")

	// 关闭日志输出
	err := logger.Close()
	if err != nil {
		t.Error(err)
	}
}
