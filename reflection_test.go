package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestReflectionBuiltins(t *testing.T) {
	cpy.Py_Initialize()

	builtins := cpy.PyEval_GetBuiltins()

	assert.NotNil(t, builtins)

	builtinsLength := cpy.PyDict_GetItemString(builtins, "len")

	assert.True(t, cpy.PyCallable_Check(builtinsLength))
}

func TestReflectionLocals(t *testing.T) {
	cpy.Py_Initialize()

	locals := cpy.PyEval_GetLocals()

	assert.Nil(t, locals)
}

func TestReflectionGlobals(t *testing.T) {
	cpy.Py_Initialize()

	globals := cpy.PyEval_GetGlobals()

	assert.Nil(t, globals)
}

func TestReflectionFuncName(t *testing.T) {
	cpy.Py_Initialize()

	builtins := cpy.PyEval_GetBuiltins()

	assert.NotNil(t, builtins)

	builtinsLength := cpy.PyDict_GetItemString(builtins, "len")

	assert.True(t, cpy.PyCallable_Check(builtinsLength))
	assert.Equal(t, "len", cpy.PyEval_GetFuncName(builtinsLength))
}

func TestReflectionFuncDesc(t *testing.T) {
	cpy.Py_Initialize()

	builtins := cpy.PyEval_GetBuiltins()

	assert.NotNil(t, builtins)

	builtinsLength := cpy.PyDict_GetItemString(builtins, "len")

	assert.True(t, cpy.PyCallable_Check(builtinsLength))
	assert.Equal(t, "()", cpy.PyEval_GetFuncDesc(builtinsLength))
}
