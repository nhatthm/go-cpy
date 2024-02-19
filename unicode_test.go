package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestUnicodeNew(t *testing.T) {
	cpy.Py_Initialize()

	s := cpy.PyUnicode_New(20, 'z')
	defer s.DecRef()

	assert.NotNil(t, s)
}

func TestUnicodeFromString(t *testing.T) {
	cpy.Py_Initialize()

	u := cpy.PyUnicode_FromString("aaa")
	defer u.DecRef()

	assert.True(t, cpy.PyUnicode_Check(u))
	assert.True(t, cpy.PyUnicode_CheckExact(u))
	assert.Equal(t, 3, cpy.PyUnicode_GetLength(u))
}

func TestUnicodeFromEncodedObject(t *testing.T) {
	cpy.Py_Initialize()

	b := cpy.PyBytes_FromString("bbb")
	defer b.DecRef()

	assert.NotNil(t, b)

	ub := cpy.PyUnicode_FromEncodedObject(b, "utf-8", "strict")
	defer ub.DecRef()

	assert.NotNil(t, ub)
}

func TestUnicodeChar(t *testing.T) {
	cpy.Py_Initialize()

	u := cpy.PyUnicode_FromString("aaa")
	defer u.DecRef()

	assert.True(t, cpy.PyUnicode_Check(u))
	assert.True(t, cpy.PyUnicode_CheckExact(u))
	assert.Equal(t, 0, cpy.PyUnicode_WriteChar(u, 1, 'd'))
	assert.Equal(t, 'd', cpy.PyUnicode_ReadChar(u, 1))
}

func TestUnicodeFill(t *testing.T) {
	cpy.Py_Initialize()

	u := cpy.PyUnicode_FromString("aaa")
	defer u.DecRef()

	assert.True(t, cpy.PyUnicode_Check(u))
	assert.True(t, cpy.PyUnicode_CheckExact(u))
	assert.Equal(t, 3, cpy.PyUnicode_Fill(u, 0, 3, 'c'))
	assert.Equal(t, "ccc", cpy.PyUnicode_AsUTF8(u))
}

func TestUnicodeCopyCharacters(t *testing.T) {
	cpy.Py_Initialize()

	u := cpy.PyUnicode_FromString("aaa")
	defer u.DecRef()

	assert.True(t, cpy.PyUnicode_Check(u))
	assert.True(t, cpy.PyUnicode_CheckExact(u))

	b := cpy.PyBytes_FromString("bbb")
	defer b.DecRef()

	assert.NotNil(t, b)

	ub := cpy.PyUnicode_FromEncodedObject(b, "utf-8", "strict")
	defer ub.DecRef()

	assert.NotNil(t, ub)
	assert.Equal(t, 3, cpy.PyUnicode_CopyCharacters(ub, u, 0, 0, 3))
	assert.Equal(t, "aaa", cpy.PyUnicode_AsUTF8(ub))
}

func TestUnicodeSubstring(t *testing.T) {
	cpy.Py_Initialize()

	u := cpy.PyUnicode_FromString("aaa")
	defer u.DecRef()

	assert.True(t, cpy.PyUnicode_Check(u))
	assert.True(t, cpy.PyUnicode_CheckExact(u))

	sub := cpy.PyUnicode_Substring(u, 0, 2)
	defer sub.DecRef()

	assert.NotNil(t, sub)
	assert.Equal(t, "aa", cpy.PyUnicode_AsUTF8(sub))
}
