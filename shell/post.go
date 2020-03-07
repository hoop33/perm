package shell

import (
	"fmt"
	"net/http"
	"strings"
)

type post int

func (post) name() string {
	return "post"
}

func (post) description() string {
	return "perform an HTTP post"
}

func (p post) usage() string {
	return fmt.Sprintf("%s <url>", p.name())
}

func (post) run(env *env, args []string) error {
	url, err := getURL(env, args)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url.String(), strings.NewReader(env.values().Encode()))
	if err != nil {
		return err
	}

	return doRequest(env, req)
}

func init() {
	p := post(0)
	allCommands[p.name()] = p
}
