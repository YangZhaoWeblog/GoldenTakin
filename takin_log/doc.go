// internal/log/log.go

package takin_log

/*
	日志系统：
		1. Applogger: 本项目统一使用自定义的 Applogger 来记录日志，而非使用 kratos log(由于 kratos log 不支持 InfoConext 等操作)
			而 applogger 是对 slog 的封装，支持更复杂的操作
		2. kratos Adapter:  将 kratos log 最终转化为 Applogger 的输出(这么做是为了，框架内部可能有地方使用的是 kratos log, 此时要保证 kratos log 使用的是 Applogger 的输出)
			日志流向：
				业务代码 → TakinLogger → slog → 多个输出目标（文件、控制台、数据库）
					↑
				Kratos框架 → Kratos log → KratosAdapter
	qs:
		1. 为何不使用 kratos helper? helper 是对 kratos logger 的简单封装，但是依然不支持 InfoContext 以及 结构化日志 等高级用法
*/
