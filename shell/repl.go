package shell

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/hoop33/perm/config"
	"github.com/peterh/liner"
)

var errExit = errors.New("exit")

// Repl is the REPL for perm
type Repl struct {
	line *liner.State
	env  *env
}

// NewRepl returns a new Repl
func NewRepl() *Repl {
	return &Repl{
		env: newEnv(),
	}
}

// Run runs the Repl loop
func (r *Repl) Run() error {
	r.createLiner()
	defer r.line.Close()

	r.readHistory()
	r.displayBanner()
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

func (r *Repl) displayBanner() {
	fmt.Printf("%s %s\n", config.Info(config.AppName), config.Info(config.AppVersion))
	fmt.Println(config.Text("Type \"help\" for more information."))
}

func (r *Repl) doLoop() error {
	for {
		if cmdLine, err := r.line.Prompt(r.env.prompt()); err == nil {
			cmds := strings.Split(cmdLine, " ")
			if cmds[0] != "" {
				r.line.AppendHistory(cmdLine)
				if cmd, ok := allCommands[cmds[0]]; ok {
					err := cmd.run(r.env, cmds[1:])
					if err == errExit {
						return nil
					}
					if err != nil {
						fmt.Fprintln(os.Stderr, config.Error(err.Error()))
					}
				} else {
					fmt.Fprintf(os.Stderr, config.Error(fmt.Sprintf("command '%s' not found\n", cmds[0])))
				}
			}
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
	r.line.SetCtrlCAborts(false)
	r.line.SetCompleter(func(line string) (c []string) {
		for _, n := range sortedCommandNames() {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})
}
