package shell

import (
  "fmt"
  "sort"
)

type command interface {
	name() string
	description() string
	run(args []string) error
}

var allCommands = make(map[string]command)

type commands int

func (commands) name() string {
  return "commands"
}

func (commands) description() string {
  return "list available commands"
}

func (commands) run(_ []string) error {
  maxLen := 0
  names := make([]string, 0, len(allCommands))
  for name := range allCommands {
    names = append(names, name)
    if len(name) > maxLen {
      maxLen = len(name)
    }
  }
  sort.Strings(names)

  for _, name := range names {
    fmt.Printf("%-*s  %s\n", maxLen, name, allCommands[name].description())
  }

  return nil
}

func init() {
  c := commands(0)
  allCommands[c.name()] = c
}
