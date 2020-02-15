package shell

import "fmt"
import "github.com/hoop33/perm/config"

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

func (version) run(_ []string) error {
	fmt.Println(config.AppVersion)
	return nil
}

func init() {
	v := version(0)
	allCommands[v.name()] = v
}
