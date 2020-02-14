package shell

type command interface {
	name() string
	run(args []string) error
}

var commands = make(map[string]command)
