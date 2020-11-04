package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("logger.level", "info")
	viper.SetDefault("logger.encoding", "console")
	viper.SetDefault("logger.color", true)
	viper.SetDefault("logger.dev_mode", true)
	viper.SetDefault("logger.disable_caller", false)
	viper.SetDefault("logger.disable_stacktrace", true)

	viper.SetDefault("server.addr", "0.0.0.0")
	viper.SetDefault("server.port", 3000)

	viper.SetDefault("author", "Michel Calheiros <michel@calheiros.dev")
	viper.SetDefault("license", "MIT")
}
