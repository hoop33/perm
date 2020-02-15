package shell

import "fmt"

type unset int

func (unset) name() string {
	return "unset"
}

func (unset) description() string {
	return "unset an environment variable"
}

func (u unset) usage() string {
	return fmt.Sprintf("%s <key>", u.name())
}

func (u unset) run(env *env, args []string) error {
	if len(args) != 1 {
		fmt.Println(u.usage())
		return nil
	}

	delete(env.vars, args[0])
	return nil
}

func init() {
	u := unset(0)
	allCommands[u.name()] = u
}
