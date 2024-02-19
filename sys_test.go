package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestSysGetSetObject(t *testing.T) {
	cpy.Py_Initialize()

	platform := cpy.PySys_GetObject("platform")

	assert.NotNil(t, platform)
	assert.True(t, cpy.PyUnicode_Check(platform))

	platform.IncRef()

	newPlatform := cpy.PyUnicode_FromString("test")
	defer newPlatform.DecRef()

	assert.Zero(t, cpy.PySys_SetObject("platform", newPlatform))
	assert.Equal(t, newPlatform, cpy.PySys_GetObject("platform"))
	assert.Zero(t, cpy.PySys_SetObject("platform", platform))
}
