package shell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReplShouldReturnNonNilRepl(t *testing.T) {
	assert.NotNil(t, NewRepl())
}

func TestCreateLinerShouldNotReturnNil(t *testing.T) {
	assert.NotNil(t, NewRepl().createLiner())
}
