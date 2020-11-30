package cmd

import (
	"github.com/clh97/ecs/pkg"
	"github.com/clh97/ecs/store"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Runs the API",
	Run: func(cmd *cobra.Command, args []string) {
		store.InitDatabase()
		pkg.InitServer()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
