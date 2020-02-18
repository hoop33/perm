package shell

import (
	"fmt"
	"net/url"
	"sort"

	"github.com/hoop33/perm/config"
)

type env struct {
	scheme  string
	host    string
	vars    map[string]string
	headers map[string]string
}

func newEnv() *env {
	e := &env{
		scheme: "http",
		host:   "localhost:3000",
	}
	e.resetVars()
	e.resetHeaders()
	allCommands[e.name()] = e

	return e
}

func (e *env) setVar(key, val string) {
	e.vars[key] = val
}

func (e *env) unsetVar(key string) {
	delete(e.vars, key)
}

func (e *env) resetVars() {
	e.vars = make(map[string]string)
}

func (e *env) setHeader(header, val string) {
	e.headers[header] = val
}

func (e *env) unsetHeader(header string) {
	delete(e.headers, header)
}

func (e *env) resetHeaders() {
	e.headers = make(map[string]string)
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

	// Get the query string and, if it exists, add it to the env vars
	qs, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}

	// TODO multiples
	if len(qs) != 0 {
		for k, v := range qs {
			e.vars[k] = v[0]
		}
	}

	scheme := config.FirstNonBlank(u.Scheme, e.scheme)
	host := config.FirstNonBlank(u.Host, e.host)
	path := config.AddPrefixIfAbsent(u.Path, "/")

	merged, err := url.Parse(fmt.Sprintf("%s://%s%s#%s", scheme, host, path, u.Fragment))
	if err != nil {
		return nil, err
	}

	// TODO multiples
	q := merged.Query()
	for k, v := range e.vars {
		q.Add(k, v)
	}
	merged.RawQuery = q.Encode()

	e.scheme = merged.Scheme
	e.host = merged.Host

	return merged, nil
}
