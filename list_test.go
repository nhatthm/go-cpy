package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestList(t *testing.T) {
	cpy.Py_Initialize()

	list := cpy.PyList_New(0)
	defer list.DecRef()

	assert.True(t, cpy.PyList_Check(list))
	assert.True(t, cpy.PyList_CheckExact(list))

	s := cpy.PyUnicode_FromString("hello")

	assert.NotNil(t, s)

	i := cpy.PyLong_FromGoInt(123)

	assert.NotNil(t, i)

	f := cpy.PyFloat_FromDouble(123.4)

	assert.NotNil(t, f)

	assert.Zero(t, cpy.PyList_Append(list, i))
	assert.Zero(t, cpy.PyList_Insert(list, 0, s))

	assert.Equal(t, 2, cpy.PyList_Size(list))

	assert.Zero(t, cpy.PyList_SetItem(list, 0, f))

	assert.Equal(t, f, cpy.PyList_GetItem(list, 0))

	assert.Zero(t, cpy.PyList_Sort(list))
	assert.Equal(t, i, cpy.PyList_GetItem(list, 0))

	assert.Zero(t, cpy.PyList_Reverse(list))
	assert.Equal(t, f, cpy.PyList_GetItem(list, 0))

	s = cpy.PyUnicode_FromString("world")

	assert.NotNil(t, s)

	list2 := cpy.PyList_New(1)
	defer list2.DecRef()

	assert.Zero(t, cpy.PyList_SetItem(list2, 0, s))
	assert.Zero(t, cpy.PyList_SetSlice(list, 0, 1, list2))

	list3 := cpy.PyList_GetSlice(list, 0, 1)
	defer list3.DecRef()

	assert.NotNil(t, list3)

	assert.Equal(t, 1, list2.RichCompareBool(list3, cpy.Py_EQ))

	tuple := cpy.PyList_AsTuple(list)
	defer tuple.DecRef()

	assert.NotNil(t, tuple)

	world := cpy.PyTuple_GetItem(tuple, 0)

	assert.NotNil(t, world)
	assert.Equal(t, "world", cpy.PyUnicode_AsUTF8(world))
}
