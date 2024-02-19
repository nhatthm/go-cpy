package cpy

/*
#include "Python.h"
#include "macro.h"
*/
import "C"

import (
	"unsafe"
)

// Py_EnterRecursiveCall marks a point where a recursive C-level call is about to be performed.
//
// If USE_STACKCHECK is defined, this function checks if the OS stack overflowed using PyOS_CheckStack(). In this is the
// case, it sets a MemoryError and returns a nonzero value.
//
// The function then checks if the recursion limit is reached. If this is the case, a RecursionError is set and a
// nonzero value is returned. Otherwise, zero is returned.
//
// where should be a UTF-8 encoded string such as "in instance check" to be concatenated to the RecursionError message
// caused by the recursion depth limit.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.Py_EnterRecursiveCall
func Py_EnterRecursiveCall(where string) int {
	cwhere := C.CString(where)
	defer C.free(unsafe.Pointer(cwhere))

	return int(C._go_Py_EnterRecursiveCall(cwhere))
}

// Py_LeaveRecursiveCall implements tp_repr for container types requires special recursion handling. In addition to
// protecting the stack, tp_repr also needs to track objects to prevent cycles. The following two functions facilitate
// this functionality. Effectively, these are the C equivalent to reprlib.recursive_repr().
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.Py_LeaveRecursiveCall
func Py_LeaveRecursiveCall() {
	C._go_Py_LeaveRecursiveCall()
}
