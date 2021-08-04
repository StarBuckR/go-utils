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

func Test_CheckIsAppAlreadyRunning(t *testing.T) {
	out, err := CheckIsAppAlreadyRunning("test")
	require.NoError(t, err)
	require.Equal(t, true, out)

	out, err = CheckIsAppAlreadyRunning("wrongtest")
	require.Error(t, err)
	require.Equal(t, false, out)
}
