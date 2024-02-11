package cpy3_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestPyLongCheck(t *testing.T) {
	cpy3.Py_Initialize()

	pyLong := cpy3.PyLong_FromGoInt(345)
	defer pyLong.DecRef()

	assert.True(t, cpy3.PyLong_Check(pyLong))
	assert.True(t, cpy3.PyLong_CheckExact(pyLong))
}

func TestPyLongFromAsLong(t *testing.T) {
	cpy3.Py_Initialize()

	v := 2354

	pyLong := cpy3.PyLong_FromLong(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy3.PyLong_AsLong(pyLong))
}

func TestPyLongFromAsUnsignedLong(t *testing.T) {
	cpy3.Py_Initialize()

	v := uint(2354)

	pyLong := cpy3.PyLong_FromUnsignedLong(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy3.PyLong_AsUnsignedLong(pyLong))
}

func TestPyLongFromAsLongLong(t *testing.T) {
	cpy3.Py_Initialize()

	v := int64(2354)

	pyLong := cpy3.PyLong_FromLongLong(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy3.PyLong_AsLongLong(pyLong))
}

func TestPyLongFromAsUnsignedLongLong(t *testing.T) {
	cpy3.Py_Initialize()

	v := uint64(2354)

	pyLong := cpy3.PyLong_FromUnsignedLongLong(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy3.PyLong_AsUnsignedLongLong(pyLong))
}

func TestPyLongFromAsDouble(t *testing.T) {
	cpy3.Py_Initialize()

	v := float64(2354.0)

	pyLong := cpy3.PyLong_FromDouble(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.InDelta(t, v, cpy3.PyLong_AsDouble(pyLong), 0.01)
}

func TestPyLongFromAsGoFloat64(t *testing.T) {
	cpy3.Py_Initialize()

	v := float64(2354.0)

	pyLong := cpy3.PyLong_FromGoFloat64(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.InDelta(t, v, cpy3.PyLong_AsDouble(pyLong), 0.01)
}

func TestPyLongFromAsString(t *testing.T) {
	cpy3.Py_Initialize()

	v := 2354
	s := strconv.Itoa(v)

	pyLong := cpy3.PyLong_FromString(s, 10)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy3.PyLong_AsLong(pyLong))
}

func TestPyLongFromAsUnicodeObject(t *testing.T) {
	cpy3.Py_Initialize()

	v := 2354
	s := strconv.Itoa(v)

	pyUnicode := cpy3.PyUnicode_FromString(s)
	defer pyUnicode.DecRef()

	assert.NotNil(t, pyUnicode)

	pyLong := cpy3.PyLong_FromUnicodeObject(pyUnicode, 10)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy3.PyLong_AsLong(pyLong))
}

func TestPyLongFromAsGoInt(t *testing.T) {
	cpy3.Py_Initialize()

	v := 2354

	pyLong := cpy3.PyLong_FromGoInt(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy3.PyLong_AsLong(pyLong))
}

func TestPyLongFromAsGoUint(t *testing.T) {
	cpy3.Py_Initialize()

	v := uint(2354)

	pyLong := cpy3.PyLong_FromGoUint(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy3.PyLong_AsUnsignedLong(pyLong))
}

func TestPyLongFromAsGoInt64(t *testing.T) {
	cpy3.Py_Initialize()

	v := int64(2354)

	pyLong := cpy3.PyLong_FromGoInt64(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy3.PyLong_AsLongLong(pyLong))
}

func TestPyLongFromAsGoUint64(t *testing.T) {
	cpy3.Py_Initialize()

	v := uint64(2354)

	pyLong := cpy3.PyLong_FromGoUint64(v)
	defer pyLong.DecRef()

	assert.NotNil(t, pyLong)
	assert.Equal(t, v, cpy3.PyLong_AsUnsignedLongLong(pyLong))
}
