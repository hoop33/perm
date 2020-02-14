package shell

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReplShouldReturnNonNilReplWhenConfigDirExists(t *testing.T) {
	r, err := NewRepl("/")
	assert.Nil(t, err)
	assert.NotNil(t, r)
}

func TestCreateLinerShouldReturnNonNilLiner(t *testing.T) {
	r, err := NewRepl("/")
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.NotNil(t, r.createLiner())
}

func TestNewReplShouldReturnErrorWhenConfigDirDoesNotExist(t *testing.T) {
	_, err := NewRepl("   foo   bar   baz   ")
	assert.NotNil(t, err)
}

func TestNewReplShouldReturnErrorWhenConfigDirIsFile(t *testing.T) {
	tmp, err := ioutil.TempFile("", "perm")
	assert.Nil(t, err)

	defer os.Remove(tmp.Name())

	_, err = NewRepl(tmp.Name())
	assert.NotNil(t, err)
}
