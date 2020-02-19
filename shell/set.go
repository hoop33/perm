package shell

import (
	"fmt"
)

type set int

func (set) name() string {
	return "set"
}

func (set) description() string {
	return "set a header or query variable"
}

func (s set) usage() string {
	return fmt.Sprintf("%s <header|var> <key> <value>", s.name())
}

func (s set) run(env *env, args []string) error {
	// Need at least 2 values: <header|var> key [value]
	if len(args) < 2 {
		s.usage()
		return nil
	}

	switch args[0] {
	case "header":
		env.setHeader(args[1], args[2:]...)
	case "var":
		env.setVar(args[1], args[2:]...)
	default:
		s.usage()
	}

	return nil
}

func init() {
	s := set(0)
	allCommands[s.name()] = s
}
