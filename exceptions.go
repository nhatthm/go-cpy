package cpy3

/*
#include "Python.h"
*/
import "C"

import (
	"unsafe"
)

// All standard Python exceptions are available as global variables whose names are PyExc_ followed by the Python
// exception name.
// These have the type PyObject*; they are all class objects.
var (
	PyExc_BaseException          = togo(C.PyExc_BaseException)
	PyExc_Exception              = togo(C.PyExc_Exception)
	PyExc_ArithmeticError        = togo(C.PyExc_ArithmeticError)
	PyExc_AssertionError         = togo(C.PyExc_AssertionError)
	PyExc_AttributeError         = togo(C.PyExc_AttributeError)
	PyExc_BlockingIOError        = togo(C.PyExc_BlockingIOError)
	PyExc_BrokenPipeError        = togo(C.PyExc_BrokenPipeError)
	PyExc_BufferError            = togo(C.PyExc_BufferError)
	PyExc_ChildProcessError      = togo(C.PyExc_ChildProcessError)
	PyExc_ConnectionAbortedError = togo(C.PyExc_ConnectionAbortedError)
	PyExc_ConnectionError        = togo(C.PyExc_ConnectionError)
	PyExc_ConnectionRefusedError = togo(C.PyExc_ConnectionRefusedError)
	PyExc_ConnectcionResetError  = togo(C.PyExc_ConnectionResetError)
	PyExc_EOFError               = togo(C.PyExc_EOFError)
	PyExc_FileExistsError        = togo(C.PyExc_FileExistsError)
	PyExc_FileNotFoundError      = togo(C.PyExc_FileNotFoundError)
	PyExc_FloatingPointError     = togo(C.PyExc_FloatingPointError)
	PyExc_GeneratorExit          = togo(C.PyExc_GeneratorExit)
	PyExc_ImportError            = togo(C.PyExc_ImportError)
	PyExc_IndentationError       = togo(C.PyExc_IndentationError)
	PyExc_IndexError             = togo(C.PyExc_IndexError)
	PyExc_InterruptedError       = togo(C.PyExc_InterruptedError)
	PyExc_IsADirectoryError      = togo(C.PyExc_IsADirectoryError)
	PyExc_KeyError               = togo(C.PyExc_KeyError)
	PyExc_KeyboardInterrupt      = togo(C.PyExc_KeyboardInterrupt)
	PyExc_LookupError            = togo(C.PyExc_LookupError)
	PyExc_MemoryError            = togo(C.PyExc_MemoryError)
	PyExc_ModuleNotFoundError    = togo(C.PyExc_ModuleNotFoundError)
	PyExc_NameError              = togo(C.PyExc_NameError)
	PyExc_NotADirectoryError     = togo(C.PyExc_NotADirectoryError)
	PyExc_NotImplementedError    = togo(C.PyExc_NotImplementedError)
	PyExc_OSError                = togo(C.PyExc_OSError)
	PyExc_OverflowError          = togo(C.PyExc_OverflowError)
	PyExc_PermissionError        = togo(C.PyExc_PermissionError)
	PyExc_ProcessLookupError     = togo(C.PyExc_ProcessLookupError)
	PyExc_RecursionError         = togo(C.PyExc_RecursionError)
	PyExc_ReferenceError         = togo(C.PyExc_ReferenceError)
	PyExc_RuntimeError           = togo(C.PyExc_RuntimeError)
	PyExc_StopAsyncIteration     = togo(C.PyExc_StopAsyncIteration)
	PyExc_StopIteration          = togo(C.PyExc_StopIteration)
	PyExc_SyntaxError            = togo(C.PyExc_SyntaxError)
	PyExc_SystemError            = togo(C.PyExc_SystemError)
	PyExc_SystemExit             = togo(C.PyExc_SystemExit)
	PyExc_TabError               = togo(C.PyExc_TabError)
	PyExc_TimeoutError           = togo(C.PyExc_TimeoutError)
	PyExc_TypeError              = togo(C.PyExc_TypeError)
	PyExc_UnboundLocalError      = togo(C.PyExc_UnboundLocalError)
	PyExc_UnicodeDecodeError     = togo(C.PyExc_UnicodeDecodeError)
	PyExc_UnicodeEncodeError     = togo(C.PyExc_UnicodeEncodeError)
	PyExc_UnicodeError           = togo(C.PyExc_UnicodeError)
	PyExc_UnicodeTranslateError  = togo(C.PyExc_UnicodeTranslateError)
	PyExc_ValueError             = togo(C.PyExc_ValueError)
	PyExc_ZeroDivisionError      = togo(C.PyExc_ZeroDivisionError)
)

// PyErr_NewException creates and returns a new exception class. The name argument must be the name of the new
// exception, a C string of the form module.classname. The base and dict arguments are normally NULL.
// This creates a class object derived from Exception (accessible in C as PyExc_Exception).
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_NewException
func PyErr_NewException(name string, base, dict *PyObject) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PyErr_NewException(cname, toc(base), toc(dict)))
}

// PyErr_NewExceptionWithDoc is the same as PyErr_NewException(), except that the new exception class can easily be
// given a docstring: If doc is non-NULL, it will be used as the docstring for the exception class.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyErr_NewExceptionWithDoc
func PyErr_NewExceptionWithDoc(name, doc string, base, dict *PyObject) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cdoc := C.CString(doc)
	defer C.free(unsafe.Pointer(cdoc))

	return togo(C.PyErr_NewExceptionWithDoc(cname, cdoc, toc(base), toc(dict)))
}

// PyException_GetTraceback returns the traceback associated with the exception as a new reference, as accessible from
// Python through the __traceback__ attribute. If there is no traceback associated, this returns NULL.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyException_GetTraceback
func PyException_GetTraceback(ex *PyObject) *PyObject {
	return togo(C.PyException_GetTraceback(toc(ex)))
}

// PyException_SetTraceback sets the traceback associated with the exception to tb. Use Py_None to clear it.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyException_SetTraceback
func PyException_SetTraceback(ex, tb *PyObject) int {
	return int(C.PyException_SetTraceback(toc(ex), toc(tb)))
}

// PyException_GetContext returns the context (another exception instance during whose handling ex was raised)
// associated with the exception as a new reference, as accessible from Python through the __context__ attribute.
// If there is no context associated, this returns NULL.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyException_GetContext
func PyException_GetContext(ex *PyObject) *PyObject {
	return togo(C.PyException_GetContext(toc(ex)))
}

// PyException_SetContext sets the context associated with the exception to ctx. Use NULL to clear it. There is no
// type check to make sure that ctx is an exception instance. This steals a reference to ctx.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyException_SetContext
func PyException_SetContext(ex, ctx *PyObject) {
	C.PyException_SetContext(toc(ex), toc(ctx))
}

// PyException_GetCause returns the cause (either an exception instance, or None, set by raise ... from ...) associated
// with the exception as a new reference, as accessible from Python through the __cause__ attribute.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyException_GetCause
func PyException_GetCause(ex *PyObject) *PyObject {
	return togo(C.PyException_GetCause(toc(ex)))
}

// PyException_SetCause sets the cause associated with the exception to cause. Use NULL to clear it. There is no type
// check to make sure that cause is either an exception instance or None. This steals a reference to cause.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.PyException_SetCause
func PyException_SetCause(ex, cause *PyObject) {
	C.PyException_SetCause(toc(ex), toc(cause))
}
