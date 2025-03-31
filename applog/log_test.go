package applog

func main() {
	// 配置 AppLoggerOptions
	opts := AppLoggerOptions{
		Level:  "info",
		Format: "json",
	}

	// 创建 AppLogger
	logger := applog.NewAppLoggerWithKratos(opts)

	// ... 使用 logger ...
	logger.Info("Server started")
}
