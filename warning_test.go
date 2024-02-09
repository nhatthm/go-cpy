package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestWarnEx(t *testing.T) {
	cpy3.Py_Initialize()

	assert.Equal(t, -1, cpy3.PyErr_WarnEx(cpy3.PyExc_RuntimeWarning, "test warning", 3))
}

func TestWarnExplicitObject(t *testing.T) {
	cpy3.Py_Initialize()

	message := cpy3.PyUnicode_FromString("test warning")
	defer message.DecRef()

	filename := cpy3.PyUnicode_FromString("test.py")
	defer filename.DecRef()

	module := cpy3.PyUnicode_FromString("test_module")
	defer module.DecRef()

	assert.Equal(t, -1, cpy3.PyErr_WarnExplicitObject(cpy3.PyExc_RuntimeError, message, filename, 4, module, nil))
}

func TestWarnExplicit(t *testing.T) {
	cpy3.Py_Initialize()

	assert.Equal(t, -1, cpy3.PyErr_WarnExplicit(cpy3.PyExc_RuntimeError, "test warning", "test.py", 4, "test_module", nil))
}
