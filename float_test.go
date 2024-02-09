package cpy3_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestPyFloatCheck(t *testing.T) {
	cpy3.Py_Initialize()

	pyFloat := cpy3.PyFloat_FromDouble(345.)
	defer pyFloat.DecRef()

	assert.True(t, cpy3.PyFloat_Check(pyFloat))
	assert.True(t, cpy3.PyFloat_CheckExact(pyFloat))
}

func TestPyFloatFromAsDouble(t *testing.T) {
	cpy3.Py_Initialize()

	v := 2354.

	pyFloat := cpy3.PyFloat_FromDouble(v)
	defer pyFloat.DecRef()

	assert.NotNil(t, pyFloat)
	assert.Equal(t, v, cpy3.PyFloat_AsDouble(pyFloat))
}

func TestPyFloatFromAsString(t *testing.T) {
	cpy3.Py_Initialize()

	pyString := cpy3.PyUnicode_FromString("2354")
	defer pyString.DecRef()

	pyFloat := cpy3.PyFloat_FromString(pyString)
	defer pyFloat.DecRef()

	assert.NotNil(t, pyFloat)
	assert.Equal(t, 2354., cpy3.PyFloat_AsDouble(pyFloat))
}

func TestPyFloatMinMax(t *testing.T) {
	cpy3.Py_Initialize()

	assert.Equal(t, math.MaxFloat64, cpy3.PyFloat_GetMax())
	assert.Equal(t, 2.2250738585072014e-308, cpy3.PyFloat_GetMin())
}

func TestPyFloatInfo(t *testing.T) {
	cpy3.Py_Initialize()

	assert.NotNil(t, cpy3.PyFloat_GetInfo())
}
