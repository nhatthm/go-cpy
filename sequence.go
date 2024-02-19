package cpy

/*
#include "Python.h"
*/
import "C"

// PySequence_Contains determines if o contains value. If an item in o is equal to value, return 1, otherwise return 0.
// On error, return -1. This is equivalent to the Python expression value in o.
//
// Reference: https://docs.python.org/3/c-api/sequence.html#c.PySequence_Contains
func PySequence_Contains(o *PyObject, value *PyObject) int {
	return int(C.PySequence_Contains(toc(o), toc(value)))
}
