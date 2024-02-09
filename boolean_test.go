package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestBoolCheck(t *testing.T) {
	cpy3.Py_Initialize()

	assert.True(t, cpy3.PyBool_Check(cpy3.Py_True))
	assert.True(t, cpy3.PyBool_Check(cpy3.Py_False))
}

func TestBoolFromLong(t *testing.T) {
	cpy3.Py_Initialize()

	assert.Equal(t, cpy3.Py_True, cpy3.PyBool_FromLong(1))
	assert.Equal(t, cpy3.Py_False, cpy3.PyBool_FromLong(0))
}
