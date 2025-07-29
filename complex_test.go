package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestComplex(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	nReal := 2.
	nImaginary := 5.

	nComplex := cpy.PyComplex_FromDoubles(nReal, nImaginary)
	defer nComplex.DecRef()

	assert.True(t, cpy.PyComplex_Check(nComplex))
	assert.True(t, cpy.PyComplex_CheckExact(nComplex))

	assert.InDelta(t, nReal, cpy.PyComplex_RealAsDouble(nComplex), 0.01)
	assert.InDelta(t, nImaginary, cpy.PyComplex_ImagAsDouble(nComplex), 0.01)
}
