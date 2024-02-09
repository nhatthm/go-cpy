package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestSysGetSetObject(t *testing.T) {
	cpy3.Py_Initialize()

	platform := cpy3.PySys_GetObject("platform")

	assert.NotNil(t, platform)
	assert.True(t, cpy3.PyUnicode_Check(platform))

	platform.IncRef()

	newPlatform := cpy3.PyUnicode_FromString("test")
	defer newPlatform.DecRef()

	assert.Zero(t, cpy3.PySys_SetObject("platform", newPlatform))
	assert.Equal(t, newPlatform, cpy3.PySys_GetObject("platform"))
	assert.Zero(t, cpy3.PySys_SetObject("platform", platform))
}
