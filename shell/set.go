package shell

import (
	"fmt"
	"strconv"
)

type set int

func (set) name() string {
	return "set"
}

func (set) description() string {
	return "set an environment variable"
}

func (s set) usage() string {
	return fmt.Sprintf("%s <key> <value>", s.name())
}

func (s set) run(env *env, args []string) error {
	if len(args) != 2 {
		fmt.Println(s.usage())
		return nil
	}

	// TODO make this cleaner
	switch args[0] {
	case "domain":
		env.domain = args[1]
	case "port":
		if port, err := strconv.Atoi(args[1]); err == nil {
			env.port = port
		}
	default:
		env.vars[args[0]] = args[1]
	}

	return nil
}

func init() {
	s := set(0)
	allCommands[s.name()] = s
}
