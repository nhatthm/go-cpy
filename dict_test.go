package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestDict(t *testing.T) {
	cpy3.Py_Initialize()

	dict := cpy3.PyDict_New()
	defer dict.DecRef()

	assert.True(t, cpy3.PyDict_Check(dict))
	assert.True(t, cpy3.PyDict_CheckExact(dict))

	proxy := cpy3.PyDictProxy_New(dict)
	defer proxy.DecRef()

	assert.NotNil(t, proxy)

	key1 := "key1"

	value1 := cpy3.PyUnicode_FromString("value1")
	defer value1.DecRef()

	assert.NotNil(t, value1)

	key2 := cpy3.PyUnicode_FromString("key2")
	defer key2.DecRef()

	assert.NotNil(t, key2)

	value2 := cpy3.PyUnicode_FromString("value2")
	defer value2.DecRef()

	assert.NotNil(t, value2)

	key3 := cpy3.PyUnicode_FromString("key3")
	defer key3.DecRef()

	assert.NotNil(t, key3)

	value3 := cpy3.PyUnicode_FromString("value3")
	defer value3.DecRef()

	assert.NotNil(t, value3)

	err := cpy3.PyDict_SetItem(dict, key2, value2)
	assert.Zero(t, err)

	err = cpy3.PyDict_SetItemString(dict, key1, value1)
	assert.Zero(t, err)

	assert.Equal(t, value2, cpy3.PyDict_GetItem(dict, key2))
	assert.Equal(t, value2, cpy3.PyDict_SetDefault(dict, key2, cpy3.Py_None))
	assert.Equal(t, value1, cpy3.PyDict_GetItemString(dict, key1))
	assert.Nil(t, cpy3.PyDict_GetItemWithError(dict, key3))

	b := cpy3.PyDict_Contains(dict, key2) != 0

	assert.True(t, b)
	assert.Equal(t, 2, cpy3.PyDict_Size(dict))

	keys := cpy3.PyDict_Keys(dict)
	defer keys.DecRef()

	assert.True(t, cpy3.PyList_Check(keys))

	values := cpy3.PyDict_Values(dict)
	defer values.DecRef()

	assert.True(t, cpy3.PyList_Check(values))

	items := cpy3.PyDict_Items(dict)
	defer items.DecRef()

	assert.True(t, cpy3.PyList_Check(items))

	err = cpy3.PyDict_SetItem(dict, key3, value3)
	assert.Zero(t, err)

	newDict := cpy3.PyDict_Copy(dict)
	defer newDict.DecRef()

	assert.Equal(t, 3, cpy3.PyDict_Size(newDict))

	err = cpy3.PyDict_DelItem(dict, key2)

	assert.Zero(t, err)

	err = cpy3.PyDict_DelItemString(dict, key1)

	assert.Zero(t, err)
	assert.Equal(t, 1, cpy3.PyDict_Size(dict))

	cpy3.PyDict_Clear(dict)

	assert.Equal(t, 0, cpy3.PyDict_Size(dict))
}
