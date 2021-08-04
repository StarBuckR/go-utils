package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_IsRoot(t *testing.T) {
	out, err := IsRoot()
	require.NoError(t, err)
	assert.Equal(t, false, out)
}
