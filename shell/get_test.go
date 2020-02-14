package shell

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGetNameShouldReturnGet(t *testing.T) {
  assert.Equal(t, "get", get(0).name())
}

func TestGetRunShouldReturnNil(t *testing.T) {
  assert.Nil(t, get(0).run(nil))
}

func TestGetShouldRegisterItself(t *testing.T) {
  _, ok := commands[get(0).name()]
  assert.True(t, ok)
}
