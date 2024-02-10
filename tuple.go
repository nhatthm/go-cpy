package cpy3

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import "unsafe"

// Tuple is an instance of PyTypeObject represents the Python tuple type; it is the same object as tuple in the Python
// layer.
//
// Reference: https://docs.python.org/3/c-api/tuple.html#c.PyTuple_Type
var Tuple = togo((*C.PyObject)(unsafe.Pointer(&C.PyTuple_Type)))

// PyTuple_Check returns true if p is a tuple object or an instance of a subtype of the tuple type. This function always
// succeeds.
//
// Reference: https://docs.python.org/3/c-api/tuple.html#c.PyTuple_Check
func PyTuple_Check(p *PyObject) bool {
	return C._go_PyTuple_Check(toc(p)) != 0
}

// PyTuple_CheckExact return true if p is a tuple object, but not an instance of a subtype of the tuple type. This
// function always succeeds.
//
// Reference: https://docs.python.org/3/c-api/tuple.html#c.PyTuple_CheckExact
func PyTuple_CheckExact(p *PyObject) bool {
	return C._go_PyTuple_CheckExact(toc(p)) != 0
}

// PyTuple_New returns a new tuple object of size len, or NULL on failure.
//
// Reference: https://docs.python.org/3/c-api/tuple.html#c.PyTuple_New
func PyTuple_New(len int) *PyObject {
	return togo(C.PyTuple_New(C.Py_ssize_t(len)))
}

// PyTuple_Size returns the size of the tuple.
//
// Reference: https://docs.python.org/3/c-api/tuple.html#c.PyTuple_Size
func PyTuple_Size(p *PyObject) int {
	return int(C.PyTuple_Size(toc(p)))
}

// PyTuple_GetItem returns the object at position pos in the tuple pointed to by p. If pos is negative or out of bounds,
// return NULL and set an IndexError exception.
//
// Reference: https://docs.python.org/3/c-api/tuple.html#c.PyTuple_GetItem
func PyTuple_GetItem(p *PyObject, pos int) *PyObject {
	return togo(C.PyTuple_GetItem(toc(p), C.Py_ssize_t(pos)))
}

// PyTuple_GetSlice returns the slice of the tuple pointed to by p between low and high, or NULL on failure. This is the
// equivalent of the Python expression p[low:high]. Indexing from the end of the tuple is not supported.
//
// Reference: https://docs.python.org/3/c-api/tuple.html#c.PyTuple_GetSlice
func PyTuple_GetSlice(p *PyObject, low, high int) *PyObject {
	return togo(C.PyTuple_GetSlice(toc(p), C.Py_ssize_t(low), C.Py_ssize_t(high)))
}

// PyTuple_SetItem inserts a reference to object o at position pos of the tuple pointed to by p. Return 0 on success.
// If pos is out of bounds, return -1 and set an IndexError exception.
//
// Reference: https://docs.python.org/3/c-api/tuple.html#c.PyTuple_SetItem
func PyTuple_SetItem(p *PyObject, pos int, o *PyObject) int {
	return int(C.PyTuple_SetItem(toc(p), C.Py_ssize_t(pos), toc(o)))
}
