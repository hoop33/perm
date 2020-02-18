package shell

import "fmt"

type unset int

func (unset) name() string {
	return "unset"
}

func (unset) description() string {
	return "unset environment variable(s)"
}

func (u unset) usage() string {
	return fmt.Sprintf("%s [var1] [var2] ... (leave blank for all)", u.name())
}

func (u unset) run(env *env, args []string) error {
	if len(args) == 0 {
		env.resetVars()
	} else {
		for _, k := range args {
			delete(env.vars, k)
		}
	}

	return nil
}

func init() {
	u := unset(0)
	allCommands[u.name()] = u
}
