package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestExceptionNew(t *testing.T) {
	cpy.Py_Initialize()

	exc := cpy.PyErr_NewException("test_module.TestException", nil, nil)
	defer exc.DecRef()

	assert.NotNil(t, exc)
}

func TestExceptionNewDoc(t *testing.T) {
	cpy.Py_Initialize()

	exc := cpy.PyErr_NewExceptionWithDoc("test_module.TestException", "docstring", nil, nil)
	defer exc.DecRef()

	assert.NotNil(t, exc)
}

func TestExceptionContext(t *testing.T) {
	t.Skip("fatal error: unexpected signal during runtime execution")
	cpy.Py_Initialize()

	exc := cpy.PyErr_NewException("test_module.TestException", nil, nil)
	defer exc.DecRef()

	assert.NotNil(t, exc)

	cpy.PyException_SetContext(exc, cpy.PyExc_BrokenPipeError)

	assert.Equal(t, cpy.PyExc_BrokenPipeError, cpy.PyException_GetContext(exc))
}
