package cpy

/*
#include "Python.h"
*/
import "C"

import (
	"unsafe"
)

// PyErr_Clear clears the error indicator. If the error indicator is not set, there is no effect.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_Clear
func PyErr_Clear() {
	C.PyErr_Clear()
}

// PyErr_PrintEx prints a standard traceback to sys.stderr and clear the error indicator. Unless the error is a
// SystemExit, in that case no traceback is printed and the Python process will exit with the error code specified by
// the SystemExit instance.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_PrintEx
func PyErr_PrintEx(setSysLastVars bool) {
	if setSysLastVars {
		C.PyErr_PrintEx(1)
	} else {
		C.PyErr_PrintEx(0)
	}
}

// PyErr_Print is an alias for PyErr_PrintEx(1).
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_Print
func PyErr_Print() {
	C.PyErr_PrintEx(1)
}

// PyErr_WriteUnraisable calls sys.unraisablehook() using the current exception and obj argument.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_WriteUnraisable
func PyErr_WriteUnraisable(obj *PyObject) {
	C.PyErr_WriteUnraisable(toc(obj))
}

// PyErr_SetString this is the most common way to set the error indicator. The first argument specifies the exception
// type; it is normally one of the standard exceptions, e.g. PyExc_RuntimeError. You need not create a new strong
// reference to it (e.g. with Py_INCREF()). The second argument is an error message; it is decoded from 'utf-8'.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetString
func PyErr_SetString(pyType *PyObject, message string) {
	cmessage := C.CString(message)

	defer C.free(unsafe.Pointer(cmessage))

	C.PyErr_SetString(toc(pyType), cmessage)
}

// PyErr_SetObject is similar to PyErr_SetString() but lets you specify an arbitrary Python object for the "value" of
// the exception.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetObject
func PyErr_SetObject(pyType, value *PyObject) {
	C.PyErr_SetObject(toc(pyType), toc(value))
}

// PyErr_SetNone this function is similar to PyErr_SetString() but lets you specify an arbitrary Python object for the
// "value" of the exception.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetNone
func PyErr_SetNone(pyType *PyObject) {
	C.PyErr_SetNone(toc(pyType))
}

// PyErr_BadArgument is a shorthand for PyErr_SetString(PyExc_TypeError, message), where message indicates that a
// built-in operation was invoked with an illegal argument. It is mostly for internal use.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_BadArgument
func PyErr_BadArgument() {
	C.PyErr_BadArgument()
}

// PyErr_NoMemory is a shorthand for PyErr_SetNone(PyExc_MemoryError); it returns NULL so an object allocation function
// can write return PyErr_NoMemory(); when it runs out of memory.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_NoMemory
func PyErr_NoMemory() *PyObject {
	return togo(C.PyErr_NoMemory())
}

// PyErr_SetImportErrorSubclass is like PyErr_SetImportError() but this function allows for specifying a subclass of
// ImportError to raise.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetImportErrorSubclass
func PyErr_SetImportErrorSubclass(msg, name, path, subclass *PyObject) *PyObject {
	return togo(C.PyErr_SetImportErrorSubclass(toc(msg), toc(name), toc(path), toc(subclass)))
}

// PyErr_SetImportError is a convenience function to raise ImportError. msg will be set as the exception's message
// string. name and path, both of which can be NULL, will be set as the ImportError's respective name and path
// attributes.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetImportError
func PyErr_SetImportError(msg, name, path *PyObject) *PyObject {
	return togo(C.PyErr_SetImportError(toc(msg), toc(name), toc(path)))
}

// PyErr_SyntaxLocationObject sets file, line, and offset information for the current exception. If the current
// exception is not a SyntaxError, then it sets additional attributes, which make the exception printing subsystem
// think the exception is a SyntaxError.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SyntaxLocationObject
func PyErr_SyntaxLocationObject(filename *PyObject, lineno, col_offset int) {
	C.PyErr_SyntaxLocationObject(toc(filename), C.int(lineno), C.int(col_offset))
}

// PyErr_SyntaxLocationEx is like PyErr_SyntaxLocationObject(), but filename is a byte string decoded from the
// filesystem encoding and error handler.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SyntaxLocationEx
func PyErr_SyntaxLocationEx(filename string, lineno, col_offset int) {
	cfilename := C.CString(filename)

	defer C.free(unsafe.Pointer(cfilename))

	C.PyErr_SyntaxLocationEx(cfilename, C.int(lineno), C.int(col_offset))
}

