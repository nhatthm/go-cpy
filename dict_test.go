package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestDict(t *testing.T) {
	cpy.Py_Initialize()

	dict := cpy.PyDict_New()
	defer dict.DecRef()

	assert.True(t, cpy.PyDict_Check(dict))
	assert.True(t, cpy.PyDict_CheckExact(dict))

	proxy := cpy.PyDictProxy_New(dict)
	defer proxy.DecRef()

	assert.NotNil(t, proxy)

	key1 := "key1"

	value1 := cpy.PyUnicode_FromString("value1")
	defer value1.DecRef()

	assert.NotNil(t, value1)

	key2 := cpy.PyUnicode_FromString("key2")
	defer key2.DecRef()

	assert.NotNil(t, key2)

	value2 := cpy.PyUnicode_FromString("value2")
	defer value2.DecRef()

	assert.NotNil(t, value2)

	key3 := cpy.PyUnicode_FromString("key3")
	defer key3.DecRef()

	assert.NotNil(t, key3)

	value3 := cpy.PyUnicode_FromString("value3")
	defer value3.DecRef()

	assert.NotNil(t, value3)

	err := cpy.PyDict_SetItem(dict, key2, value2)
	assert.Zero(t, err)

	err = cpy.PyDict_SetItemString(dict, key1, value1)
	assert.Zero(t, err)

	assert.Equal(t, value2, cpy.PyDict_GetItem(dict, key2))
	assert.Equal(t, value2, cpy.PyDict_SetDefault(dict, key2, cpy.Py_None))
	assert.Equal(t, value1, cpy.PyDict_GetItemString(dict, key1))
	assert.Nil(t, cpy.PyDict_GetItemWithError(dict, key3))

	b := cpy.PyDict_Contains(dict, key2) != 0

	assert.True(t, b)
	assert.Equal(t, 2, cpy.PyDict_Size(dict))

	keys := cpy.PyDict_Keys(dict)
	defer keys.DecRef()

	assert.True(t, cpy.PyList_Check(keys))

	values := cpy.PyDict_Values(dict)
	defer values.DecRef()

	assert.True(t, cpy.PyList_Check(values))

	items := cpy.PyDict_Items(dict)
	defer items.DecRef()

	assert.True(t, cpy.PyList_Check(items))

	err = cpy.PyDict_SetItem(dict, key3, value3)
	assert.Zero(t, err)

	newDict := cpy.PyDict_Copy(dict)
	defer newDict.DecRef()

	assert.Equal(t, 3, cpy.PyDict_Size(newDict))

	err = cpy.PyDict_DelItem(dict, key2)

	assert.Zero(t, err)

	err = cpy.PyDict_DelItemString(dict, key1)

	assert.Zero(t, err)
	assert.Equal(t, 1, cpy.PyDict_Size(dict))

	cpy.PyDict_Clear(dict)

	assert.Equal(t, 0, cpy.PyDict_Size(dict))
}
