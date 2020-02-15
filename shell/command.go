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
var sorted []string
var maxCommandLength int

func sortedCommandNames() []string {
  if len(sorted) == 0 {
    sorted = make([]string, 0, len(allCommands))
    for name := range allCommands {
      sorted = append(sorted, name)
    }
    sort.Strings(sorted)
  }
  return sorted
}

func maxLen() int {
  if maxCommandLength == 0 {
    for _, name := range sortedCommandNames() {
      if len(name) > maxCommandLength {
        maxCommandLength = len(name)
      }
    }
  }
  return maxCommandLength
}

type commands int

func (commands) name() string {
  return "commands"
}

func (commands) description() string {
  return "list available commands"
}

func (commands) run(_ []string) error {
  for _, name := range sortedCommandNames() {
    fmt.Printf("%-*s  %s\n", maxLen(), name, allCommands[name].description())
  }
  return nil
}

func init() {
  c := commands(0)
  allCommands[c.name()] = c
}
