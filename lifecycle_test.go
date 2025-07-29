package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nhat.io/cpy/v3"
)

func TestInitialization(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	assert.True(t, cpy.Py_IsInitialized())

	cpy.Py_Finalize()

	assert.False(t, cpy.Py_IsInitialized())
}

func TestInitializationEx(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	assert.True(t, cpy.Py_IsInitialized())
	assert.Zero(t, cpy.Py_FinalizeEx())
	assert.False(t, cpy.Py_IsInitialized())
}

func TestPrefix(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	prefix, err := cpy.Py_GetPrefix()

	require.NoError(t, err)
	assert.IsType(t, "", prefix)
}

func TestExecPrefix(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	execPrefix, err := cpy.Py_GetExecPrefix()

	require.NoError(t, err)
	assert.IsType(t, "", execPrefix)
}

func TestProgramFullPath(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	programFullPath, err := cpy.Py_GetProgramFullPath()

	require.NoError(t, err)
	assert.IsType(t, "", programFullPath)
}

func TestVersion(t *testing.T) {
	version := cpy.Py_GetVersion()

	assert.IsType(t, "", version)
}

func TestPlatform(t *testing.T) {
	platform := cpy.Py_GetPlatform()

	assert.IsType(t, "", platform)
}

func TestCopyright(t *testing.T) {
	copyright := cpy.Py_GetCopyright()

	assert.IsType(t, "", copyright)
}

func TestCompiler(t *testing.T) {
	compiler := cpy.Py_GetCompiler()

	assert.IsType(t, "", compiler)
}

func TestBuildInfo(t *testing.T) {
	buildInfo := cpy.Py_GetBuildInfo()

	assert.IsType(t, "", buildInfo)
}
