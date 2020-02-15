package shell

type get int

func (get) name() string {
  return "get"
}

func (get) description() string {
  return "perform an HTTP GET"
}

func (get) run(args []string) error {
  return nil
}

func init() {
  g := get(0)
  allCommands[g.name()] = g
}
