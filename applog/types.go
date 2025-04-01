package applog

import (
	"context"
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

	FileLogOption *FileLogOption // 往文件写入的配置
	//DBOption      *DBLogOption      // 往 DB 落盘配置
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
