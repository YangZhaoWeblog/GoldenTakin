package adapter

import (
	"github.com/YangZhaoWeblog/GoldenTakin/takin_log"
	"github.com/go-kratos/kratos/v2/log"
)

/*
		用于将 kratos log，转化为调用 Applogger
	 	kratos 可设置 kratosAdapter
*/
type KratosAdapter struct {
	logger *takin_log.TakinLogger
}

func NewKratosAdapter(logger *takin_log.TakinLogger) *KratosAdapter {
	return &KratosAdapter{logger: logger}
}

// 实现 kratos 的 log.Logger 接口
func (a *KratosAdapter) Log(level log.Level, keyvals ...interface{}) error {
	// 从 keyvals 中提取消息和其他属性
	var msg string
	attrs := make([]interface{}, 0, len(keyvals))

	// 检查是否有消息
	if len(keyvals) > 0 {
		if v, ok := keyvals[0].(string); ok {
			msg = v
			keyvals = keyvals[1:]
		}
	}

	// 如果没有明确的消息，生成一个默认消息
	if msg == "" {
		msg = "default kratos log"
	}

	// 将剩余的 keyvals 视为属性
	attrs = append(attrs, keyvals...)

	// 根据 Kratos 日志级别调用相应的 TakinLogger 方法
	// 循环拆解 key, values 的责任下推到 Applogger 来执行
	switch level {
	case log.LevelDebug:
		a.logger.Debug(msg, attrs...)
	case log.LevelInfo:
		a.logger.Info(msg, attrs...)
	case log.LevelWarn:
		a.logger.Warn(msg, attrs...)
	case log.LevelError:
		a.logger.Error(msg, attrs...)
	default:
		a.logger.Info(msg, attrs...)
	}

	return nil
}
