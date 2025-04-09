package takin_log

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"github.com/YangZhaoWeblog/GoldenTakin/takin_log/outputer"
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

// ParseLogLevel 将字符串转换为LogLevel类型
// 支持的字符串格式：debug, info, warn, error, fatal（不区分大小写）
// 如果传入的字符串无法识别，返回默认级别(InfoLevel)和错误
func ParseLogLevel(levelStr string) (LogLevel, error) {
	switch strings.ToLower(strings.TrimSpace(levelStr)) {
	case "debug":
		return DebugLevel, nil
	case "info":
		return InfoLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "error", "err":
		return ErrorLevel, nil
	case "fatal":
		return FatalLevel, nil
	default:
		return InfoLevel, fmt.Errorf("unknown log level: %s, using default level: info", levelStr)
	}
}

type TakinLoggerOptions struct {
	Component string   // 组件名称
	AppName   string   // 应用名称
	MinLevel  LogLevel // 最小日志级别

	Outputs []io.Writer // 日志输出目标列表，支持同时输出到多个目标

	FileLogOption *outputer.FileLogOption // 往文件写入的配置
}

// 定义应用日志接口, 一组行为
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
