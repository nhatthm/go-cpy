package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestExceptionNew(t *testing.T) {
	cpy3.Py_Initialize()

	exc := cpy3.PyErr_NewException("test_module.TestException", nil, nil)
	defer exc.DecRef()

	assert.NotNil(t, exc)
}

func TestExceptionNewDoc(t *testing.T) {
	cpy3.Py_Initialize()

	exc := cpy3.PyErr_NewExceptionWithDoc("test_module.TestException", "docstring", nil, nil)
	defer exc.DecRef()

	assert.NotNil(t, exc)
}

func TestExceptionContext(t *testing.T) {
	t.Skip("fatal error: unexpected signal during runtime execution")
	cpy3.Py_Initialize()

	exc := cpy3.PyErr_NewException("test_module.TestException", nil, nil)
	defer exc.DecRef()

	assert.NotNil(t, exc)

	cpy3.PyException_SetContext(exc, cpy3.PyExc_BrokenPipeError)

	assert.Equal(t, cpy3.PyExc_BrokenPipeError, cpy3.PyException_GetContext(exc))
}
