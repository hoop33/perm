package shell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNameShouldReturnGet(t *testing.T) {
	assert.Equal(t, "get", get(0).name())
}

func TestGetDescriptionShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", get(0).description())
}

func TestGetUsageShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", get(0).usage())
}

func TestGetShouldRegisterItself(t *testing.T) {
	_, ok := allCommands[get(0).name()]
	assert.True(t, ok)
}
