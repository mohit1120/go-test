package check

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	var a = "abc"
	var b = "abc"
	var c = Check(a, b)
	assert.Equal(t, c, true)

	a = "abc"
	b = "bc"
	c = Check(a, b)
	assert.Equal(t, c, false)
}
