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