// PyErr_SyntaxLocation is like PyErr_SyntaxLocationEx(), but the col_offset parameter is omitted.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SyntaxLocation
func PyErr_SyntaxLocation(filename string, lineno int) {
	cfilename := C.CString(filename)

	defer C.free(unsafe.Pointer(cfilename))

	C.PyErr_SyntaxLocation(cfilename, C.int(lineno))
}

// PyErr_BadInternalCall is a shorthand for PyErr_SetString(PyExc_SystemError, message), where message indicates that
// an internal operation (e.g. a Python/C API function) was invoked with an illegal argument. It is mostly for internal
// use.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_BadInternalCall
func PyErr_BadInternalCall() {
	C.PyErr_BadInternalCall()
}

// PyErr_Occurred tests whether the error indicator is set. If set, return the exception type (the first argument to
// the last call to one of the PyErr_Set* functions or to PyErr_Restore()). If not set, return NULL. You do not own a
// reference to the return value, so you do not need to Py_DECREF() it.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_Occurred
func PyErr_Occurred() *PyObject {
	return togo(C.PyErr_Occurred())
}

// PyErr_GivenExceptionMatches returns true if the given exception matches the exception type in exc. If exc is a
// class object, this also returns true when given is an instance of a subclass. If exc is a tuple, all exception types
// in the tuple (and recursively in subtuples) are searched for a match.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_GivenExceptionMatches
func PyErr_GivenExceptionMatches(given, exc *PyObject) bool {
	ret := C.PyErr_GivenExceptionMatches(toc(given), toc(exc))

	return ret == 1
}

// PyErr_ExceptionMatches is equivalent to PyErr_GivenExceptionMatches(PyErr_Occurred(), exc). This should only be
// called when an exception is actually set; a memory access violation will occur if no exception has been raised.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_ExceptionMatches
func PyErr_ExceptionMatches(exc *PyObject) bool {
	ret := C.PyErr_ExceptionMatches(toc(exc))

	return ret == 1
}

// PyErr_GetRaisedException Return the exception currently being raised, clearing the error indicator at the same time.
// Return NULL if the error indicator is not set.
//
// This function is used by code that needs to catch exceptions, or code that needs to save and restore the error
// indicator temporarily.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_GetRaisedException
func PyErr_GetRaisedException() *PyObject {
	return togo(C.PyErr_GetRaisedException())
}

// PyErr_SetRaisedException Set exc as the exception currently being raised, clearing the existing exception if one is
// set.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetRaisedException
func PyErr_SetRaisedException(exec *PyObject) {
	C.PyErr_SetRaisedException(toc(exec))
}

// PyErr_GetExcInfo retrieves the old-style representation of the exception info, as known from sys.exc_info(). This
// refers to an exception that was already caught, not to an exception that was freshly raised. Returns new references
// for the three objects, any of which may be NULL. Does not modify the exception info state. This function is kept for
// backwards compatibility.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_GetExcInfo
func PyErr_GetExcInfo() (*PyObject, *PyObject, *PyObject) {
	var pyType, value, traceback *C.PyObject

	C.PyErr_GetExcInfo(&pyType, &value, &traceback) //nolint: gocritic

	return togo(pyType), togo(value), togo(traceback)
}

// PyErr_SetExcInfo sets the exception info, as known from sys.exc_info(). This refers to an exception that was already
// caught, not to an exception that was freshly raised. This function steals the references of the arguments. To clear
// the exception state, pass NULL for all three arguments. This function is kept for backwards compatibility. Prefer
// using PyErr_SetHandledException().
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetExcInfo
func PyErr_SetExcInfo(pyType *PyObject, value *PyObject, traceback *PyObject) {
	C.PyErr_SetExcInfo(toc(pyType), toc(value), toc(traceback))
}

// PyErr_CheckSignals interacts with Python's signal handling.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_CheckSignals
func PyErr_CheckSignals() int {
	return int(C.PyErr_CheckSignals())
}

// PyErr_SetInterrupt simulates the effect of a SIGINT signal arriving. This is equivalent to
// PyErr_SetInterruptEx(SIGINT).
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetInterrupt
func PyErr_SetInterrupt() {
	C.PyErr_SetInterrupt()
}
