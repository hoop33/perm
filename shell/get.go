package shell

type get int

func (get) name() string {
  return "get"
}

func (get) run(args []string) error {
  return nil
}

func init() {
  g := get(0)
  commands[g.name()] = g
}
