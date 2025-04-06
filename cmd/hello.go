package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(helloCommand)
}

var helloCommand = &cobra.Command{
	Use: "hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(os.Stdout, *cmd, args)
	},
}
