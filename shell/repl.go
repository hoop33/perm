package shell

import (
	"fmt"
	"os"

	"github.com/peterh/liner"
)

// Repl is the REPL for perm
type Repl struct {
	configDir string
}

// NewRepl returns a new Repl
func NewRepl(configDir string) (*Repl, error) {
	info, err := os.Stat(configDir)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("config directory '%s' is a file", configDir)
	}

	return &Repl{
		configDir,
	}, nil
}

// Run runs the Repl loop
func (r *Repl) Run() error {
	line := r.createLiner()
	defer line.Close()

	for {
		if cmd, err := line.Prompt("> "); err == nil {
			fmt.Println(cmd)
		} else if err == liner.ErrPromptAborted {
			return nil
		} else {
			return err
		}
	}
}

func (r *Repl) createLiner() *liner.State {
	line := liner.NewLiner()
	line.SetCtrlCAborts(true)
	return line
}
