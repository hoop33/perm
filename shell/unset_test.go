package shell

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	assert.Nil(t, unset(0).run(nil, nil))
}

func TestUnsetShouldRegisterItself(t *testing.T) {
	_, ok := allCommands[unset(0).name()]
	assert.True(t, ok)
}

func TestUnsetShouldRemoveVarFromEnvWhenPresent(t *testing.T) {
	e := newEnv()
	e.vars["foo"] = "bar"
	_, ok := e.vars["foo"]
	assert.True(t, ok)
	assert.Nil(t, unset(0).run(e, []string{"foo"}))
	_, ok = e.vars["foo"]
	assert.False(t, ok)
}

func TestUnsetShouldDoNothingWhenVarNotPresent(t *testing.T) {
	e := newEnv()
	assert.Nil(t, unset(0).run(e, []string{"foo"}))
	_, ok := e.vars["foo"]
	assert.False(t, ok)
}

func TestUnsetShouldDoNothingWhenNoVarPassed(t *testing.T) {
	e := newEnv()
	assert.Nil(t, unset(0).run(e, []string{""}))
	assert.Equal(t, 0, len(e.vars))
}
