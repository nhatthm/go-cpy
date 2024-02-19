package cpy

/*
#include "Python.h"
#include "macro.h"
*/
import "C"

import (
	"unsafe"
)

// Long is an instance of PyTypeObject represents the Python integer type. This is the same object as int in the
// Python layer.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_Type
var Long = togo((*C.PyObject)(unsafe.Pointer(&C.PyLong_Type)))

// PyLong_Check returns true if its argument is a PyLongObject or a subtype of PyLongObject. This function always
// succeeds.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_Check
func PyLong_Check(p *PyObject) bool {
	return C._go_PyLong_Check(toc(p)) != 0
}

// PyLong_CheckExact returns true if its argument is a PyLongObject, but not a subtype of PyLongObject. This function
// always succeeds.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_CheckExact
func PyLong_CheckExact(p *PyObject) bool {
	return C._go_PyLong_CheckExact(toc(p)) != 0
}

// PyLong_FromLong returns a new PyLongObject object from v, or NULL on failure.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_FromLong
func PyLong_FromLong(v int) *PyObject {
	return togo(C.PyLong_FromLong(C.long(v)))
}

// PyLong_FromUnsignedLong returns a new PyLongObject object from a C unsigned long, or NULL on failure.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_FromUnsignedLong
func PyLong_FromUnsignedLong(v uint) *PyObject {
	return togo(C.PyLong_FromUnsignedLong(C.ulong(v)))
}

// PyLong_FromLongLong returns a new PyLongObject object from a C unsigned long, or NULL on failure.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_FromLongLong
func PyLong_FromLongLong(v int64) *PyObject {
	return togo(C.PyLong_FromLongLong(C.longlong(v)))
}

// PyLong_FromUnsignedLongLong returns a new PyLongObject object from a C unsigned long long, or NULL on failure.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_FromUnsignedLongLong
func PyLong_FromUnsignedLongLong(v uint64) *PyObject {
	return togo(C.PyLong_FromUnsignedLongLong(C.ulonglong(v)))
}

// PyLong_FromDouble returns a new PyLongObject object from the integer part of v, or NULL on failure.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_FromDouble
func PyLong_FromDouble(v float64) *PyObject {
	return togo(C.PyLong_FromDouble(C.double(v)))
}

// PyLong_FromString returns a new PyLongObject based on the string value in str, which is interpreted according to the
// radix in base, or NULL on failure. If pend is non-NULL, *pend will point to the end of str on success or to the first
// character that could not be processed on error. If base is 0, str is interpreted using the Integer literals
// definition; in this case, leading zeros in a non-zero decimal number raises a ValueError. If base is not 0, it must
// be between 2 and 36, inclusive. Leading and trailing whitespace and single underscores after a base specifier and
// between digits are ignored. If there are no digits or str is not NULL-terminated following the digits and trailing
// whitespace, ValueError will be raised.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_FromString
func PyLong_FromString(str string, base int) *PyObject {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return togo(C.PyLong_FromString(cstr, nil, C.int(base)))
}

// PyLong_FromUnicodeObject converts a sequence of Unicode digits in the string u to a Python integer value.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_FromUnicodeObject
func PyLong_FromUnicodeObject(u *PyObject, base int) *PyObject {
	return togo(C.PyLong_FromUnicodeObject(toc(u), C.int(base)))
}

// PyLong_FromGoInt ensures the go integer type does not overflow.
func PyLong_FromGoInt(v int) *PyObject {
	return togo(C.PyLong_FromLongLong(C.longlong(v)))
}

// PyLong_FromGoUint ensures the go integer type does not overflow.
func PyLong_FromGoUint(v uint) *PyObject {
	return togo(C.PyLong_FromUnsignedLongLong(C.ulonglong(v)))
}

// PyLong_FromGoInt64 ensures the go integer type does not overflow.
func PyLong_FromGoInt64(v int64) *PyObject {
	return togo(C.PyLong_FromLongLong(C.longlong(v)))
}

// PyLong_FromGoUint64 ensures the go integer type does not overflow.
func PyLong_FromGoUint64(v uint64) *PyObject {
	return togo(C.PyLong_FromUnsignedLongLong(C.ulonglong(v)))
}

