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

func TestSetShouldSetNothingWhenNoArgs(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{}))
	assert.Equal(t, 0, len(e.vars))
	assert.Equal(t, 0, len(e.headers))
}

func TestSetShouldSetNothingWhenOneArg(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"var"}))
	assert.Equal(t, 0, len(e.vars))
	assert.Equal(t, 0, len(e.headers))
}

func TestSetShouldSetVarWhenNotVarOrHeader(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"foo", "bar"}))
	assert.Equal(t, 1, len(e.vars))
	assert.Equal(t, 0, len(e.headers))
}

func TestSetShouldSetVarToTrueWhenVarKey(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"var", "foo"}))
	assert.Equal(t, "true", e.vars["foo"][0])
	assert.Equal(t, 0, len(e.headers))
}

func TestSetShouldSetVarWhenVarKeyValue(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"var", "foo", "bar"}))
	assert.Equal(t, "bar", e.vars["foo"][0])
	assert.Equal(t, 0, len(e.headers))
}

func TestSetShouldSetVarWhenVarKeyValueExtra(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"var", "foo", "bar", "baz"}))
	assert.Equal(t, "bar", e.vars["foo"][0])
	assert.Equal(t, "baz", e.vars["foo"][1])
	assert.Equal(t, 0, len(e.headers))
}

func TestSetShouldSetHeaderToTrueWhenVarKey(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"header", "foo"}))
	assert.Equal(t, "true", e.headers["foo"][0])
	assert.Equal(t, 0, len(e.vars))
}

func TestSetShouldSetHeaderWhenVarKeyValue(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"header", "foo", "bar"}))
	assert.Equal(t, "bar", e.headers["foo"][0])
	assert.Equal(t, 0, len(e.vars))
}

func TestSetShouldSetHeaderWhenVarKeyValueExtra(t *testing.T) {
	e := newEnv()
	assert.Nil(t, set(0).run(e, []string{"header", "foo", "bar", "baz"}))
	assert.Equal(t, "bar", e.headers["foo"][0])
	assert.Equal(t, "baz", e.headers["foo"][1])
	assert.Equal(t, 0, len(e.vars))
}
