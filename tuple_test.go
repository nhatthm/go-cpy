package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestTupleCheck(t *testing.T) {
	cpy3.Py_Initialize()

	tuple := cpy3.PyTuple_New(0)
	defer tuple.DecRef()

	assert.True(t, cpy3.PyTuple_Check(tuple))
	assert.True(t, cpy3.PyTuple_CheckExact(tuple))
}

func TestTupleNew(t *testing.T) {
	cpy3.Py_Initialize()

	tuple := cpy3.PyTuple_New(0)
	defer tuple.DecRef()

	assert.NotNil(t, tuple)
}

func TestTupleSize(t *testing.T) {
	cpy3.Py_Initialize()

	size := 45

	tuple := cpy3.PyTuple_New(size)
	defer tuple.DecRef()

	assert.Equal(t, size, cpy3.PyTuple_Size(tuple))
}

func TestTupleGetSetItem(t *testing.T) {
	cpy3.Py_Initialize()

	s := cpy3.PyUnicode_FromString("test")

	i := cpy3.PyLong_FromGoInt(34)

	tuple := cpy3.PyTuple_New(2)
	defer tuple.DecRef()

	assert.Zero(t, cpy3.PyTuple_SetItem(tuple, 0, s))
	assert.Zero(t, cpy3.PyTuple_SetItem(tuple, 1, i))

	assert.Equal(t, i, cpy3.PyTuple_GetItem(tuple, 1))
}

func TestTupleGetSlice(t *testing.T) {
	cpy3.Py_Initialize()

	s := cpy3.PyUnicode_FromString("test")

	i := cpy3.PyLong_FromGoInt(34)

	tuple := cpy3.PyTuple_New(2)
	defer tuple.DecRef()

	assert.Zero(t, cpy3.PyTuple_SetItem(tuple, 0, s))
	assert.Zero(t, cpy3.PyTuple_SetItem(tuple, 1, i))

	slice := cpy3.PyTuple_GetSlice(tuple, 0, 1)
	defer slice.DecRef()

	assert.True(t, cpy3.PyTuple_Check(slice))
	assert.Equal(t, 1, cpy3.PyTuple_Size(slice))
	assert.Equal(t, s, cpy3.PyTuple_GetItem(slice, 0))
}
