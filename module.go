package cpy

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import "unsafe"

// Module is an instance of PyTypeObject represents the Python module type. This is exposed to Python programs as
// types.ModuleType.
//
// Reference: https://docs.python.org/3/c-api/module.html#c.PyModule_Type
var Module = togo((*C.PyObject)(unsafe.Pointer(&C.PyModule_Type)))

// PyModule_Check returns true if p is a module object, or a subtype of a module object. This function always succeeds.
//
// Reference: https://docs.python.org/3/c-api/module.html#c.PyModule_Check
func PyModule_Check(p *PyObject) bool {
	return C._go_PyModule_Check(toc(p)) != 0
}

// PyModule_CheckExact returns true if p is a module object, but not a subtype of PyModule_Type. This function always
// succeeds.
//
// Reference: https://docs.python.org/3/c-api/module.html#c.PyModule_CheckExact
func PyModule_CheckExact(p *PyObject) bool {
	return C._go_PyModule_CheckExact(toc(p)) != 0
}

// PyModule_NewObject returns a new module object with the __name__ attribute set to name. The module's __name__,
// __doc__, __package__, and __loader__ attributes are filled in (all but __name__ are set to None); the caller is
// responsible for providing a __file__ attribute.
//
// Reference: https://docs.python.org/3/c-api/module.html#c.PyModule_NewObject
func PyModule_NewObject(name *PyObject) *PyObject {
	return togo(C.PyModule_NewObject(toc(name)))
}

// PyModule_New is similar to PyModule_NewObject(), but the name is a UTF-8 encoded string instead of a Unicode object.
//
// Reference: https://docs.python.org/3/c-api/module.html#c.PyModule_New
func PyModule_New(name string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PyModule_New(cname))
}

// PyModule_GetDict returns the dictionary object that implements module's namespace; this object is the same as the
// __dict__ attribute of the module object. If module is not a module object (or a subtype of a module object),
// SystemError is raised and NULL is returned.
//
// Reference: https://docs.python.org/3/c-api/module.html#c.PyModule_GetDict
func PyModule_GetDict(module *PyObject) *PyObject {
	return togo(C.PyModule_GetDict(toc(module)))
}

// PyModule_GetNameObject return module's __name__ value. If the module does not provide one, or if it is not a string,
// SystemError is raised and NULL is returned.
//
// Reference: https://docs.python.org/3/c-api/module.html#c.PyModule_GetNameObject
func PyModule_GetNameObject(module *PyObject) *PyObject {
	return togo(C.PyModule_GetNameObject(toc(module)))
}

// PyModule_GetName is similar to PyModule_GetNameObject() but return the name encoded to 'utf-8'.
//
// Reference: https://docs.python.org/3/c-api/module.html#c.PyModule_GetName
func PyModule_GetName(module *PyObject) string {
	cname := C.PyModule_GetName(toc(module))
	return C.GoString(cname)
}

// PyModule_GetState returns the "state" of the module, that is, a pointer to the block of memory allocated at module
// creation time, or NULL. See PyModuleDef.m_size.
//
// Reference: https://docs.python.org/3/c-api/module.html#c.PyModule_GetState
func PyModule_GetState(module *PyObject) unsafe.Pointer {
	return unsafe.Pointer(C.PyModule_GetState(toc(module))) //nolint: unconvert
}

// PyModule_GetFilenameObject returns the name of the file from which module was loaded using module's __file__
// attribute. If this is not defined, or if it is not a unicode string, raise SystemError and return NULL; otherwise
// return a reference to a Unicode object.
//
// Reference: https://docs.python.org/3/c-api/module.html#c.PyModule_GetFilenameObject
func PyModule_GetFilenameObject(module *PyObject) *PyObject {
	return togo(C.PyModule_GetFilenameObject(toc(module)))
}
