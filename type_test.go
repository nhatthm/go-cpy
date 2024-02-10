package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestTypeCheck(t *testing.T) {
	cpy3.Py_Initialize()

	assert.True(t, cpy3.PyType_Check(cpy3.Type))
	assert.True(t, cpy3.PyType_CheckExact(cpy3.Type))
}
