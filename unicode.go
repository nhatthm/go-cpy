package cpy3

/*
#include "Python.h"
#include "macro.h"
*/
import "C"

import (
	"unsafe"
)

// Unicode is an instance of PyTypeObject represents the Python Unicode type. It is exposed to Python code as str.
//
// Reference: https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_Type
var Unicode = togo((*C.PyObject)(unsafe.Pointer(&C.PyUnicode_Type)))

// PyUnicode_Check returns true if the object obj is a Unicode object or an instance of a Unicode subtype. This function
// always succeeds.
//
// Reference: https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_Check
func PyUnicode_Check(o *PyObject) bool {
	return C._go_PyUnicode_Check(toc(o)) != 0
}

// PyUnicode_CheckExact returns true if the object obj is a Unicode object, but not an instance of a subtype. This
// function always succeeds.
//
// Reference: https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_CheckExact
func PyUnicode_CheckExact(o *PyObject) bool {
	return C._go_PyUnicode_CheckExact(toc(o)) != 0
}

// PyUnicode_New creates a new Unicode object. maxchar should be the true maximum code point to be placed in the string.
// As an approximation, it can be rounded up to the nearest value in the sequence 127, 255, 65535, 1114111.
//
// This is the recommended way to allocate a new Unicode object. Objects created using this function are not resizable.
//
// Reference: https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_New
func PyUnicode_New(size int, maxchar rune) *PyObject {
	return togo(C.PyUnicode_New(C.Py_ssize_t(size), C.Py_UCS4(maxchar)))
}

// PyUnicode_FromString creates a Unicode object from a UTF-8 encoded null-terminated char buffer str.
//
// Reference: https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_FromString
func PyUnicode_FromString(u string) *PyObject {
	cu := C.CString(u)
	defer C.free(unsafe.Pointer(cu))

	return togo(C.PyUnicode_FromString(cu))
}

// PyUnicode_FromEncodedObject decodes an encoded object obj to a Unicode object.
//
// bytes, bytearray and other bytes-like objects are decoded according to the given encoding and using the error
// handling defined by errors. Both can be NULL to have the interface use the default values
// (see Built-in Codecs for details).
//
// All other objects, including Unicode objects, cause a TypeError to be set.
//
// The API returns NULL if there was an error. The caller is responsible for decref’ing the returned objects.
//
// Reference: https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_FromEncodedObject
func PyUnicode_FromEncodedObject(obj *PyObject, encoding, errors string) *PyObject {
	cencoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(cencoding))

	cerrors := C.CString(errors)
	defer C.free(unsafe.Pointer(cerrors))

	return togo(C.PyUnicode_FromEncodedObject(toc(obj), cencoding, cerrors))
}

// PyUnicode_GetLength returns the length of the Unicode object, in code points.
//
// Reference: https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_GetLength
func PyUnicode_GetLength(unicode *PyObject) int {
	return int(C.PyUnicode_GetLength(toc(unicode)))
}

// PyUnicode_CopyCharacters copy characters from one Unicode object into another. This function performs character
// conversion when necessary and falls back to memcpy() if possible. Returns -1 and sets an exception on error,
// otherwise returns the number of copied characters.
//
// Reference: https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_CopyCharacters
func PyUnicode_CopyCharacters(to, from *PyObject, to_start, from_start, how_many int) int {
	return int(C.PyUnicode_CopyCharacters(toc(to), C.Py_ssize_t(to_start), toc(from), C.Py_ssize_t(from_start), C.Py_ssize_t(how_many)))
}

// PyUnicode_Fill fills a string with a character: write fill_char into unicode[start:start+length].
//
// Fail if fill_char is bigger than the string maximum character, or if the string has more than 1 reference.
//
// The function returns the number of written character, or return -1 and raise an exception on error.
//
// Reference: https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_Fill
func PyUnicode_Fill(unicode *PyObject, start, length int, fill_char rune) int {
	return int(C.PyUnicode_Fill(toc(unicode), C.Py_ssize_t(start), C.Py_ssize_t(length), C.Py_UCS4(fill_char)))
}

// PyUnicode_WriteChar writes a character to a string. The string must have been created through PyUnicode_New().
// Since Unicode strings are supposed to be immutable, the string must not be shared, or have been hashed yet.
//
// This function checks that unicode is a Unicode object, that the index is not out of bounds, and that the object can
// be modified safely (i.e. that it its reference count is one).
//
// Reference: https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_WriteChar
func PyUnicode_WriteChar(unicode *PyObject, index int, character rune) int {
	return int(C.PyUnicode_WriteChar(toc(unicode), C.Py_ssize_t(index), C.Py_UCS4(character)))
}

// PyUnicode_ReadChar reads a character from a string. This function checks that unicode is a Unicode object and the
// index is not out of bounds, in contrast to PyUnicode_READ_CHAR(), which performs no error checking.
//
// Reference: https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_ReadChar
func PyUnicode_ReadChar(unicode *PyObject, index int) rune {
	return rune(C.PyUnicode_ReadChar(toc(unicode), C.Py_ssize_t(index)))
}

// PyUnicode_Substring returns a substring of unicode, from character index start (included) to character index end
// (excluded). Negative indices are not supported.
//
// Reference: https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_Substring
func PyUnicode_Substring(unicode *PyObject, start, end int) *PyObject {
	return togo(C.PyUnicode_Substring(toc(unicode), C.Py_ssize_t(start), C.Py_ssize_t(end)))
}

// PyUnicode_AsUTF8 is the same as PyUnicode_AsUTF8AndSize(), but does not store the size.
//
// Reference: https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_AsUTF8
func PyUnicode_AsUTF8(unicode *PyObject) string {
	cutf8 := C.PyUnicode_AsUTF8(toc(unicode))
	return C.GoString(cutf8)
}
