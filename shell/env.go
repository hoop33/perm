package shell

import (
	"fmt"
	"sort"
)

type env struct {
	vars map[string]string
}

func newEnv() *env {
	e := &env{
		vars: make(map[string]string),
	}
	allCommands[e.name()] = e

	return e
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
