package shell

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type get int

func (get) name() string {
	return "get"
}

func (get) description() string {
	return "perform an HTTP GET"
}

func (g get) usage() string {
	return fmt.Sprintf("%s <url>", g.name())
}

func (g get) run(_ *env, args []string) error {
	// TODO centralize client, make it configurable over the life of the program
	if len(args) == 0 {
		fmt.Println(g.usage())
		return nil
	}

	resp, err := http.Get(args[0])
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}

func init() {
	g := get(0)
	allCommands[g.name()] = g
}
