package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestList(t *testing.T) {
	cpy3.Py_Initialize()

	list := cpy3.PyList_New(0)
	defer list.DecRef()

	assert.True(t, cpy3.PyList_Check(list))
	assert.True(t, cpy3.PyList_CheckExact(list))

	s := cpy3.PyUnicode_FromString("hello")

	assert.NotNil(t, s)

	i := cpy3.PyLong_FromGoInt(123)

	assert.NotNil(t, i)

	f := cpy3.PyFloat_FromDouble(123.4)

	assert.NotNil(t, f)

	assert.Zero(t, cpy3.PyList_Append(list, i))
	assert.Zero(t, cpy3.PyList_Insert(list, 0, s))

	assert.Equal(t, 2, cpy3.PyList_Size(list))

	assert.Zero(t, cpy3.PyList_SetItem(list, 0, f))

	assert.Equal(t, f, cpy3.PyList_GetItem(list, 0))

	assert.Zero(t, cpy3.PyList_Sort(list))
	assert.Equal(t, i, cpy3.PyList_GetItem(list, 0))

	assert.Zero(t, cpy3.PyList_Reverse(list))
	assert.Equal(t, f, cpy3.PyList_GetItem(list, 0))

	s = cpy3.PyUnicode_FromString("world")

	assert.NotNil(t, s)

	list2 := cpy3.PyList_New(1)
	defer list2.DecRef()

	assert.Zero(t, cpy3.PyList_SetItem(list2, 0, s))
	assert.Zero(t, cpy3.PyList_SetSlice(list, 0, 1, list2))

	list3 := cpy3.PyList_GetSlice(list, 0, 1)
	defer list3.DecRef()

	assert.NotNil(t, list3)

	assert.Equal(t, 1, list2.RichCompareBool(list3, cpy3.Py_EQ))

	tuple := cpy3.PyList_AsTuple(list)
	defer tuple.DecRef()

	assert.NotNil(t, tuple)

	world := cpy3.PyTuple_GetItem(tuple, 0)

	assert.NotNil(t, world)
	assert.Equal(t, "world", cpy3.PyUnicode_AsUTF8(world))
}
