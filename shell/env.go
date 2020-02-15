package shell

import (
	"fmt"
	"github.com/hoop33/perm/config"
	"net/url"
	"sort"
)

type env struct {
	scheme string
	host   string
	vars   map[string]string
}

func newEnv() *env {
	e := &env{
		scheme: "http",
		host:   "localhost:3000",
		vars:   make(map[string]string),
	}
	allCommands[e.name()] = e

	return e
}

func (e *env) prompt() string {
	return fmt.Sprintf("%s://%s> ", e.scheme, e.host)
}

func (*env) name() string {
	return "env"
}

func (*env) description() string {
	return "display the current environment"
}

func (e *env) usage() string {
	return e.name()
}

func (e *env) run(_ *env, _ []string) error {
	sorted := make([]string, 0, len(e.vars))
	for key := range e.vars {
		sorted = append(sorted, key)
	}
	sort.Strings(sorted)

	for _, key := range sorted {
		fmt.Printf("%s = %s\n", key, e.vars[key])
	}

	return nil
}

func (e *env) mergeURL(urlStr string) (*url.URL, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	scheme := config.FirstNonBlank(u.Scheme, e.scheme)
	host := config.FirstNonBlank(u.Host, e.host)
	path := config.AddPrefixIfAbsent(u.Path, "/")

	// TODO build the whole thing
	merged, err := url.Parse(fmt.Sprintf("%s://%s%s", scheme, host, path))
	if err != nil {
		return nil, err
	}

	e.scheme = merged.Scheme
	e.host = merged.Host

	return merged, nil
}
