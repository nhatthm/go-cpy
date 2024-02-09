package cpy3

/*
#include "Python.h"
*/
import "C"

// PyEval_GetBuiltins returns a dictionary of the builtins in the current execution frame, or the interpreter of the
// thread state if no frame is currently executing.
//
// Reference: https://docs.python.org/3/c-api/reflection.html?highlight=reflection#c.PyEval_GetBuiltins
func PyEval_GetBuiltins() *PyObject {
	return togo(C.PyEval_GetBuiltins())
}

// PyEval_GetLocals returns a dictionary of the local variables in the current execution frame, or NULL if no frame is
// currently executing.
//
// Reference: https://docs.python.org/3/c-api/reflection.html?highlight=reflection#c.PyEval_GetLocals
func PyEval_GetLocals() *PyObject {
	return togo(C.PyEval_GetLocals())
}

// PyEval_GetGlobals returns a dictionary of the global variables in the current execution frame, or NULL if no frame is
// currently executing.
//
// Reference: https://docs.python.org/3/c-api/reflection.html?highlight=reflection#c.PyEval_GetGlobals
func PyEval_GetGlobals() *PyObject {
	return togo(C.PyEval_GetGlobals())
}

// PyEval_GetFuncName returns the name of func if it is a function, class or instance object, else the name of funcs
// type.
//
// Reference: https://docs.python.org/3/c-api/reflection.html?highlight=reflection#c.PyEval_GetFuncName
func PyEval_GetFuncName(pyFunc *PyObject) string {
	return C.GoString(C.PyEval_GetFuncName(toc(pyFunc)))
}

// PyEval_GetFuncDesc returns a description string, depending on the type of func. Return values include "()" for
// functions and methods, " constructor", " instance", and " object". Concatenated with the result of
// PyEval_GetFuncName(), the result will be a description of func.
//
// Reference: https://docs.python.org/3/c-api/reflection.html?highlight=reflection#c.PyEval_GetFuncDesc
func PyEval_GetFuncDesc(pyFunc *PyObject) string {
	return C.GoString(C.PyEval_GetFuncDesc(toc(pyFunc)))
}
