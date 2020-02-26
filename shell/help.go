package shell

import (
	"fmt"

	"github.com/hoop33/perm/config"
)

type help int

func (help) name() string {
	return "help"
}

func (help) description() string {
	return "show help"
}

func (h help) usage() string {
	return h.name()
}

func (help) run(_ *env, _ []string) error {
	fmt.Println(config.Info(fmt.Sprintf("%s %s", config.AppName, config.AppVersion)))
	fmt.Println(config.Default("a REPL for HTTP(S)"))
	fmt.Println(config.Default("\nEnter a command at the prompt"))
	fmt.Println(config.Header("\nCommands:"))
	commands(0).run(nil, nil)
	return nil
}

func init() {
	h := help(0)
	allCommands[h.name()] = h
}
