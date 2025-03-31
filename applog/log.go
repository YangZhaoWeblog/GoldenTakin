// internal/log/log.go

package applog

import (
	"io"

	"github.com/go-kratos/kratos/v2/log"
)

/*
	日志系统：
		1. Applogger: 本项目统一使用自定义的 Applogger 来记录日志，而非使用 kratos log(由于 kratos log 不支持 InfoConext 等操作)
			而 applogger 是对 slog 的封装，支持更复杂的操作
		2. kratos Adapter:  将 kratos log 最终转化为 Applogger 的输出(这么做是为了，框架内部可能有地方使用的是 kratos log, 此时要保证 kratos log 使用的是 Applogger 的输出)
			日志流向：
				业务代码 → AppLogger → slog → 多个输出目标（文件、控制台、数据库）
					↑
				Kratos框架 → Kratos log → KratosAdapter
	qs:
		1. 为何不使用 kratos helper? helper 是对 kratos logger 的简单封装，但是依然不支持 InfoContext 以及 结构化日志 等高级用法
*/

// 创建应用日志系统的核心组件，并配置 Kratos 日志适配
func NewAppLoggerWithKratos(options AppLoggerOptions, option FileLogOption) *AppLogger {
	// 创建文件输出并添加到选项中
	if options.Outputs == nil {
		options.Outputs = []io.Writer{fileOutput}
	} else {
		options.Outputs = append(options.Outputs, fileOutput)
	}

	// 创建应用日志记录器
	appLogger := NewAppLogger(options)

	// 创建 Kratos 适配器
	kratosLogger := NewKratosAdapter(appLogger)

	// 设置 Kratos 全局日志记录器
	log.SetLogger(kratosLogger)

	return appLogger
}
