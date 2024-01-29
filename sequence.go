package python3

/*
#include "Python.h"
*/
import "C"

// PySequence_Contains: https://docs.python.org/3/c-api/sequence.html#c.PySequence_Contains
func PySequence_Contains(o *PyObject, value *PyObject) int {
	return int(C.PySequence_Contains(toc(o), toc(value)))
}
