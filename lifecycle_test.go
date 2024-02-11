package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nhat.io/cpy3"
)

func TestInitialization(t *testing.T) {
	cpy3.Py_Initialize()

	assert.True(t, cpy3.Py_IsInitialized())

	cpy3.Py_Finalize()

	assert.False(t, cpy3.Py_IsInitialized())
}

func TestInitializationEx(t *testing.T) {
	cpy3.Py_Initialize()

	assert.True(t, cpy3.Py_IsInitialized())
	assert.Zero(t, cpy3.Py_FinalizeEx())
	assert.False(t, cpy3.Py_IsInitialized())
}

func TestPrefix(t *testing.T) {
	prefix, err := cpy3.Py_GetPrefix()

	require.NoError(t, err)
	assert.IsType(t, "", prefix)
}

func TestExecPrefix(t *testing.T) {
	execPrefix, err := cpy3.Py_GetExecPrefix()

	require.NoError(t, err)
	assert.IsType(t, "", execPrefix)
}

func TestProgramFullPath(t *testing.T) {
	programFullPath, err := cpy3.Py_GetProgramFullPath()

	require.NoError(t, err)
	assert.IsType(t, "", programFullPath)
}

func TestVersion(t *testing.T) {
	version := cpy3.Py_GetVersion()

	assert.IsType(t, "", version)
}

func TestPlatform(t *testing.T) {
	platform := cpy3.Py_GetPlatform()

	assert.IsType(t, "", platform)
}

func TestCopyright(t *testing.T) {
	copyright := cpy3.Py_GetCopyright()

	assert.IsType(t, "", copyright)
}

func TestCompiler(t *testing.T) {
	compiler := cpy3.Py_GetCompiler()

	assert.IsType(t, "", compiler)
}

func TestBuildInfo(t *testing.T) {
	buildInfo := cpy3.Py_GetBuildInfo()

	assert.IsType(t, "", buildInfo)
}
