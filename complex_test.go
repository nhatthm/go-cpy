package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestComplex(t *testing.T) {
	cpy3.Py_Initialize()

	nReal := 2.
	nImaginary := 5.

	nComplex := cpy3.PyComplex_FromDoubles(nReal, nImaginary)
	defer nComplex.DecRef()

	assert.True(t, cpy3.PyComplex_Check(nComplex))
	assert.True(t, cpy3.PyComplex_CheckExact(nComplex))

	assert.Equal(t, nReal, cpy3.PyComplex_RealAsDouble(nComplex))
	assert.Equal(t, nImaginary, cpy3.PyComplex_ImagAsDouble(nComplex))
}
