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

	builtinsLength := cpy3.PyDict_GetItemString(builtins, "len")

	assert.True(t, cpy3.PyCallable_Check(builtinsLength))
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

	builtinsLength := cpy3.PyDict_GetItemString(builtins, "len")

	assert.True(t, cpy3.PyCallable_Check(builtinsLength))
	assert.Equal(t, "len", cpy3.PyEval_GetFuncName(builtinsLength))
}

func TestReflectionFuncDesc(t *testing.T) {
	cpy3.Py_Initialize()

	builtins := cpy3.PyEval_GetBuiltins()

	assert.NotNil(t, builtins)

	builtinsLength := cpy3.PyDict_GetItemString(builtins, "len")

	assert.True(t, cpy3.PyCallable_Check(builtinsLength))
	assert.Equal(t, "()", cpy3.PyEval_GetFuncDesc(builtinsLength))
}
