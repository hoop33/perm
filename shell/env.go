package shell

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/hoop33/perm/config"
)

type env struct {
	scheme  string
	host    string
	vars    map[string][]string
	headers map[string][]string
	client  *http.Client
}

func newEnv() *env {
	e := &env{
		scheme: "http",
		host:   "localhost:3000",
	}
	e.resetVars()
	e.resetHeaders()
	e.resetClient()
	allCommands[e.name()] = e

	return e
}

func (e *env) setVar(key string, vals ...string) {
	if len(vals) == 0 {
		e.vars[key] = []string{"true"}
	} else {
		e.vars[key] = vals
	}
}

func (e *env) unsetVar(key string) {
	delete(e.vars, key)
}

func (e *env) resetVars() {
	e.vars = make(map[string][]string)
}

func (e *env) resetClient() {
	// TODO configuration
	e.client = &http.Client{
		Timeout: time.Second * 30,
	}
}

func (e *env) setHeader(header string, vals ...string) {
	if len(vals) == 0 {
		e.headers[header] = []string{"true"}
	} else {
		e.headers[header] = vals
	}
}

func (e *env) unsetHeader(header string) {
	delete(e.headers, header)
}

func (e *env) resetHeaders() {
	e.headers = make(map[string][]string)
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
	if len(e.headers) > 0 {
		fmt.Println(config.Header("Headers:"))
		e.showHeaders()
	}
	if len(e.vars) > 0 {
		fmt.Println(config.Header("Variables:"))
		e.showVars()
	}

	return nil
}

func (e *env) showHeaders() {
	showSorted(e.headers)
}

func (e *env) showVars() {
	showSorted(e.vars)
}

func showSorted(m map[string][]string) {
	sorted := make([]string, 0, len(m))
	for key := range m {
		sorted = append(sorted, key)
	}
	sort.Strings(sorted)

	for _, key := range sorted {
		fmt.Printf("  %s = %s\n", config.Info(key), config.Text(strings.Join(m[key], ", ")))
	}
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

	if len(qs) != 0 {
		for k, v := range qs {
			e.setVar(k, v...)
		}
	}

	scheme := config.FirstNonBlank(u.Scheme, e.scheme)
	host := config.FirstNonBlank(u.Host, e.host)
	path := config.AddPrefixIfAbsent(u.Path, "/")

	merged, err := url.Parse(fmt.Sprintf("%s://%s%s#%s", scheme, host, path, u.Fragment))
	if err != nil {
		return nil, err
	}

	q := merged.Query()
	for k, vals := range e.vars {
		for _, v := range vals {
			q.Add(k, v)
		}
	}
	merged.RawQuery = q.Encode()

	e.scheme = merged.Scheme
	e.host = merged.Host

	return merged, nil
}
