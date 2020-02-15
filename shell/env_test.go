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
	assert.Equal(t, "localhost:3000> ", newEnv().prompt())
}

func TestPromptShouldUpdateWithDomainAndPort(t *testing.T) {
	e := newEnv()
	e.domain = "foo"
	e.port = 12
	assert.Equal(t, "foo:12> ", e.prompt())
}

func TestEnvRunShouldReturnNilWhenEnvIsSelf(t *testing.T) {
	e := newEnv()
	e.vars["foo"] = "bar"
	assert.Nil(t, e.run(e, nil))
}
