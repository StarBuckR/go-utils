package sbutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CheckDNSAccess(t *testing.T) {
	out, err := CheckDNSAccess("github.com")
	require.NoError(t, err)
	assert.True(t, out)

	out, err = CheckDNSAccess("github.nonexistent")
	require.Error(t, err)
	assert.False(t, out)
}

func Test_CheckPortAccess(t *testing.T) {
	out, err := CheckPortAccess("portquiz.net", 80)
	require.NoError(t, err)
	assert.True(t, out)

	out, err = CheckPortAccess("portquiz.net", 0)
	require.Error(t, err)
	assert.False(t, out)
}
