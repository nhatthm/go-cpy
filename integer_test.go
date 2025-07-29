package cpy_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestPyLongCheck(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	pyLong := cpy.PyLong_FromGoInt(345)
	defer pyLong.DecRef()

	assert.True(t, cpy.PyLong_Check(pyLong))
	assert.True(t, cpy.PyLong_CheckExact(pyLong))
}

func TestPyLongFromAsLong(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	v := 2354

	pyLong := cpy.PyLong_FromLong(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy.PyLong_AsLong(pyLong))
}

func TestPyLongFromAsUnsignedLong(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	v := uint(2354)

	pyLong := cpy.PyLong_FromUnsignedLong(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy.PyLong_AsUnsignedLong(pyLong))
}

func TestPyLongFromAsLongLong(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	v := int64(2354)

	pyLong := cpy.PyLong_FromLongLong(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy.PyLong_AsLongLong(pyLong))
}

func TestPyLongFromAsUnsignedLongLong(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	v := uint64(2354)

	pyLong := cpy.PyLong_FromUnsignedLongLong(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy.PyLong_AsUnsignedLongLong(pyLong))
}

func TestPyLongFromAsDouble(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	v := float64(2354.0)

	pyLong := cpy.PyLong_FromDouble(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.InDelta(t, v, cpy.PyLong_AsDouble(pyLong), 0.01)
}

func TestPyLongFromAsGoFloat64(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	v := float64(2354.0)

	pyLong := cpy.PyLong_FromGoFloat64(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.InDelta(t, v, cpy.PyLong_AsDouble(pyLong), 0.01)
}

func TestPyLongFromAsString(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	v := 2354
	s := strconv.Itoa(v)

	pyLong := cpy.PyLong_FromString(s, 10)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy.PyLong_AsLong(pyLong))
}

func TestPyLongFromAsUnicodeObject(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	v := 2354
	s := strconv.Itoa(v)

	pyUnicode := cpy.PyUnicode_FromString(s)
	defer pyUnicode.DecRef()

	assert.NotNil(t, pyUnicode)

	pyLong := cpy.PyLong_FromUnicodeObject(pyUnicode, 10)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy.PyLong_AsLong(pyLong))
}

func TestPyLongFromAsGoInt(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	v := 2354

	pyLong := cpy.PyLong_FromGoInt(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy.PyLong_AsLong(pyLong))
}

func TestPyLongFromAsGoUint(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	v := uint(2354)

	pyLong := cpy.PyLong_FromGoUint(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy.PyLong_AsUnsignedLong(pyLong))
}

func TestPyLongFromAsGoInt64(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	v := int64(2354)

	pyLong := cpy.PyLong_FromGoInt64(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy.PyLong_AsLongLong(pyLong))
}

func TestPyLongFromAsGoUint64(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	v := uint64(2354)

	pyLong := cpy.PyLong_FromGoUint64(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy.PyLong_AsUnsignedLongLong(pyLong))
}
