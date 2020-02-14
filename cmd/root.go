package cmd

import (
	"fmt"
	"os"

	"github.com/hoop33/perm/config"
	"github.com/hoop33/perm/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   config.AppName,
	Short: "A REPL for HTTP(S)",
	Long: `perm is a REPL for running HTTP and HTTPS commands.
It maintains state across requests, so you type less and get results quicker.`,
	Version: config.AppVersion,
	Run: func(cmd *cobra.Command, args []string) {
		repl, err := shell.NewRepl(config.Dir())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := repl.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
