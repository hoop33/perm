package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "perm",
	Short: "A REPL for HTTP(S)",
	Long: `perm is a REPL for running HTTP and HTTPS commands.
It maintains state across requests, so you type less and get results quicker.`,
	Version: "0.1.0",
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
