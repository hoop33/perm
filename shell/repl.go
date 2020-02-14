package shell

import (
	"fmt"

	"github.com/peterh/liner"
)

// Repl is the REPL for perm
type Repl struct {
}

// NewRepl returns a new Repl
func NewRepl() *Repl {
	return &Repl{}
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
