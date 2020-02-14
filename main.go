package main

import (
	"fmt"
	"os"

	"github.com/hoop33/perm/cmd"
	"github.com/hoop33/perm/config"
)

func main() {
	// Don't launch if we aren't going to be able to save our configuration
	if err := verifyConfigDir(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	cmd.Execute()
}

func verifyConfigDir() error {
	// Get the configuration directory as a string
	dir := config.Dir()

	// If it doesn't exist, create it
	info, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		return os.MkdirAll(dir, 0700)
	}

	// If Stat returned error, bail
	if err != nil {
		return err
	}

	// If it exists, but it's a file, bail
	if !info.IsDir() {
		return fmt.Errorf("config directory '%s' is a file", dir)
	}

	return nil
}
