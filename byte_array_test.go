package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestByteArrayCheck(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	s1 := "aaaaaaaa"

	array1 := cpy.PyByteArray_FromStringAndSize(s1)
	defer array1.DecRef()

	assert.True(t, cpy.PyByteArray_Check(array1))
	assert.True(t, cpy.PyByteArray_CheckExact(array1))
}

func TestByteArrayFromAsString(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	s1 := "aaaaaaaa"

	array1 := cpy.PyByteArray_FromStringAndSize(s1)
	defer array1.DecRef()

	assert.Equal(t, s1, cpy.PyByteArray_AsString(array1))
}

func TestByteArrayConcat(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	s1 := "aaaaaaaa"
	s2 := "bbbbbbbb"

	array1 := cpy.PyByteArray_FromStringAndSize(s1)
	defer array1.DecRef()

	bytes := cpy.PyBytes_FromString(s2)
	defer bytes.DecRef()

	assert.NotNil(t, bytes)

	array2 := cpy.PyByteArray_FromObject(bytes)
	defer array2.DecRef()

	assert.NotNil(t, array2)

	newArray := cpy.PyByteArray_Concat(array1, array2)
	defer newArray.DecRef()

	assert.Equal(t, s1+s2, cpy.PyByteArray_AsString(newArray))
}

func TestByteArrayResize(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	s1 := "aaaaaaaa"

	array1 := cpy.PyByteArray_FromStringAndSize(s1)
	defer array1.DecRef()

	length := 20

	cpy.PyByteArray_Resize(array1, 20)

	assert.Equal(t, length, cpy.PyByteArray_Size(array1))
}
