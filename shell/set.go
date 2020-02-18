package shell

import (
	"fmt"
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
	switch len(args) {
	case 1:
		env.vars[args[0]] = "true"
	case 2:
		env.vars[args[0]] = args[1]
	default:
		s.usage()
	}

	return nil
}

func init() {
	s := set(0)
	allCommands[s.name()] = s
}
