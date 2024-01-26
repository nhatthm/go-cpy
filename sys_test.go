package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSysGetSetObject(t *testing.T) {
	Py_Initialize()

	platform := PySys_GetObject("platform")
	assert.NotNil(t, platform)
	assert.True(t, PyUnicode_Check(platform))
	platform.IncRef()

	newPlatform := PyUnicode_FromString("test")
	defer newPlatform.DecRef()

	assert.Zero(t, PySys_SetObject("platform", newPlatform))

	assert.Equal(t, newPlatform, PySys_GetObject("platform"))

	assert.Zero(t, PySys_SetObject("platform", platform))
}
