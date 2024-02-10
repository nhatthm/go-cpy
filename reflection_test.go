package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestReflectionBuiltins(t *testing.T) {
	cpy3.Py_Initialize()

	builtins := cpy3.PyEval_GetBuiltins()

	assert.NotNil(t, builtins)

	len := cpy3.PyDict_GetItemString(builtins, "len")

	assert.True(t, cpy3.PyCallable_Check(len))
}

func TestReflectionLocals(t *testing.T) {
	cpy3.Py_Initialize()

	locals := cpy3.PyEval_GetLocals()

	assert.Nil(t, locals)
}

func TestReflectionGlobals(t *testing.T) {
	cpy3.Py_Initialize()

	globals := cpy3.PyEval_GetGlobals()

	assert.Nil(t, globals)
}

func TestReflectionFuncName(t *testing.T) {
	cpy3.Py_Initialize()

	builtins := cpy3.PyEval_GetBuiltins()

	assert.NotNil(t, builtins)

	len := cpy3.PyDict_GetItemString(builtins, "len")

	assert.True(t, cpy3.PyCallable_Check(len))
	assert.Equal(t, "len", cpy3.PyEval_GetFuncName(len))
}
func TestReflectionFuncDesc(t *testing.T) {
	cpy3.Py_Initialize()

	builtins := cpy3.PyEval_GetBuiltins()

	assert.NotNil(t, builtins)

	len := cpy3.PyDict_GetItemString(builtins, "len")

	assert.True(t, cpy3.PyCallable_Check(len))
	assert.Equal(t, "()", cpy3.PyEval_GetFuncDesc(len))
}
