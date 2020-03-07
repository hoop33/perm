package shell

import (
	"fmt"
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

func (g get) run(env *env, args []string) error {
	url, err := getURL(env, args)
	if err != nil {
		return err
	}

	url.RawQuery = env.values().Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return err
	}

	return doRequest(env, req)
}

func init() {
	g := get(0)
	allCommands[g.name()] = g
}
