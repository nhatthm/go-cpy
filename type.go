package cpy3

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import "unsafe"

// Type is the type object for type objects; it is the same object as type in the Python layer.
//
// Reference: https://docs.python.org/3/c-api/type.html#c.PyType_Type
var Type = togo((*C.PyObject)(unsafe.Pointer(&C.PyType_Type)))

// PyType_Check returns non-zero if the object o is a type object, including instances of types derived from the
// standard type object. Return 0 in all other cases. This function always succeeds.
//
// Reference: https://docs.python.org/3/c-api/type.html#c.PyType_Check
func PyType_Check(o *PyObject) bool {
	return C._go_PyType_Check(toc(o)) != 0
}

// PyType_CheckExact returns non-zero if the object o is a type object, but not a subtype of the standard type object.
// Return 0 in all other cases. This function always succeeds.
//
// Reference: https://docs.python.org/3/c-api/type.html#c.PyType_CheckExact
func PyType_CheckExact(o *PyObject) bool {
	return C._go_PyType_CheckExact(toc(o)) != 0
}
