package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ExecuteCommand(t *testing.T) {
	out, err := ExecuteCommand("echo 'test'")
	require.NoError(t, err)
	assert.Equal(t, "test", out)
}

func Test_ExecuteCommandWithErrorBuffer(t *testing.T) {
	out, err := ExecuteCommandWithErrorBuffer("echo 'test'")
	require.NoError(t, err)
	assert.Equal(t, "test", out)

	_, err = ExecuteCommandWithErrorBuffer("ech 'test'")
	require.Error(t, err)
}
