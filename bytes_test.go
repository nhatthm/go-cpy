package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestBytesCheck(t *testing.T) {
	cpy3.Py_Initialize()

	s1 := "aaaaaaaa"

	bytes1 := cpy3.PyBytes_FromString(s1)
	defer bytes1.DecRef()

	assert.True(t, cpy3.PyBytes_Check(bytes1))
	assert.True(t, cpy3.PyBytes_CheckExact(bytes1))
}

func TestBytesFromAsString(t *testing.T) {
	cpy3.Py_Initialize()

	s1 := "aaaaaaaa"

	bytes1 := cpy3.PyBytes_FromString(s1)
	defer bytes1.DecRef()

	assert.Equal(t, s1, cpy3.PyBytes_AsString(bytes1))
}

func TestBytesSize(t *testing.T) {
	cpy3.Py_Initialize()

	s1 := "aaaaaaaa"

	bytes1 := cpy3.PyBytes_FromString(s1)
	defer bytes1.DecRef()

	assert.Equal(t, 8, cpy3.PyBytes_Size(bytes1))
}

func TestBytesConcat(t *testing.T) {
	cpy3.Py_Initialize()

	s1 := "aaaaaaaa"
	s2 := "bbbbbbbb"

	bytes1 := cpy3.PyBytes_FromString(s1)

	array := cpy3.PyByteArray_FromStringAndSize(s2)
	defer array.DecRef()

	bytes2 := cpy3.PyBytes_FromObject(array)
	defer bytes2.DecRef()

	assert.NotNil(t, bytes2)

	bytes1 = cpy3.PyBytes_Concat(bytes1, bytes2)
	defer bytes1.DecRef()

	assert.NotNil(t, bytes1)
	assert.Equal(t, s1+s2, cpy3.PyBytes_AsString(bytes1))
}

func TestBytesConcatAndDel(t *testing.T) {
	cpy3.Py_Initialize()

	s1 := "aaaaaaaa"
	s2 := "bbbbbbbb"

	bytes1 := cpy3.PyBytes_FromString(s1)

	bytes2 := cpy3.PyBytes_FromString(s2)

	assert.NotNil(t, bytes2)

	bytes1 = cpy3.PyBytes_ConcatAndDel(bytes1, bytes2)
	defer bytes1.DecRef()

	assert.NotNil(t, bytes1)
	assert.Equal(t, s1+s2, cpy3.PyBytes_AsString(bytes1))
}

func TestByteSlices(t *testing.T) {
	cpy3.Py_Initialize()

	s1 := []byte("aaaaaaaa")
	s2 := []byte("bbbbbbbb")

	bytes1 := cpy3.PyBytes_FromByteSlice(s1)
	defer bytes1.DecRef()

	bytes2 := cpy3.PyBytes_FromByteSlice(s2)
	defer bytes2.DecRef()

	assert.Equal(t, s1, cpy3.PyBytes_AsByteSlice(bytes1))
	assert.Equal(t, s2, cpy3.PyBytes_AsByteSlice(bytes2))
}
