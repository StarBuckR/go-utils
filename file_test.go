package sbutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CheckRequiredFileOrDir(t *testing.T) {
	out := CheckRequiredFileOrDir("./go.mod")
	assert.Equal(t, true, out)

	out = CheckRequiredFileOrDir("./wrong.file")
	assert.Equal(t, false, out)

	out = CheckRequiredFileOrDir("./cmd")
	assert.Equal(t, true, out)

	out = CheckRequiredFileOrDir("./nonexist")
	assert.Equal(t, false, out)
}
