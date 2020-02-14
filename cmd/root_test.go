package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCmdShouldHaveNonEmptyUse(t *testing.T) {
	assert.NotEqual(t, "", rootCmd.Use)
}

func TestRootCmdShouldHaveNonEmptyShort(t *testing.T) {
	assert.NotEqual(t, "", rootCmd.Short)
}

func TestRootCmdShouldHaveNonEmptyLong(t *testing.T) {
	assert.NotEqual(t, "", rootCmd.Long)
}

func TestRootCmdShouldHaveNonEmptyVersion(t *testing.T) {
	assert.NotEqual(t, "", rootCmd.Version)
}
