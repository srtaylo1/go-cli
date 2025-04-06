package cmd

import (
	format "fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "hello",
	Run: func(cmd *cobra.Command, args []string) {
		format.Fprintln(os.Stderr, "hello World")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		format.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
