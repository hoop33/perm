package shell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionNameShouldReturnVersion(t *testing.T) {
	assert.Equal(t, "version", version(0).name())
}

func TestVersionRunShouldReturnNil(t *testing.T) {
	assert.Nil(t, version(0).run(nil))
}

func TestVersionShouldRegisterItself(t *testing.T) {
	_, ok := commands[version(0).name()]
	assert.True(t, ok)
}
