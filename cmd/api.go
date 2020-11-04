package cmd

import (
	"github.com/clh97/ecs/internal"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Runs the API",
	Run: func(cmd *cobra.Command, args []string) {
		internal.InitServer()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}