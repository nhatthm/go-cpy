package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestTypeCheck(t *testing.T) {
	cpy.Py_Initialize()

	assert.True(t, cpy.PyType_Check(cpy.Type))
	assert.True(t, cpy.PyType_CheckExact(cpy.Type))
}

func TestTypeGetName(t *testing.T) {
	cpy.Py_Initialize()

	py := cpy.PyLong_FromGoInt64(42).Type()
	str := py.GetAttrString("__name__").Str()

	assert.Equal(t, "int", cpy.PyUnicode_AsUTF8(str))
}
