package shell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsetNameShouldReturnUnset(t *testing.T) {
	assert.Equal(t, "unset", unset(0).name())
}

func TestUnsetsDescriptionShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", unset(0).description())
}

func TestUnsetUsageShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", unset(0).usage())
}

func TestUnsetRunShouldReturnNil(t *testing.T) {
	e := newEnv()
	assert.Nil(t, unset(0).run(e, nil))
}

func TestUnsetShouldRegisterItself(t *testing.T) {
	_, ok := allCommands[unset(0).name()]
	assert.True(t, ok)
}

func TestUnsetShouldRemoveVarFromEnvWhenPresent(t *testing.T) {
	e := newEnv()
	e.vars["foo"] = "bar"
	e.vars["baz"] = "bat"
	_, ok := e.vars["foo"]
	assert.True(t, ok)
	_, ok = e.vars["baz"]
	assert.True(t, ok)
	assert.Nil(t, unset(0).run(e, []string{"foo"}))
	_, ok = e.vars["foo"]
	assert.False(t, ok)
	_, ok = e.vars["baz"]
	assert.True(t, ok)
}

func TestUnsetShouldDoNothingWhenVarNotPresent(t *testing.T) {
	e := newEnv()
	assert.Nil(t, unset(0).run(e, []string{"foo"}))
	_, ok := e.vars["foo"]
	assert.False(t, ok)
}

func TestUnsetShouldRemoveAllWhenNoVarPassed(t *testing.T) {
	e := newEnv()
	e.vars["foo"] = "bar"
	e.vars["baz"] = "bat"
	_, ok := e.vars["foo"]
	assert.True(t, ok)
	_, ok = e.vars["baz"]
	assert.True(t, ok)
	assert.Nil(t, unset(0).run(e, []string{}))
	_, ok = e.vars["foo"]
	assert.False(t, ok)
	_, ok = e.vars["baz"]
	assert.False(t, ok)
	assert.Equal(t, 0, len(e.vars))
}
