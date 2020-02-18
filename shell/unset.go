package shell

import "fmt"

type unset int

func (unset) name() string {
	return "unset"
}

func (unset) description() string {
	return "unset header(s) or query variable(s)"
}

func (u unset) usage() string {
	return fmt.Sprintf("%s <header|var> [key1] [key2] ... (leave blank for all)", u.name())
}

func (u unset) run(env *env, args []string) error {
	if len(args) == 0 {
		u.usage()
		return nil
	}

	switch args[0] {
	case "header":
		if len(args) == 1 {
			env.resetHeaders()
		} else {
			for _, k := range args {
				env.unsetHeader(k)
			}
		}
	case "var":
		if len(args) == 1 {
			env.resetVars()
		} else {
			for _, k := range args {
				env.unsetVar(k)
			}
		}
	default:
		u.usage()
	}

	return nil
}

func init() {
	u := unset(0)
	allCommands[u.name()] = u
}
