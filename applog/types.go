package applog

import (
	"context"
	"fmt"
	"log/slog"
)

// 日志级别类型
type LogLevel = slog.Level

const (
	DebugLevel = slog.LevelDebug
	InfoLevel  = slog.LevelInfo
	WarnLevel  = slog.LevelWarn
	ErrorLevel = slog.LevelError
	FatalLevel = slog.LevelError + 4 // slog 没有 Fatal 级别，我们自定义一个
)

type AppLoggerOptions struct {
	Component string   // 组件名称
	AppName   string   // 应用名称
	MinLevel  LogLevel // 最小日志级别

	FileLogOption FileLogOption // 往文件写入的配置
	//DBOption      DBLogOption      // 往 DB 落盘配置
}

// 定义应用日志接口
type Logger interface {
	Debug(msg string, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)

	Info(msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)

	Warn(msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)

	Error(msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)

	Close() error
}

// 返回日志级别的字符串表示
func LevelString(level LogLevel) string {
	switch {
	case level <= DebugLevel:
		return "DEBUG"
	case level <= InfoLevel:
		return "INFO"
	case level <= WarnLevel:
		return "WARN"
	case level <= ErrorLevel:
		return "ERROR"
	case level >= FatalLevel:
		return "FATAL"
	default:
		return fmt.Sprintf("LEVEL(%d)", level)
	}
}

// 将 普通参数 类型转换为 slog参数 的形式
func slogAttrsFromAny(args []any) []slog.Attr {
	// slog 只能接受偶数个参数，所以额外添加一个
	if len(args)%2 != 0 {
		args = append(args, "MISSING_VALUE")
	}
	attrs := make([]slog.Attr, 0, len(args)/2)
	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string)
		if !ok {
			key = fmt.Sprintf("%v", args[i])
		}
		attrs = append(attrs, slog.Any(key, args[i+1]))
	}
	return attrs
}
