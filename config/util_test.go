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

func TestHistoryFileShouldEndWithHistory(t *testing.T) {
	assert.True(t, strings.HasSuffix(HistoryFile(), ".history"))
}

func TestAddPrefixIfAbsentShouldReturnBlankWhenBothAreBlank(t *testing.T) {
	assert.Equal(t, "", AddPrefixIfAbsent("", ""))
}

func TestAddPrefixIfAbsentShouldAddPrefixWhenAbsent(t *testing.T) {
	assert.Equal(t, "foobar", AddPrefixIfAbsent("bar", "foo"))
}

func TestAddPrefixIfAbsentShouldNotAddPrefixWhenPresnt(t *testing.T) {
	assert.Equal(t, "foobar", AddPrefixIfAbsent("foobar", "foo"))
}

func TestFirstNonBlankShouldReturnBlankWhenEmpty(t *testing.T) {
	assert.Equal(t, "", FirstNonBlank())
}

func TestFirstNonBlankShouldReturnBlankWhenAllBlank(t *testing.T) {
	assert.Equal(t, "", FirstNonBlank("", "", "", "", ""))
}

func TestFirstNonBlankShouldReturnFirstElementWhenNotBlank(t *testing.T) {
	assert.Equal(t, "one", FirstNonBlank("one", "two", "three"))
}

func TestFirstNonBlankShouldReturnFirstNonBlankElement(t *testing.T) {
	assert.Equal(t, "three", FirstNonBlank("", "", "three"))
}
