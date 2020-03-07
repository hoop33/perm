package shell

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getURL(env *env, args []string) (*url.URL, error) {
	// Allow them to not pass a URL at all
	urlStr := "/"
	if len(args) > 0 {
		urlStr = args[0]
	}

	return env.mergeURL(urlStr)
}

func doRequest(env *env, req *http.Request) error {
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
