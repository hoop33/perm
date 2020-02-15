package shell

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetNameShouldReturnSet(t *testing.T) {
	assert.Equal(t, "set", set(0).name())
}

func TestSetsDescriptionShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", set(0).description())
}

func TestSetUsageShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", set(0).usage())
}

func TestSetRunShouldReturnNil(t *testing.T) {
	assert.Nil(t, set(0).run(nil, nil))
}

func TestSetShouldRegisterItself(t *testing.T) {
	_, ok := allCommands[set(0).name()]
	assert.True(t, ok)
}

func TestSetShouldSetEnvVarWhenTwoArgs(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"foo", "bar"}))
	assert.Equal(t, "bar", e.vars["foo"])
}

func TestSetShouldNotSetEnvVarWhenOneArg(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"foo"}))
	_, ok := e.vars["foo"]
	assert.False(t, ok)
}

func TestSetShouldNotSetEnvVarWhenNoArgs(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{}))
	assert.Equal(t, 0, len(e.vars))
}

func TestSetDomainShouldSetTheDomainWhenHasArg(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"domain", "example.com"}))
	assert.Equal(t, "example.com", e.domain)
}

func TestSetDomainShouldNotSetTheDomainWhenNoArg(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"domain"}))
	assert.Equal(t, "localhost", e.domain)
}

func TestSetDomainShouldNotSetVar(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"domain", "example.com"}))
	_, ok := e.vars["domain"]
	assert.False(t, ok)
}

func TestSetPortShouldSetThePortWhenHasIntArg(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"port", "8000"}))
	assert.Equal(t, 8000, e.port)
}

func TestSetPortShouldNotSetThePortWhenNotIntArg(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"port", "aaa"}))
	assert.Equal(t, 3000, e.port)
}

func TestSetPortShouldNotSetThePortWhenNoArg(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"port"}))
	assert.Equal(t, 3000, e.port)
}
func TestSetPortShouldNotSetVar(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"domain", "example.com"}))
	_, ok := e.vars["port"]
	assert.False(t, ok)
}
