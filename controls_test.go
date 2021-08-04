package sbutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_IsRoot(t *testing.T) {
	out, err := IsRoot()
	require.NoError(t, err)
	assert.False(t, out)
}

func Test_CheckIsAppAlreadyRunning(t *testing.T) {
	out, err := CheckIsAppAlreadyRunning("test")
	require.NoError(t, err)
	require.True(t, out)

	out, err = CheckIsAppAlreadyRunning("wrongtest")
	require.Error(t, err)
	require.False(t, out)
}
