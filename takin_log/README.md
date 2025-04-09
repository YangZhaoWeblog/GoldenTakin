# takin_log
一个高性能、可扩展的日志库，为Go微服务项目提供统一的日志处理方案。

设计理念：
  1. 基于 slog 实现高性能日志处理
  2. 支持多输出目标和结构化日志
  3. 与 Kratos 框架无缝集成

日志系统：
  1. Applogger: 本项目统一使用自定义的 Applogger 来记录日志，而非使用 kratos log(由于 kratos log 不支持 InfoConext 等操作), 而 applogger 是对 slog 的封装，支持更复杂的操作
  2. kratos Adapter:  将 kratos log 最终转化为 Applogger 的输出(这么做是为了，框架内部可能有地方使用的是 kratos log, 此时要保证 kratos log 使用的是 Applogger 的输出)

日志流向：
  1. 业务代码 → TakinLogger → slog → 多个输出目标（文件、控制台、数据库）
  2. Kratos框架 → Kratos log → KratosAdapter → TakinLogger → slog → 多个输出目标（文件、控制台、数据库）

qs:
  1. 为何不使用 kratos helper? helper 是对 kratos logger 的简单封装，但是依然不支持 InfoContext 以及 结构化日志 等高级用法

## 特性

- 基于Go 1.21+ 的`slog`实现
- 多输出目标支持（控制台、文件、自定义）
- 结构化日志
- 支持链路追踪集成
- Kratos框架适配器

## 安装

```go
go get github.com/YangZhaoWeblog/GoldenTakin/takin_log
```

## 快速开始

```go
import (
    "github.com/YangZhaoWeblog/GoldenTakin/takin_log"
    "github.com/YangZhaoWeblog/GoldenTakin/takin_log/outputer"
)

func main() {
    opts := takin_log.AppLoggerOptions{
        Component: "my-service",
        AppName:   "my-app",
        MinLevel:  takin_log.InfoLevel,
        FileLogOption: &outputer.FileLogOption{
            FilePath:   "./logs/app.log",
            MaxSize:    100,
            MaxBackups: 3,
            MaxAge:     7,
            Compress:   true,
        },
    }
    
    logger := takin_log.NewTakinLogger(opts)
    defer logger.Close()
    
    logger.Info("服务启动成功", "port", 8080)
}
```
