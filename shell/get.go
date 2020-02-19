package shell

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	urlStr := "/"
	if len(args) > 0 {
		urlStr = args[0]
	}

	url, err := env.mergeURL(urlStr)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return err
	}

	for h, vals := range env.headers {
		req.Header.Set(h, strings.Join(vals, ","))
	}

	resp, err := env.client.Do(req)
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
