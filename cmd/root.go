package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/clh97/ecs/config"
)

var (
	cfgFile     string
	userLicense string
	rootCmd     = &cobra.Command{
		Use:   "ecs",
		Short: "A rest api for ECS",
		Long:  "ECS is a rest api that serves as embedded comment system for JAMStack applications",
	}
)

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(func() { config.Initialize(cfgFile) })
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ecs.config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "MIT")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
}
