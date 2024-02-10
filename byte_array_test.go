package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestByteArrayCheck(t *testing.T) {
	cpy3.Py_Initialize()

	s1 := "aaaaaaaa"

	array1 := cpy3.PyByteArray_FromStringAndSize(s1)
	defer array1.DecRef()

	assert.True(t, cpy3.PyByteArray_Check(array1))
	assert.True(t, cpy3.PyByteArray_CheckExact(array1))
}

func TestByteArrayFromAsString(t *testing.T) {
	cpy3.Py_Initialize()

	s1 := "aaaaaaaa"

	array1 := cpy3.PyByteArray_FromStringAndSize(s1)
	defer array1.DecRef()

	assert.Equal(t, s1, cpy3.PyByteArray_AsString(array1))
}

func TestByteArrayConcat(t *testing.T) {
	cpy3.Py_Initialize()

	s1 := "aaaaaaaa"
	s2 := "bbbbbbbb"

	array1 := cpy3.PyByteArray_FromStringAndSize(s1)
	defer array1.DecRef()

	bytes := cpy3.PyBytes_FromString(s2)
	defer bytes.DecRef()

	assert.NotNil(t, bytes)

	array2 := cpy3.PyByteArray_FromObject(bytes)
	defer array2.DecRef()

	assert.NotNil(t, array2)

	newArray := cpy3.PyByteArray_Concat(array1, array2)
	defer newArray.DecRef()

	assert.Equal(t, s1+s2, cpy3.PyByteArray_AsString(newArray))
}

func TestByteArrayResize(t *testing.T) {
	cpy3.Py_Initialize()

	s1 := "aaaaaaaa"

	array1 := cpy3.PyByteArray_FromStringAndSize(s1)
	defer array1.DecRef()

	length := 20
	cpy3.PyByteArray_Resize(array1, 20)

	assert.Equal(t, length, cpy3.PyByteArray_Size(array1))
}
