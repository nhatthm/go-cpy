package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestUnicodeNew(t *testing.T) {
	cpy3.Py_Initialize()

	s := cpy3.PyUnicode_New(20, 'z')
	defer s.DecRef()

	assert.NotNil(t, s)
}

func TestUnicodeFromString(t *testing.T) {
	cpy3.Py_Initialize()

	u := cpy3.PyUnicode_FromString("aaa")
	defer u.DecRef()

	assert.True(t, cpy3.PyUnicode_Check(u))
	assert.True(t, cpy3.PyUnicode_CheckExact(u))
	assert.Equal(t, 3, cpy3.PyUnicode_GetLength(u))
}

func TestUnicodeFromEncodedObject(t *testing.T) {
	cpy3.Py_Initialize()

	b := cpy3.PyBytes_FromString("bbb")
	defer b.DecRef()

	assert.NotNil(t, b)

	ub := cpy3.PyUnicode_FromEncodedObject(b, "utf-8", "strict")
	defer ub.DecRef()

	assert.NotNil(t, ub)
}

func TestUnicodeChar(t *testing.T) {
	cpy3.Py_Initialize()

	u := cpy3.PyUnicode_FromString("aaa")
	defer u.DecRef()

	assert.True(t, cpy3.PyUnicode_Check(u))
	assert.True(t, cpy3.PyUnicode_CheckExact(u))
	assert.Equal(t, 0, cpy3.PyUnicode_WriteChar(u, 1, 'd'))
	assert.Equal(t, 'd', cpy3.PyUnicode_ReadChar(u, 1))
}

func TestUnicodeFill(t *testing.T) {
	cpy3.Py_Initialize()

	u := cpy3.PyUnicode_FromString("aaa")
	defer u.DecRef()

	assert.True(t, cpy3.PyUnicode_Check(u))
	assert.True(t, cpy3.PyUnicode_CheckExact(u))
	assert.Equal(t, 3, cpy3.PyUnicode_Fill(u, 0, 3, 'c'))
	assert.Equal(t, "ccc", cpy3.PyUnicode_AsUTF8(u))
}

func TestUnicodeCopyCharacters(t *testing.T) {
	cpy3.Py_Initialize()

	u := cpy3.PyUnicode_FromString("aaa")
	defer u.DecRef()

	assert.True(t, cpy3.PyUnicode_Check(u))
	assert.True(t, cpy3.PyUnicode_CheckExact(u))

	b := cpy3.PyBytes_FromString("bbb")
	defer b.DecRef()

	assert.NotNil(t, b)

	ub := cpy3.PyUnicode_FromEncodedObject(b, "utf-8", "strict")
	defer ub.DecRef()

	assert.NotNil(t, ub)
	assert.Equal(t, 3, cpy3.PyUnicode_CopyCharacters(ub, u, 0, 0, 3))
	assert.Equal(t, "aaa", cpy3.PyUnicode_AsUTF8(ub))
}

func TestUnicodeSubstring(t *testing.T) {
	cpy3.Py_Initialize()

	u := cpy3.PyUnicode_FromString("aaa")
	defer u.DecRef()

	assert.True(t, cpy3.PyUnicode_Check(u))
	assert.True(t, cpy3.PyUnicode_CheckExact(u))

	sub := cpy3.PyUnicode_Substring(u, 0, 2)
	defer sub.DecRef()

	assert.NotNil(t, sub)
	assert.Equal(t, "aa", cpy3.PyUnicode_AsUTF8(sub))
}
