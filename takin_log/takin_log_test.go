package takin_log

import (
	"github.com/YangZhaoWeblog/GoldenTakin/takin_log/outputer"
	"io"
	"testing"
)

func TestAppLogger(t *testing.T) {
	fileOutput := outputer.NewFileOutput(&outputer.FileLogOption{
		FilePath:   "./app.log", // 日志文件路径
		MaxSize:    100,         // 每个日志文件最大尺寸，单位为MB
		MaxBackups: 3,           // 保留的旧日志文件最大数量
		MaxAge:     7,           // 保留的旧日志文件最大寿命，单位为天
		Compress:   true,        // 压缩旧日志文件
		LocalTime:  true,        // 使用本地时间
	})

	opts := TakinLoggerOptions{
		Component: "test-component",
		AppName:   "test-app",
		//MinLevel:  InfoLevel,
		MinLevel: DebugLevel,
		Outputs:  []io.Writer{fileOutput}, // 注入文件输出
	}

	// 创建 TakinLogger
	takinLogger := NewTakinLogger(opts)

	// 测试基本日志输出
	takinLogger.Debug("这是一条Debug日志")
	takinLogger.Info("这是一条Info日志", "key1", "value1")
	takinLogger.Warn("这是一条Warn日志")
	takinLogger.Error("这是一条Error日志")

	// 测试奇数个参数
	takinLogger.Info("奇数参数测试", "key1", "value1", "key2")
}

// // 测试使用内置FileLogOption的开箱即用方式
func TestAppLoggerWithFileOption(t *testing.T) {
	// 使用开箱即用方式配置文件日志
	opts := TakinLoggerOptions{
		Component: "test-component",
		AppName:   "test-app-builtin",
		MinLevel:  DebugLevel,
		// 直接使用FileLogOption，无需手动创建fileOutput
		FileLogOption: &outputer.FileLogOption{
			FilePath:   "./app_builtin.log",
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     7,
			Compress:   true,
			LocalTime:  true,
		},
	}

	// 创建 TakinLogger
	takinLogger := NewTakinLogger(opts)

	// 测试日志输出
	takinLogger.Debug("内置FileLogOption方式 - Debug日志")
	takinLogger.Info("内置FileLogOption方式 - Info日志", "key1", "value1")
	takinLogger.Warn("内置FileLogOption方式 - Warn日志")
	takinLogger.Error("内置FileLogOption方式 - Error日志")
	takinLogger.Info("奇数参数测试", "key1", "value1", "key2")

	// 关闭日志输出
	err := takinLogger.Close()
	if err != nil {
		t.Error(err)
	}
}
