package config

import (
	config "github.com/spf13/viper"
)

func init() {
	config.SetDefault("logger.level", "info")
	config.SetDefault("logger.encoding", "console")
	config.SetDefault("logger.color", true)
	config.SetDefault("logger.dev_mode", true)
	config.SetDefault("logger.disable_caller", false)
	config.SetDefault("logger.disable_stacktrace", true)

	config.SetDefault("server.host", "0.0.0.0")
	config.SetDefault("server.port", "3000")
}
