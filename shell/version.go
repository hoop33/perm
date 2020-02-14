package shell

import "fmt"
import "github.com/hoop33/perm/config"

type version int

func (version) name() string {
	return "version"
}

func (version) run(_ []string) error {
	fmt.Println(config.AppVersion)
	return nil
}

func init() {
	v := version(0)
	commands[v.name()] = v
}
