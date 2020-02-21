package shell

import (
	"fmt"

	"github.com/hoop33/perm/config"
)

type exit int

func (exit) name() string {
	return "exit"
}

func (exit) description() string {
	return fmt.Sprintf("exit %s", config.AppName)
}

func (e exit) usage() string {
	return e.name()
}

func (exit) run(_ *env, _ []string) error {
	return errExit
}

func init() {
	e := exit(0)
	allCommands[e.name()] = e
}