// PyLong_FromGoFloat64 ensures the go integer type does not overflow.
func PyLong_FromGoFloat64(v float64) *PyObject {
	return togo(C.PyLong_FromDouble(C.double(v)))
}

// PyLong_AsLong returns a C long representation of obj. If obj is not an instance of PyLongObject, first call its
// __index__() method (if present) to convert it to a PyLongObject.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_AsLong
func PyLong_AsLong(obj *PyObject) int {
	return int(C.PyLong_AsLong(toc(obj)))
}

// PyLong_AsLongAndOverflow returns a C long representation of obj. If obj is not an instance of PyLongObject, first
// call its __index__() method (if present) to convert it to a PyLongObject.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_AsLongAndOverflow
func PyLong_AsLongAndOverflow(obj *PyObject) (int, int) {
	overflow := C.int(0)
	ret := C.PyLong_AsLongAndOverflow(toc(obj), &overflow)

	return int(ret), int(overflow)
}

// PyLong_AsLongLong returns a C long long representation of obj. If obj is not an instance of PyLongObject, first call
// its __index__() method (if present) to convert it to a PyLongObject.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_AsLongLong
// nolint: dupword
func PyLong_AsLongLong(obj *PyObject) int64 {
	return int64(C.PyLong_AsLongLong(toc(obj)))
}

// PyLong_AsLongLongAndOverflow returns a C long long representation of obj. If obj is not an instance of PyLongObject,
// first call its __index__() method (if present) to convert it to a PyLongObject.
//
// If the value of obj is greater than LLONG_MAX or less than LLONG_MIN, set *overflow to 1 or -1, respectively, and
// return -1; otherwise, set *overflow to 0. If any other exception occurs set *overflow to 0 and return -1 as usual.
//
// Returns -1 on error. Use PyErr_Occurred() to disambiguate.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_AsLongLongAndOverflow
// nolint: dupword
func PyLong_AsLongLongAndOverflow(obj *PyObject) (int64, int) {
	overflow := C.int(0)
	ret := C.PyLong_AsLongLongAndOverflow(toc(obj), &overflow)

	return int64(ret), int(overflow)
}

// PyLong_AsUnsignedLong returns a C long long representation of obj. If obj is not an instance of PyLongObject, first
// call its __index__() method (if present) to convert it to a PyLongObject.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_AsUnsignedLong
// nolint: dupword
func PyLong_AsUnsignedLong(obj *PyObject) uint {
	return uint(C.PyLong_AsUnsignedLong(toc(obj)))
}

// PyLong_AsUnsignedLongLong returns a C unsigned long long representation of pylong. pylong must be an instance of
// PyLongObject.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_AsUnsignedLongLong
// nolint: dupword
func PyLong_AsUnsignedLongLong(obj *PyObject) uint64 {
	return uint64(C.PyLong_AsUnsignedLongLong(toc(obj)))
}

// PyLong_AsUnsignedLongMask returns a C unsigned long representation of obj. If obj is not an instance of PyLongObject,
// first call its __index__() method (if present) to convert it to a PyLongObject.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_AsUnsignedLongMask
func PyLong_AsUnsignedLongMask(obj *PyObject) uint {
	return uint(C.PyLong_AsUnsignedLongMask(toc(obj)))
}

// PyLong_AsUnsignedLongLongMask returns a C unsigned long long representation of obj. If obj is not an instance of
// PyLongObject, first call its __index__() method (if present) to convert it to a PyLongObject.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_AsUnsignedLongLongMask
// nolint: dupword
func PyLong_AsUnsignedLongLongMask(obj *PyObject) uint64 {
	return uint64(C.PyLong_AsUnsignedLongLongMask(toc(obj)))
}

// PyLong_AsDouble returns a C double representation of pylong. pylong must be an instance of PyLongObject.
//
// Reference: https://docs.python.org/3/c-api/long.html#c.PyLong_AsDouble
func PyLong_AsDouble(obj *PyObject) float64 {
	return float64(C.PyLong_AsDouble(toc(obj)))
}
