package sbutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Difference(t *testing.T) {
	diff := DifferenceStr([]string{"a", "b"}, []string{"a"})
	assert.Equal(t, []string{"b"}, diff)

	diff = DifferenceStr([]string{"a", "b"}, []string{"a", "b"})
	assert.Nil(t, diff)

	diff = DifferenceStr([]string{"a", "b", "c"}, []string{"a"})
	assert.Equal(t, []string{"b", "c"}, diff)
}

func Test_Contains(t *testing.T) {
	out := Contains([]string{"a", "b", "c"}, "c")
	assert.True(t, out)

	out = Contains([]string{"a", "b", "c"}, "d")
	assert.False(t, out)
}
