package takin_log

import (
	"context"
	"io"
	"log/slog"
	"os"
)

/*
		提供应用级日志功能
	 	内部封装 slog 作为底层写入 log，以替代 kratos log 不支持 InfoContext 等高级特性的缺陷
*/
type TakinLogger struct {
	// 标识日志来源的组件或服务名称, 在微服务架构中用于快速定位日志来源
	// 例如: user-service, video-service
	component string

	// 标识应用名称, 在多应用部署时区分不同应用
	// 例如: tiktok, tiktok-admin
	appName string

	// 表示该条日志的最低级别, 低于该级别的日志不会输出
	minLevel LogLevel

	// 定义日志输出目标列表, 支持同时输出到多个目标
	// 如：- 控制台输出（开发环境）- 文件输出（生产环境）- 数据库存储（错误落盘）
	Outputs []io.Writer // 多个输出，优先级高于单个Output

	// 内部实际使用的日志记录器
	slogLogger *slog.Logger
}

// 创建一个新的应用日志记录器，支持多输出目标
func NewAppLogger(opts AppLoggerOptions) *TakinLogger {
	// 初始化输出目标集合
	writers := []io.Writer{
		os.Stdout, // 默认总是包含控制台输出
	}

	// 添加文件输出（如果配置了文件路径）
	if opts.FileLogOption != nil && opts.FileLogOption.FilePath != "" {
		fileOutput := NewFileOutput(opts.FileLogOption)
		writers = append(writers, fileOutput)
	}

	// 添加外部自定义输出
	if len(opts.Outputs) > 0 {
		writers = append(writers, opts.Outputs...)
	}

	// 确定最低日志级别，默认为 Info
	minLevel := opts.MinLevel
	if minLevel == 0 {
		minLevel = InfoLevel
	}

	// 创建多目标输出writer
	multiWriter := io.MultiWriter(writers...)

	// 创建JSON格式日志处理器
	handler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
		Level: minLevel,
	})

	// 构造并返回日志记录器实例
	return &TakinLogger{
		component:  opts.Component,
		appName:    opts.AppName,
		Outputs:    writers,
		minLevel:   minLevel,
		slogLogger: slog.New(handler),
	}
}

// 构建通用属性
func (l *TakinLogger) commonAttrs() []any {
	return []any{
		"component", l.component,
		"app_name", l.appName,
		//"time", time.Now(),
	}
}

// 添加跟踪信息到日志属性中
func (l *TakinLogger) addTraceInfo(ctx context.Context, attrs []any) []any {
	if traceID := ctx.Value("trace_id"); traceID != nil {
		attrs = append(attrs, "trace_id", traceID)
	}
	if spanID := ctx.Value("span_id"); spanID != nil {
		attrs = append(attrs, "span_id", spanID)
	}
	if userID := ctx.Value("user_id"); userID != nil {
		attrs = append(attrs, "user_id", userID)
	}
	return attrs
}

// 将 普通参数 类型转换为 slog参数 的形式
func slogAttrsFromAny(args []any) []any {
	// slog 只能接受偶数个参数，所以额外添加一个
	if len(args)%2 != 0 {
		args = append(args, "MISSING_VALUE")
	}
	return args
}

func (l *TakinLogger) Debug(msg string, args ...any) {
	attrs := append(l.commonAttrs(), slogAttrsFromAny(args)...)
	l.slogLogger.Debug(msg, attrs...)
}

func (l *TakinLogger) DebugContext(ctx context.Context, msg string, args ...any) {
	attrs := l.addTraceInfo(ctx, append(l.commonAttrs(), slogAttrsFromAny(args)...))
	l.slogLogger.DebugContext(ctx, msg, attrs...)
}

func (l *TakinLogger) Info(msg string, args ...any) {
	attrs := append(l.commonAttrs(), slogAttrsFromAny(args)...)
	l.slogLogger.Info(msg, attrs...)
}

func (l *TakinLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	attrs := l.addTraceInfo(ctx, append(l.commonAttrs(), slogAttrsFromAny(args)...))
	l.slogLogger.InfoContext(ctx, msg, attrs...)
}

func (l *TakinLogger) Warn(msg string, args ...any) {
	attrs := append(l.commonAttrs(), slogAttrsFromAny(args)...)
	l.slogLogger.Warn(msg, attrs...)
}

func (l *TakinLogger) WarnContext(ctx context.Context, msg string, args ...any) {
	attrs := l.addTraceInfo(ctx, append(l.commonAttrs(), slogAttrsFromAny(args)...))
	l.slogLogger.WarnContext(ctx, msg, attrs...)
}

func (l *TakinLogger) Error(msg string, args ...any) {
	attrs := append(l.commonAttrs(), slogAttrsFromAny(args)...)
	l.slogLogger.Error(msg, attrs...)
}

func (l *TakinLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	attrs := l.addTraceInfo(ctx, append(l.commonAttrs(), slogAttrsFromAny(args)...))
	l.slogLogger.ErrorContext(ctx, msg, attrs...)
}

// 关闭所有日志输出
func (l *TakinLogger) Close() error {
	var lastErr error
	for _, output := range l.Outputs {
		// 检查是否支持关闭操作
		if closer, ok := output.(io.Closer); ok {
			if err := closer.Close(); err != nil {
				lastErr = err
			}
		}
	}
	return lastErr
}
