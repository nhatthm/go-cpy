package cpy

/*
#include "Python.h"
*/
import "C"

import (
	"unsafe"
)

// PySys_GetObject returns the object name from the sys module or NULL if it does not exist, without setting an exception.
//
// Reference: https://docs.python.org/3/c-api/sys.html#c.PySys_GetObject
func PySys_GetObject(name string) *PyObject {
	cname := C.CString(name)

	defer C.free(unsafe.Pointer(cname))

	return togo(C.PySys_GetObject(cname))
}

// PySys_SetObject sets name in the sys module to v unless v is NULL, in which case name is deleted from the sys module.
// Returns 0 on success, -1 on error.
//
// Reference: https://docs.python.org/3/c-api/sys.html#c.PySys_SetObject
func PySys_SetObject(name string, v *PyObject) int {
	cname := C.CString(name)

	defer C.free(unsafe.Pointer(cname))

	return int(C.PySys_SetObject(cname, toc(v)))
}

// PySys_ResetWarnOptions resets sys.warnoptions to an empty list. This function may be called prior to Py_Initialize().
//
// Reference: https://docs.python.org/3/c-api/sys.html#c.PySys_ResetWarnOptions
func PySys_ResetWarnOptions() {
	C.PySys_ResetWarnOptions()
}

// PySys_GetXOptions returns the current dictionary of -X options, similarly to sys._xoptions. On error, NULL is
// returned and an exception is set.
//
// Reference: https://docs.python.org/3/c-api/sys.html#c.PySys_GetXOptions
func PySys_GetXOptions() *PyObject {
	return togo(C.PySys_GetXOptions())
}
