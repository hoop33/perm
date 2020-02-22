package shell

import (
	"fmt"

	"github.com/hoop33/perm/config"
)

type version int

func (version) name() string {
	return "version"
}

func (version) description() string {
	return "display the application version"
}

func (v version) usage() string {
	return v.name()
}

func (version) run(_ *env, _ []string) error {
	fmt.Println(config.Info(fmt.Sprintf("%s %s", config.AppName, config.AppVersion)))
	return nil
}

func init() {
	v := version(0)
	allCommands[v.name()] = v
}
