package sbutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CheckRequiredFileOrDir(t *testing.T) {
	out := CheckRequiredFileOrDir("./go.mod")
	assert.True(t, out)

	out = CheckRequiredFileOrDir("./wrong.file")
	assert.False(t, out)

	out = CheckRequiredFileOrDir("./cmd")
	assert.True(t, out)

	out = CheckRequiredFileOrDir("./nonexist")
	assert.False(t, out)
}
