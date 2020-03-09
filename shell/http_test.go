package shell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetURLShouldUseSlashWhenNoArgs(t *testing.T) {
	e := newEnv()
	u, err := getURL(e, nil)
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost:3000/", u.String())
}

func TestGetURLShouldUseURLFromArgsWhenPresent(t *testing.T) {
	e := newEnv()
	u, err := getURL(e, []string{"/foo"})
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost:3000/foo", u.String())
}
