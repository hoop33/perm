package shell

// Repl is the REPL for perm
type Repl struct {
}

// NewRepl returns a new Repl
func NewRepl() *Repl {
	return &Repl{
	}
}

// Run runs the Repl loop
func (r *Repl) Run() error {
	return nil
}
