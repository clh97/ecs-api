package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Initialize reads up config file passed by cli or tries to read in current dir
func Initialize(cfgFile string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName("ecs.viper")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
