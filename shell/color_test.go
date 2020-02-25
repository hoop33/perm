package shell

import (
	"testing"

	"github.com/hoop33/perm/config"
	"github.com/stretchr/testify/assert"
)

func TestColorNameShouldReturnColor(t *testing.T) {
	assert.Equal(t, "color", color(0).name())
}

func TestColorsDescriptionShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", color(0).description())
}

func TestColorUsageShouldNotBeEmpty(t *testing.T) {
	assert.NotEqual(t, "", color(0).usage())
}

func TestColorRunShouldReturnNil(t *testing.T) {
	assert.Nil(t, color(0).run(nil, nil))
}

func TestColorRunShouldTurnColorOnWhenNoArgs(t *testing.T) {
	color(0).run(nil, nil)
	assert.Equal(t, "\x1b[35mhello\x1b[0m", config.Text("hello"))
}

func TestColorRunShouldTurnColorOnWhenOn(t *testing.T) {
	color(0).run(nil, []string{"on"})
	assert.Equal(t, "\x1b[35mhello\x1b[0m", config.Text("hello"))
}

func TestColorRunShouldTurnColorOnWhenON(t *testing.T) {
	color(0).run(nil, []string{"ON"})
	assert.Equal(t, "\x1b[35mhello\x1b[0m", config.Text("hello"))
}

func TestColorRunShouldTurnColorOnWhenTrue(t *testing.T) {
	color(0).run(nil, []string{"true"})
	assert.Equal(t, "\x1b[35mhello\x1b[0m", config.Text("hello"))
}

func TestColorRunShouldTurnColorOffWhenOff(t *testing.T) {
	color(0).run(nil, []string{"off"})
	assert.Equal(t, "hello", config.Text("hello"))
}

func TestColorRunShouldTurnColorOffWhenFalse(t *testing.T) {
	color(0).run(nil, []string{"false"})
	assert.Equal(t, "hello", config.Text("hello"))
}

func TestColorShouldRegisterItself(t *testing.T) {
	_, ok := allCommands[color(0).name()]
	assert.True(t, ok)
}
