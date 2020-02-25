package shell

import (
	"fmt"
	"strings"

	"github.com/hoop33/perm/config"
)

type color int

func (color) name() string {
	return "color"
}

func (color) description() string {
	return "turn color output on or off"
}

func (c color) usage() string {
	return fmt.Sprintf("%s [on|off]", c.name())
}

func (color) run(_ *env, args []string) error {
	enable := len(args) == 0 || "on" == strings.ToLower(args[0]) || "true" == strings.ToLower(args[0])
	config.EnableColor(enable)
	return nil
}

func init() {
	c := color(0)
	allCommands[c.name()] = c
}
