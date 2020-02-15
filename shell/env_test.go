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
