package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestBoolCheck(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	assert.True(t, cpy.PyBool_Check(cpy.Py_True))
	assert.True(t, cpy.PyBool_Check(cpy.Py_False))
}

func TestBoolFromLong(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	assert.Equal(t, cpy.Py_True, cpy.PyBool_FromLong(1))
	assert.Equal(t, cpy.Py_False, cpy.PyBool_FromLong(0))
}
