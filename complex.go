package cpy3

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import "unsafe"

// Complex is an instance of PyTypeObject represents the Python `complex` number type. It is the same object as
// `complex` in the Python layer.
//
// Reference: https://docs.python.org/3/c-api/complex.html#c.PyComplex_Type
var Complex = togo((*C.PyObject)(unsafe.Pointer(&C.PyComplex_Type)))

// PyComplex_Check returns true if its argument is a PyComplexObject or a subtype of PyComplexObject. This function
// always succeeds.
//
// Reference: https://docs.python.org/3/c-api/complex.html#c.PyComplex_Check
func PyComplex_Check(p *PyObject) bool {
	return C._go_PyComplex_Check(toc(p)) != 0
}

// PyComplex_CheckExact returns true if its argument is a PyComplexObject or a subtype of PyComplexObject. This
// function always succeeds.
//
// Reference: https://docs.python.org/3/c-api/complex.html#c.PyComplex_CheckExact
func PyComplex_CheckExact(p *PyObject) bool {
	return C._go_PyComplex_CheckExact(toc(p)) != 0
}

// PyComplex_FromDoubles returns a new PyComplexObject object from real and imag.
//
// Reference: https://docs.python.org/3/c-api/complex.html#c.PyComplex_FromDoubles
func PyComplex_FromDoubles(realNumber, imaginaryUnit float64) *PyObject {
	return togo(C.PyComplex_FromDoubles(C.double(realNumber), C.double(imaginaryUnit)))
}

// PyComplex_RealAsDouble returns the real part of op as a C double.
//
// Reference: https://docs.python.org/3/c-api/complex.html#c.PyComplex_RealAsDouble
func PyComplex_RealAsDouble(op *PyObject) float64 {
	return float64(C.PyComplex_RealAsDouble(toc(op)))
}

// PyComplex_ImagAsDouble returns the imaginary part of op as a C double.
//
// Reference: https://docs.python.org/3/c-api/complex.html#c.PyComplex_ImagAsDouble
func PyComplex_ImagAsDouble(op *PyObject) float64 {
	return float64(C.PyComplex_ImagAsDouble(toc(op)))
}
