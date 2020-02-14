package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirShouldReturnNonEmptyPath(t *testing.T) {
	assert.NotEqual(t, "", Dir())
}

func TestDirShouldReturnPathWithAppNameSuffix(t *testing.T) {
	assert.True(t, strings.HasSuffix(Dir(), AppName))
}
