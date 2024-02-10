package cpy3

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import "unsafe"

// ByteArray is an instance of PyTypeObject represents the Python bytearray type; it is the same object as `bytearray`
// in the Python layer.
//
// Reference: https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_Type
var ByteArray = togo((*C.PyObject)(unsafe.Pointer(&C.PyByteArray_Type)))

// PyByteArray_Check returns true if the object o is a bytearray object or an instance of a subtype of the `bytearray`
// type. This function always succeeds.
//
// Reference: https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_Check
func PyByteArray_Check(o *PyObject) bool {
	return C._go_PyByteArray_Check(toc(o)) != 0
}

// PyByteArray_CheckExact returns true if the object o is a bytearray object, but not an instance of a subtype of the
// `bytearray` type. This function always succeeds.
//
// Reference: https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_CheckExact
func PyByteArray_CheckExact(o *PyObject) bool {
	return C._go_PyByteArray_CheckExact(toc(o)) != 0
}

// PyByteArray_FromObject returns a new `bytearray` object from any object, o, that implements the buffer protocol.
//
// Reference: https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_FromObject
func PyByteArray_FromObject(o *PyObject) *PyObject {
	return togo(C.PyByteArray_FromObject(toc(o)))
}

// PyByteArray_FromStringAndSize creates a new `bytearray` object from `string` and its length, `len`.
// On failure, `NULL` is returned.
//
// Reference: https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_FromStringAndSize
func PyByteArray_FromStringAndSize(str string) *PyObject {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return togo(C.PyByteArray_FromStringAndSize(cstr, C.Py_ssize_t(len(str))))
}

// PyByteArray_Concat concatenates `bytearray`s a and b and return a new `bytearray` with the result.
//
// Reference: https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_Concat
func PyByteArray_Concat(a, b *PyObject) *PyObject {
	return togo(C.PyByteArray_Concat(toc(a), toc(b)))
}

// PyByteArray_Size returns the size of bytearray after checking for a NULL pointer.
//
// Reference: https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_Size
func PyByteArray_Size(o *PyObject) int {
	return int(C.PyByteArray_Size(toc(o)))
}

// PyByteArray_AsString returns the contents of `bytearray` as a char array after checking for a `NULL` pointer.
// The returned array always has an extra null byte appended.
//
// Reference: https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_AsString
func PyByteArray_AsString(o *PyObject) string {
	return C.GoStringN(C.PyByteArray_AsString(toc(o)), C.int(C.PyByteArray_Size(toc(o))))
}

// PyByteArray_Resize resizes the internal buffer of bytearray to len.
//
// Reference: https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_Resize
func PyByteArray_Resize(bytearray *PyObject, len int) {
	C.PyByteArray_Resize(toc(bytearray), C.Py_ssize_t(len))
}
