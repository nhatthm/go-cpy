package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestSequenceContains_Dict(t *testing.T) {
	cpy3.Py_Initialize()

	dict := cpy3.PyDict_New()
	defer dict.DecRef()

	dict.SetItem(cpy3.PyUnicode_FromString("foo"), cpy3.PyUnicode_FromString("bar"))

	assert.Equal(t, 1, cpy3.PySequence_Contains(dict, cpy3.PyUnicode_FromString("foo")))
	assert.Equal(t, 0, cpy3.PySequence_Contains(dict, cpy3.PyUnicode_FromString("bar")))
}
