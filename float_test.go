package cpy_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestPyFloatCheck(t *testing.T) {
	cpy.Py_Initialize()

	pyFloat := cpy.PyFloat_FromDouble(345.)
	defer pyFloat.DecRef()

	assert.True(t, cpy.PyFloat_Check(pyFloat))
	assert.True(t, cpy.PyFloat_CheckExact(pyFloat))
}

func TestPyFloatFromAsDouble(t *testing.T) {
	cpy.Py_Initialize()

	v := 2354.

	pyFloat := cpy.PyFloat_FromDouble(v)
	defer pyFloat.DecRef()

	assert.NotNil(t, pyFloat)
	assert.InDelta(t, v, cpy.PyFloat_AsDouble(pyFloat), 0.01)
}

func TestPyFloatFromAsString(t *testing.T) {
	cpy.Py_Initialize()

	pyString := cpy.PyUnicode_FromString("2354")
	defer pyString.DecRef()

	pyFloat := cpy.PyFloat_FromString(pyString)
	defer pyFloat.DecRef()

	assert.NotNil(t, pyFloat)
	assert.InDelta(t, 2354., cpy.PyFloat_AsDouble(pyFloat), 0.01)
}

func TestPyFloatMinMax(t *testing.T) {
	cpy.Py_Initialize()

	assert.InDelta(t, math.MaxFloat64, cpy.PyFloat_GetMax(), 0.01)
	assert.InDelta(t, 2.2250738585072014e-308, cpy.PyFloat_GetMin(), 0.01)
}

func TestPyFloatInfo(t *testing.T) {
	cpy.Py_Initialize()

	assert.NotNil(t, cpy.PyFloat_GetInfo())
}
