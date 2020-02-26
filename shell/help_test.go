package shell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelpNameShouldReturnHelp(t *testing.T) {
	assert.Equal(t, "help", help(0).name())
}

func TestHelpsDescriptionShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", help(0).description())
}

func TestHelpUsageShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", help(0).usage())
}

func TestHelpRunShouldReturnNil(t *testing.T) {
	assert.Nil(t, help(0).run(nil, nil))
}

func TestHelpShouldRegisterItself(t *testing.T) {
	_, ok := allCommands[help(0).name()]
	assert.True(t, ok)
}
