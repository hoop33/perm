package shell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionNameShouldReturnVersion(t *testing.T) {
	assert.Equal(t, "version", version(0).name())
}

func TestVersionsDescriptionShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", version(0).description())
}

func TestVersionUsageShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", version(0).usage())
}

func TestVersionRunShouldReturnNil(t *testing.T) {
	assert.Nil(t, version(0).run(nil, nil))
}

func TestVersionShouldRegisterItself(t *testing.T) {
	_, ok := allCommands[version(0).name()]
	assert.True(t, ok)
}
