package cpy3

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import "unsafe"

// Bytes is an instance of PyTypeObject represents the Python `bytes` type; it is the same object as `bytes` in the
// Python layer.
//
// Reference: https://docs.python.org/3/c-api/bytes.html#c.PyBytes_Type
var Bytes = togo((*C.PyObject)(unsafe.Pointer(&C.PyBytes_Type)))

// PyBytes_Check returns true if the object o is a `bytes` object or an instance of a subtype of the `bytes` type.
// This function always succeeds.
//
// Reference: https://docs.python.org/3/c-api/bytes.html#c.PyBytes_Check
func PyBytes_Check(o *PyObject) bool {
	return C._go_PyBytes_Check(toc(o)) != 0
}

// PyBytes_CheckExact returns true if the object o is a `bytes` object, but not an instance of a subtype of the `bytes`
// type. This function always succeeds.
//
// Reference: https://docs.python.org/3/c-api/bytes.html#c.PyBytes_CheckExact
func PyBytes_CheckExact(o *PyObject) bool {
	return C._go_PyBytes_CheckExact(toc(o)) != 0
}

// PyBytes_FromString returns a new bytes object with a copy of the string v as value on success, and NULL on failure.
// The parameter v must not be NULL; it will not be checked.
//
// Reference: https://docs.python.org/3/c-api/bytes.html#c.PyBytes_FromString
func PyBytes_FromString(str string) *PyObject {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return togo(C.PyBytes_FromString(cstr))
}

// PyBytes_FromStringAndSize returns a new `bytes` object with a copy of the string v as value and length len on
// success, and `NULL` on failure. If v is `NULL`, the contents of the bytes object are uninitialized.
//
// Reference: https://docs.python.org/3/c-api/bytes.html#c.PyBytes_FromStringAndSize
func PyBytes_FromStringAndSize(str string) *PyObject {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return togo(C.PyBytes_FromStringAndSize(cstr, C.Py_ssize_t(len(str))))
}

// PyBytes_FromObject returns the `bytes` representation of object o that implements the buffer protocol.
//
// Reference: https://docs.python.org/3/c-api/bytes.html#c.PyBytes_FromObject
func PyBytes_FromObject(o *PyObject) *PyObject {
	return togo(C.PyBytes_FromObject(toc(o)))
}

// PyBytes_Size returns the length of the bytes in bytes object o.
//
// Reference: https://docs.python.org/3/c-api/bytes.html#c.PyBytes_Size
func PyBytes_Size(o *PyObject) int {
	return int(C.PyBytes_Size(toc(o)))
}

// PyBytes_AsString returns a pointer to the contents of o. The pointer refers to the internal buffer of o, which
// consists of len(o) + 1 bytes. The last byte in the buffer is always null, regardless of whether there are any other
// null bytes. The data must not be modified in any way, unless the object was just created using
// PyBytes_FromStringAndSize(NULL, size). It must not be deallocated.
//
// If o is not a bytes object at all, PyBytes_AsString() returns NULL and raises TypeError.
//
// Reference: https://docs.python.org/3/c-api/bytes.html#c.PyBytes_AsString
func PyBytes_AsString(o *PyObject) string {
	return C.GoStringN(C.PyBytes_AsString(toc(o)), C.int(C.PyBytes_Size(toc(o))))
}

// PyBytes_Concat creates a new bytes object in *bytes containing the contents of newpart appended to bytes; the caller
// will own the new reference. The reference to the old value of bytes will be stolen. If the new object cannot be
// created, the old reference to bytes will still be discarded and the value of *bytes will be set to NULL; the
// appropriate exception will be set.
//
// Reference: https://docs.python.org/3/c-api/bytes.html#c.PyBytes_Concat
func PyBytes_Concat(bytes, newpart *PyObject) *PyObject {
	cbytes := toc(bytes)

	C.PyBytes_Concat(&cbytes, toc(newpart)) //nolint: gocritic

	return togo(cbytes)
}

// PyBytes_ConcatAndDel creates a new bytes object in *bytes containing the contents of newpart appended to bytes. This
// version releases the strong reference to newpart (i.e. decrements its reference count).
//
// Reference: https://docs.python.org/3/c-api/bytes.html#c.PyBytes_ConcatAndDel
func PyBytes_ConcatAndDel(bytes, newpart *PyObject) *PyObject {
	cbytes := toc(bytes)

	C.PyBytes_ConcatAndDel(&cbytes, toc(newpart)) //nolint: gocritic

	return togo(cbytes)
}

// PyBytes_FromByteSlice uses https://docs.python.org/3/c-api/bytes.html#c.PyBytes_FromStringAndSize but with []byte
func PyBytes_FromByteSlice(bytes []byte) *PyObject {
	pbytes := C.CBytes(bytes)
	defer C.free(pbytes)

	cstr := (*C.char)(pbytes)

	return togo(C.PyBytes_FromStringAndSize(cstr, C.Py_ssize_t(len(bytes))))
}

// PyBytes_AsByteSlice is equivalent to PyBytes_AsString but returns byte slices.
func PyBytes_AsByteSlice(o *PyObject) []byte {
	cstr := C.PyBytes_AsString(toc(o))
	size := C.PyBytes_Size(toc(o))

	return C.GoBytes(unsafe.Pointer(cstr), C.int(size))
}
