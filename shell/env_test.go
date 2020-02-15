package shell

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnvNameShouldReturnEnv(t *testing.T) {
	assert.Equal(t, "env", newEnv().name())
}

func TestEnvsDescriptionShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", newEnv().description())
}

func TestEnvUsageShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", newEnv().usage())
}

func TestEnvRunShouldReturnNil(t *testing.T) {
	assert.Nil(t, newEnv().run(nil, nil))
}

func TestEnvShouldRegisterItself(t *testing.T) {
	_, ok := allCommands[newEnv().name()]
	assert.True(t, ok)
}

func TestPromptShouldDefaultToLocalhost3000(t *testing.T) {
	assert.Equal(t, "http://localhost:3000> ", newEnv().prompt())
}

func TestPromptShouldUpdateWithSchemeAndHost(t *testing.T) {
	e := newEnv()
	e.scheme = "https"
	e.host = "foo"
	assert.Equal(t, "https://foo> ", e.prompt())
}

func TestEnvRunShouldReturnNilWhenEnvIsSelf(t *testing.T) {
	e := newEnv()
	e.vars["foo"] = "bar"
	assert.Nil(t, e.run(e, nil))
}

func TestEnvMergeURLShouldReturnTheSameURLWhenURLIsAbsolute(t *testing.T) {
	e := newEnv()
	e.scheme = "http"
	e.host = "localhost:3000"
	url, err := e.mergeURL("https://grailbox.com")
	assert.Nil(t, err)
	assert.Equal(t, "https://grailbox.com/", url.String())
}

func TestEnvMergeURLShouldChangeTheBaseURLWhenURLIsAbsolute(t *testing.T) {
	e := newEnv()
	e.scheme = "http"
	e.host = "localhost:3000"
	_, err := e.mergeURL("https://grailbox.com")
	assert.Nil(t, err)
	assert.Equal(t, "https", e.scheme)
	assert.Equal(t, "grailbox.com", e.host)
}

func TestEnvMergeURLShouldReturnTheSameURLWhenURLIsAbsoluteAndHasPath(t *testing.T) {
	e := newEnv()
	e.scheme = "http"
	e.host = "localhost:3000"
	url, err := e.mergeURL("https://grailbox.com/uses")
	assert.Nil(t, err)
	assert.Equal(t, "https://grailbox.com/uses", url.String())
}

func TestEnvMergeURLShouldChangeTheBaseURLWhenURLIsAbsoluteAndHasPath(t *testing.T) {
	e := newEnv()
	e.scheme = "http"
	e.host = "localhost:3000"
	_, err := e.mergeURL("https://grailbox.com/uses")
	assert.Nil(t, err)
	assert.Equal(t, "https", e.scheme)
	assert.Equal(t, "grailbox.com", e.host)
}

func TestEnvMergeURLShouldReturnMergedURLWhenURLIsRelative(t *testing.T) {
	e := newEnv()
	e.scheme = "http"
	e.host = "localhost:3000"
	url, err := e.mergeURL("/")
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost:3000/", url.String())
}

func TestEnvMergeURLShouldRetainTheBaseURLWhenURLIsRelative(t *testing.T) {
	e := newEnv()
	e.scheme = "http"
	e.host = "localhost:3000"
	_, err := e.mergeURL("/")
	assert.Nil(t, err)
	assert.Equal(t, "http", e.scheme)
	assert.Equal(t, "localhost:3000", e.host)
}

func TestEnvMergeURLShouldReturnMergedURLWhenURLIsRelativeAndHasPath(t *testing.T) {
	e := newEnv()
	e.scheme = "http"
	e.host = "localhost:3000"
	_, err := e.mergeURL("/foo/bar")
	assert.Nil(t, err)
	assert.Equal(t, "http", e.scheme)
	assert.Equal(t, "localhost:3000", e.host)
}

func TestEnvMergeURLShouldRetainTheBaseURLWhenURLIsRelativeAndHasPath(t *testing.T) {
	e := newEnv()
	e.scheme = "http"
	e.host = "localhost:3000"
	_, err := e.mergeURL("/foo/bar")
	assert.Nil(t, err)
	assert.Equal(t, "http", e.scheme)
	assert.Equal(t, "localhost:3000", e.host)
}

func TestEnvMergeURLShouldReturnMergedURLWhenURLIsRelativeAndHasPathNoLeadingSlash(t *testing.T) {
	e := newEnv()
	e.scheme = "http"
	e.host = "localhost:3000"
	_, err := e.mergeURL("foo/bar")
	assert.Nil(t, err)
	assert.Equal(t, "http", e.scheme)
	assert.Equal(t, "localhost:3000", e.host)
}

func TestEnvMergeURLShouldRetainTheBaseURLWhenURLIsRelativeAndHasPathNoLeadingSlash(t *testing.T) {
	e := newEnv()
	e.scheme = "http"
	e.host = "localhost:3000"
	_, err := e.mergeURL("foo/bar")
	assert.Nil(t, err)
	assert.Equal(t, "http", e.scheme)
	assert.Equal(t, "localhost:3000", e.host)
}
