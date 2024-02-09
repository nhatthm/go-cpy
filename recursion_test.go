package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestRecursion(t *testing.T) {
	cpy3.Py_Initialize()

	assert.Zero(t, cpy3.Py_EnterRecursiveCall("in test function"))

	cpy3.Py_LeaveRecursiveCall()
}
