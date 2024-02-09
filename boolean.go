package cpy3

/*
#include "Python.h"
#include "macro.h"
*/
import "C"

import (
	"unsafe"
)

// Python boolean constants.
var (
	Py_False = togo(C.Py_False)
	Py_True  = togo(C.Py_True)
)

// Bool is an instance of PyTypeObject represents the Python boolean type; it is the same object as `bool` in the
// Python layer.
//
// Reference: https://docs.python.org/3/c-api/bool.html#c.PyBool_Type
var Bool = togo((*C.PyObject)(unsafe.Pointer(&C.PyBool_Type)))

// PyBool_Check returns true if o is of type PyBool_Type. This function always succeeds.
//
// Reference: https://docs.python.org/3/c-api/bool.html#c.PyBool_Check
func PyBool_Check(o *PyObject) bool {
	return C._go_PyBool_Check(toc(o)) != 0
}

// PyBool_FromLong returns Py_True or Py_False, depending on the truth value of v.
//
// Reference: https://docs.python.org/3/c-api/bool.html#c.PyBool_FromLong
func PyBool_FromLong(v int) *PyObject {
	return togo(C.PyBool_FromLong(C.long(v)))
}
