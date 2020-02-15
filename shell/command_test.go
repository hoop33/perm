package shell

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestCommandsNameShouldReturnCommands(t *testing.T) {
  assert.Equal(t, "commands", commands(0).name())
}

func TestCommandssDescriptionShouldNotBeEmpty(t *testing.T) {
  assert.NotEqual(t, "", commands(0).description())
}

func TestCommandsRunShouldReturnNil(t *testing.T) {
  assert.Nil(t, commands(0).run(nil))
}

func TestCommandsShouldRegisterItself(t *testing.T) {
  _, ok := allCommands[commands(0).name()]
  assert.True(t, ok)
}

func TestSortedNamesShouldHaveTheSameLengthAsAllCommands(t *testing.T) {
  assert.Equal(t, len(sortedCommandNames()), len(allCommands))
}

func TestMaxLenShouldBeGreaterThanZero(t *testing.T) {
  assert.True(t, maxLen() > 0)
}
