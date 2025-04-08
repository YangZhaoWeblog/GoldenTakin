package adapter

import (
	"github.com/YangZhaoWeblog/GoldenTakin/takin_log"
	"github.com/YangZhaoWeblog/GoldenTakin/takin_log/outputer"
	"github.com/go-kratos/kratos/v2/log"
	"testing"
)

func TestKratosAdapter(t *testing.T) {
	// 1. 创建 TakinLogger
	takinLogger := takin_log.NewAppLogger(takin_log.AppLoggerOptions{
		Component: "test-KratosAdapter-to-takinLog",
		AppName:   "test-KratosAdapter-to-takinLog",
		MinLevel:  takin_log.DebugLevel,
		// 直接使用FileLogOption，无需手动创建fileOutput
		FileLogOption: &outputer.FileLogOption{
			FilePath:   "./app_builtin.log",
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     7,
			Compress:   true,
			LocalTime:  true,
		},
	})

	defer func() {
		err := takinLogger.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	// 2. 初始化 Kratos logger, 并将其设置为 kratos log 的适配器
	kratosLogger := NewKratosAdapter(takinLogger)
	log.SetLogger(kratosLogger) // 设置 Kratos 全局日志记录器

	// 3. 使用 kratos logger, 会发现
	log.Log(log.LevelInfo, "kratos消息", "key", "value")
}
