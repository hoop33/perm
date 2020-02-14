package shell

import (
	"fmt"
	"io"
	"os"

	"github.com/hoop33/perm/config"
	"github.com/peterh/liner"
)

// Repl is the REPL for perm
type Repl struct {
	line *liner.State
}

// NewRepl returns a new Repl
func NewRepl() *Repl {
	return &Repl{}
}

// Run runs the Repl loop
func (r *Repl) Run() error {
	r.createLiner()
	defer r.line.Close()

	r.readHistory()
	err := r.doLoop()
	if err2 := r.saveHistory(); err2 != nil {
		fmt.Fprintln(os.Stderr, err2)
	}

	return err
}

func (r *Repl) readHistory() {
	// Don't worry if we can't read history
	if f, err := os.Open(config.HistoryFile()); err == nil {
		r.line.ReadHistory(f)
		f.Close()
	}
}

func (r *Repl) doLoop() error {
	for {
		if cmd, err := r.line.Prompt("> "); err == nil {
			r.line.AppendHistory(cmd)
			fmt.Println(cmd)
		} else if err == liner.ErrPromptAborted || err == io.EOF {
			return nil
		} else {
			return err
		}
	}
}

func (r *Repl) saveHistory() error {
	f, err := os.Create(config.HistoryFile())
	if err != nil {
		return err
	}

	r.line.WriteHistory(f)
	f.Close()
	return nil
}

func (r *Repl) createLiner() {
	r.line = liner.NewLiner()
	r.line.SetCtrlCAborts(true)
}
