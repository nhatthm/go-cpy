package cpy3

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import "unsafe"

// Float is an instance of PyTypeObject represents the Python floating point type. This is the same object as float in
// the Python layer.
//
// Reference: https://docs.python.org/3/c-api/float.html#c.PyFloat_Type
var Float = togo((*C.PyObject)(unsafe.Pointer(&C.PyFloat_Type)))

// PyFloat_Check returns true if its argument is a PyFloatObject or a subtype of PyFloatObject. This function always
// succeeds.
//
// Reference: https://docs.python.org/3/c-api/float.html#c.PyFloat_Check
func PyFloat_Check(p *PyObject) bool {
	return C._go_PyFloat_Check(toc(p)) != 0
}

// PyFloat_CheckExact returns true if its argument is a PyFloatObject, but not a subtype of PyFloatObject. This
// function always succeeds.
//
// Reference: https://docs.python.org/3/c-api/float.html#c.PyFloat_CheckExact
func PyFloat_CheckExact(p *PyObject) bool {
	return C._go_PyFloat_CheckExact(toc(p)) != 0
}

// PyFloat_FromDouble creates a PyFloatObject object from v, or NULL on failure.
//
// Reference: https://docs.python.org/3/c-api/float.html#c.PyFloat_FromDouble
func PyFloat_FromDouble(v float64) *PyObject {
	return togo(C.PyFloat_FromDouble(C.double(v)))
}

// PyFloat_FromString creates a PyFloatObject object based on the string value in str, or NULL on failure.
//
// Reference: https://docs.python.org/3/c-api/float.html#c.PyFloat_FromString
func PyFloat_FromString(str *PyObject) *PyObject {
	return togo(C.PyFloat_FromString(toc(str)))
}

// PyFloat_AsDouble creates a PyFloatObject object based on the string value in str, or NULL on failure.
//
// Reference: https://docs.python.org/3/c-api/float.html#c.PyFloat_AsDouble
func PyFloat_AsDouble(obj *PyObject) float64 {
	return float64(C.PyFloat_AsDouble(toc(obj)))
}

// PyFloat_GetInfo returns a structseq instance which contains information about the precision, minimum and maximum
// values of a float. It's a thin wrapper around the header file float.h.
//
// Reference: https://docs.python.org/3/c-api/float.html#c.PyFloat_GetInfo
func PyFloat_GetInfo() *PyObject {
	return togo(C.PyFloat_GetInfo())
}

// PyFloat_GetMax returns the maximum representable finite float DBL_MAX as C double.
//
// Reference: https://docs.python.org/3/c-api/float.html#c.PyFloat_GetMax
func PyFloat_GetMax() float64 {
	return float64(C.PyFloat_GetMax())
}

// PyFloat_GetMin returns the minimum normalized positive float DBL_MIN as C double.
//
// Reference: https://docs.python.org/3/c-api/float.html#c.PyFloat_GetMin
func PyFloat_GetMin() float64 {
	return float64(C.PyFloat_GetMin())
}
