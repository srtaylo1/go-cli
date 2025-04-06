package cmd

import (
	format "fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "hello",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		format.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
