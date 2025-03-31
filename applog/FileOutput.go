package applog

import (
	"io"
)

// 文件日志配置
type FileLogConfig struct {
	Filename   string // 日志文件路径
	MaxSize    int    // 每个日志文件最大尺寸，单位为MB
	MaxBackups int    // 保留的旧日志文件最大数量
	MaxAge     int    // 保留的旧日志文件最大寿命，单位为天
	Compress   bool   // 是否压缩旧日志文件
	LocalTime  bool   // 是否使用本地时间而非UTC时间
}

// 创建一个基于lumberjack的文件输出
// 可以将返回的writer注入到AppLoggerConfig.Output中
func NewFileOutput(config *FileLogConfig) io.Writer {
	if config == nil {
		return nil
	}

	return &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
		LocalTime:  config.LocalTime,
	}
}

// 返回默认的文件日志配置
func DefaultFileLogConfig(filename string) *FileLogConfig {
	return &FileLogConfig{
		Filename:   filename,
		MaxSize:    100,  // 100MB
		MaxBackups: 10,   // 保留10个备份
		MaxAge:     30,   // 保留30天
		Compress:   true, // 压缩旧日志
		LocalTime:  true, // 使用本地时间
	}
}
