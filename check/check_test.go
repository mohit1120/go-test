package check

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	var a = "abc"
	var b = "abc"
	var c = check(a, b)
	assert.Equal(t, c, true)

	a = "abc"
	b = "bc"
	c = check(a, b)
	assert.Equal(t, c, false)
}
