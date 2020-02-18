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

func TestUnsetShouldUnsetNothingWhenNoArgs(t *testing.T) {
	e := newEnv()
	e.setVar("foo", "bar")
	e.setHeader("baz", "bat")
	assert.Nil(t, unset(0).run(e, []string{""}))
	assert.Equal(t, 1, len(e.vars))
	assert.Equal(t, 1, len(e.headers))
}

func TestUnsetShouldUnsetNothingWhenNotVarOrHeader(t *testing.T) {
	e := newEnv()
	e.setVar("foo", "bar")
	e.setHeader("baz", "bat")
	assert.Nil(t, unset(0).run(e, []string{"foo"}))
	assert.Equal(t, 1, len(e.vars))
	assert.Equal(t, 1, len(e.headers))
}

func TestUnsetShouldUnsetAllVarsWhenVar(t *testing.T) {
	e := newEnv()
	e.setVar("foo", "bar")
	e.setVar("fooa", "bara")
	e.setHeader("baz", "bat")
	assert.Nil(t, unset(0).run(e, []string{"var"}))
	assert.Equal(t, 0, len(e.vars))
	assert.Equal(t, 1, len(e.headers))
}

func TestUnsetShouldUnsetAllHeadersWhenHeader(t *testing.T) {
	e := newEnv()
	e.setVar("foo", "bar")
	e.setHeader("baz", "bat")
	e.setHeader("baza", "bata")
	assert.Nil(t, unset(0).run(e, []string{"header"}))
	assert.Equal(t, 1, len(e.vars))
	assert.Equal(t, 0, len(e.headers))
}

func TestUnsetShouldUnsetVarWhenVarAndPresent(t *testing.T) {
	e := newEnv()
	e.setVar("foo", "bar")
	e.setVar("fooa", "bara")
	e.setHeader("baz", "bat")
	assert.Nil(t, unset(0).run(e, []string{"var", "foo"}))
	assert.Equal(t, 1, len(e.vars))
	assert.Equal(t, 1, len(e.headers))
}

func TestUnsetShouldUnsetHeaderWhenHeaderAndPresent(t *testing.T) {
	e := newEnv()
	e.setVar("foo", "bar")
	e.setHeader("baz", "bat")
	e.setHeader("baza", "bata")
	assert.Nil(t, unset(0).run(e, []string{"header", "baz"}))
	assert.Equal(t, 1, len(e.vars))
	assert.Equal(t, 1, len(e.headers))
}

func TestUnsetShouldNotUnsetVarWhenVarAndAbsent(t *testing.T) {
	e := newEnv()
	e.setVar("foo", "bar")
	e.setVar("fooa", "bara")
	e.setHeader("baz", "bat")
	assert.Nil(t, unset(0).run(e, []string{"var", "foob"}))
	assert.Equal(t, 2, len(e.vars))
	assert.Equal(t, 1, len(e.headers))
}

func TestUnsetShouldNotUnsetHeaderWhenHeaderAndAbsent(t *testing.T) {
	e := newEnv()
	e.setVar("foo", "bar")
	e.setHeader("baz", "bat")
	e.setHeader("baza", "bata")
	assert.Nil(t, unset(0).run(e, []string{"header", "bazb"}))
	assert.Equal(t, 1, len(e.vars))
	assert.Equal(t, 2, len(e.headers))
}
