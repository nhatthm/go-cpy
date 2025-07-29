package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestSequenceContains_Dict(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	dict := cpy.PyDict_New()
	defer dict.DecRef()

	dict.SetItem(cpy.PyUnicode_FromString("foo"), cpy.PyUnicode_FromString("bar"))

	assert.Equal(t, 1, cpy.PySequence_Contains(dict, cpy.PyUnicode_FromString("foo")))
	assert.Equal(t, 0, cpy.PySequence_Contains(dict, cpy.PyUnicode_FromString("bar")))
}
