package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequenceContains_Dict(t *testing.T) {
	Py_Initialize()

	dict := PyDict_New()

	dict.SetItem(PyUnicode_FromString("foo"), PyUnicode_FromString("bar"))

	assert.Equal(t, 1, PySequence_Contains(dict, PyUnicode_FromString("foo")))
	assert.Equal(t, 0, PySequence_Contains(dict, PyUnicode_FromString("bar")))
}
