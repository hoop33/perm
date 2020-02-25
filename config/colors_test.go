package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorShouldReturnColorizedTextWhenColorEnabled(t *testing.T) {
	EnableColor(true)
	assert.Equal(t, "\x1b[31mhello\x1b[0m", Error("hello"))
}

func TestErrorShouldReturnNonColorizedTextWhenColorDisabled(t *testing.T) {
	EnableColor(false)
	assert.Equal(t, "hello", Error("hello"))
}

func TestWarningShouldReturnColorizedTextWhenColorEnabled(t *testing.T) {
	EnableColor(true)
	assert.Equal(t, "\x1b[33mhello\x1b[0m", Warning("hello"))
}

func TestWarningShouldReturnNonColorizedTextWhenColorDisabled(t *testing.T) {
	EnableColor(false)
	assert.Equal(t, "hello", Warning("hello"))
}

func TestInfoShouldReturnColorizedTextWhenColorEnabled(t *testing.T) {
	EnableColor(true)
	assert.Equal(t, "\x1b[32mhello\x1b[0m", Info("hello"))
}

func TestInfoShouldReturnNonColorizedTextWhenColorDisabled(t *testing.T) {
	EnableColor(false)
	assert.Equal(t, "hello", Info("hello"))
}

func TestHeaderShouldReturnColorizedTextWhenColorEnabled(t *testing.T) {
	EnableColor(true)
	assert.Equal(t, "\x1b[34mhello\x1b[0m", Header("hello"))
}

func TestHeaderShouldReturnNonColorizedTextWhenColorDisabled(t *testing.T) {
	EnableColor(false)
	assert.Equal(t, "hello", Header("hello"))
}

func TestDefaultShouldReturnNonColorizedTextWhenColorEnabled(t *testing.T) {
	EnableColor(true)
	assert.Equal(t, "hello", Default("hello"))
}

func TestDefaultShouldReturnNonColorizedTextWhenColorDisabled(t *testing.T) {
	EnableColor(false)
	assert.Equal(t, "hello", Default("hello"))
}

func TestTextShouldReturnColorizedTextWhenColorEnabled(t *testing.T) {
	EnableColor(true)
	assert.Equal(t, "\x1b[35mhello\x1b[0m", Text("hello"))
}

func TestTextShouldReturnNonColorizedTextWhenColorDisabled(t *testing.T) {
	EnableColor(false)
	assert.Equal(t, "hello", Text("hello"))
}
