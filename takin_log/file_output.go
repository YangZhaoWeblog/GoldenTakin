package takin_log

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

// 文件日志配置
type FileLogOption struct {
	FilePath   string // 日志文件路径
	MaxSize    int    // 每个日志文件最大尺寸，单位为MB
	MaxBackups int    // 保留的旧日志文件最大数量
	MaxAge     int    // 保留的旧日志文件最大寿命，单位为天
	Compress   bool   // 是否压缩旧日志文件
	LocalTime  bool   // 是否使用本地时间而非UTC时间
}

// 创建一个基于lumberjack的文件输出
// 可以将返回的writer注入到AppLoggerConfig.Output中
func NewFileOutput(config *FileLogOption) io.Writer {
	// 提供默认值
	maxSize := config.MaxSize
	if maxSize <= 0 {
		maxSize = 100 // 默认100MB
	}

	maxBackups := config.MaxBackups
	if maxBackups <= 0 {
		maxBackups = 3 // 默认3个备份
	}

	MaxAge := config.MaxAge
	if MaxAge <= 0 {
		MaxAge = 7 // 默认保留 7 天
	}

	return &lumberjack.Logger{
		Filename:   config.FilePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     MaxAge,
		Compress:   config.Compress,
	}
}
