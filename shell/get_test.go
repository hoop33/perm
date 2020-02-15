package shell

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestGetRunShouldReturnNil(t *testing.T) {
	assert.Nil(t, get(0).run(nil))
}

func TestGetShouldRegisterItself(t *testing.T) {
	_, ok := allCommands[get(0).name()]
	assert.True(t, ok)
}
