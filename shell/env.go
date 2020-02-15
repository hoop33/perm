package shell

import (
	"fmt"
	"sort"
)

type env struct {
	domain string
	port   int
	vars   map[string]string
}

func newEnv() *env {
	e := &env{
		domain: "localhost",
		port:   3000,
		vars:   make(map[string]string),
	}
	allCommands[e.name()] = e

	return e
}

func (e env) prompt() string {
	return fmt.Sprintf("%s:%d> ", e.domain, e.port)
}

func (env) name() string {
	return "env"
}

func (env) description() string {
	return "display the current environment"
}

func (e env) usage() string {
	return e.name()
}

func (e env) run(_ *env, _ []string) error {
	sorted := make([]string, 0, len(e.vars))
	for key := range e.vars {
		sorted = append(sorted, key)
	}
	sort.Strings(sorted)

	for _, key := range sorted {
		fmt.Printf("%s = %s\n", key, e.vars[key])
	}

	return nil
}
