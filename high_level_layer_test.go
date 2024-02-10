package cpy3_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nhat.io/cpy3"
)

func TestRunFile(t *testing.T) {
	cpy3.Py_Initialize()

	pyErr, err := cpy3.PyRun_AnyFile("resources/fixtures/test.py")
	assert.Zero(t, pyErr)
	require.NoError(t, err)

	stdout := cpy3.PySys_GetObject("stdout")

	result := stdout.CallMethodArgs("getvalue")
	defer result.DecRef()

	assert.Equal(t, "hello world\n", cpy3.PyUnicode_AsUTF8(result))
}

func TestRunString(t *testing.T) {
	cpy3.Py_Initialize()

	pythonCode, err := os.ReadFile("resources/fixtures/test.py")
	require.NoError(t, err)

	assert.Zero(t, cpy3.PyRun_SimpleString(string(pythonCode)))

	stdout := cpy3.PySys_GetObject("stdout")

	result := stdout.CallMethodArgs("getvalue")
	defer result.DecRef()

	assert.Equal(t, "hello world\n", cpy3.PyUnicode_AsUTF8(result))
}

func TestPyMain(t *testing.T) {
	cpy3.Py_Initialize()

	pyErr, err := cpy3.Py_Main([]string{"resources/fixtures/test.py"})

	assert.Zero(t, pyErr)
	require.NoError(t, err)
}
