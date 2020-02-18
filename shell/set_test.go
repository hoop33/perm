package shell

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestSetShouldSetEnvVarToTrueWhenOneArg(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"foo"}))
	assert.Equal(t, "true", e.vars["foo"])
}

func TestSetShouldSetEnvVarWhenTwoArgs(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"foo", "bar"}))
	assert.Equal(t, "bar", e.vars["foo"])
}

func TestSetShouldNotSetEnvVarWhenThreeArgs(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"foo", "bar", "baz"}))
	_, ok := e.vars["foo"]
	assert.False(t, ok)
}

func TestSetShouldNotSetEnvVarWhenNoArgs(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{}))
	assert.Equal(t, 0, len(e.vars))
}
