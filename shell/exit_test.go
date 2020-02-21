package shell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExitNameShouldReturnExit(t *testing.T) {
	assert.Equal(t, "exit", exit(0).name())
}

func TestExitsDescriptionShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", exit(0).description())
}

func TestExitUsageShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", exit(0).usage())
}

func TestExitRunShouldReturnErrExit(t *testing.T) {
	assert.Equal(t, errExit, exit(0).run(nil, nil))
}

func TestExitShouldRegisterItself(t *testing.T) {
	_, ok := allCommands[exit(0).name()]
	assert.True(t, ok)
}
