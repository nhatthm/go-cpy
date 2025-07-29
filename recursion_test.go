package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestRecursion(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	assert.Zero(t, cpy.Py_EnterRecursiveCall("in test function"))

	cpy.Py_LeaveRecursiveCall()
}
