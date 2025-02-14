package api

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "consumer-cli",
	Short: "Cli manager",
}

func Execute() error {
	return rootCmd.Execute()
}
