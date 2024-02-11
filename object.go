package cpy3

//go:generate go run resources/scripts/variadic.go

/*
#include "Python.h"
#include "macro.h"
#include "variadic.h"
*/
import "C"

import (
	"unsafe"
)

// MaxVariadicLength is the maximum number of arguments that can be passed to a variadic C function due to a cgo limitation
const MaxVariadicLength = 20

// Constants used for comparison in PyObject_RichCompareBool.
var (
	Py_LT = int(C.Py_LT)
	Py_LE = int(C.Py_LE)
	Py_EQ = int(C.Py_EQ)
	Py_NE = int(C.Py_NE)
	Py_GT = int(C.Py_GT)
	Py_GE = int(C.Py_GE)
)

// Py_None is the Python None object, denoting lack of value. This object has no methods and is immortal.
//
// Reference: https://docs.python.org/3/c-api/none.html#c.Py_None
var Py_None = togo(C.Py_None)

// PyObject is the base of all object types. This is a type which contains the information Python needs to treat a
// pointer to an object as an object. In a normal "release" build, it contains only the object's reference count and a
// pointer to the corresponding type object. Nothing is actually declared to be a PyObject, but every pointer to a
// Python object can be cast to a PyObject*. Access to the members must be done by using the macros Py_REFCNT and
// Py_TYPE.
//
// Reference: https://docs.python.org/3/c-api/structures.html?highlight=pyobject#c.PyObject
type PyObject C.PyObject

// IncRef indicates taking a new strong reference to object o, indicating it is in use and should not be destroyed.
//
// Reference: https://docs.python.org/3/c-api/refcounting.html#c.Py_INCREF
func (pyObject *PyObject) IncRef() {
	C.Py_IncRef(toc(pyObject))
}

// DecRef releases a strong reference to object o, indicating the reference is no longer used.
//
// Reference: https://docs.python.org/3/c-api/refcounting.html#c.Py_DECREF
func (pyObject *PyObject) DecRef() {
	C.Py_DecRef(toc(pyObject))
}

// ReprEnter is called at the beginning of the tp_repr implementation to detect cycles.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.Py_ReprEnter
func (pyObject *PyObject) ReprEnter() int {
	return int(C.Py_ReprEnter(toc(pyObject)))
}

// ReprLeave ends a PyObject.ReprEnter(). Must be called once for each invocation of Py_ReprEnter() that returns zero.
//
// Reference: https://docs.python.org/3/c-api/exceptions.html#c.Py_ReprLeave
func (pyObject *PyObject) ReprLeave() {
	C.Py_ReprLeave(toc(pyObject))
}

// HasAttr returns 1 if o has the attribute attr_name, and 0 otherwise. This is equivalent to the Python expression
// hasattr(o, attr_name). This function always succeeds.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_HasAttr
func (pyObject *PyObject) HasAttr(attr_name *PyObject) bool {
	return C.PyObject_HasAttr(toc(pyObject), toc(attr_name)) == 1
}

// HasAttrString is the same as PyObject.HasAttr(), but attr_name is specified as a const char* UTF-8 encoded bytes
// string, rather than a PyObject*.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_HasAttrString
func (pyObject *PyObject) HasAttrString(attr_name string) bool {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return C.PyObject_HasAttrString(toc(pyObject), cattr_name) == 1
}

// GetAttr retrieves an attribute named attr_name from object o. Returns the attribute value on success, or NULL on
// failure. This is the equivalent of the Python expression o.attr_name.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_GetAttr
func (pyObject *PyObject) GetAttr(attr_name *PyObject) *PyObject {
	return togo(C.PyObject_GetAttr(toc(pyObject), toc(attr_name)))
}

// GetAttrString is the same as PyObject.GetAttr(), but attr_name is specified as a const char* UTF-8 encoded bytes
// string, rather than a PyObject*.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_GetAttrString
func (pyObject *PyObject) GetAttrString(attr_name string) *PyObject {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return togo(C.PyObject_GetAttrString(toc(pyObject), cattr_name))
}

// SetAttr sets the value of the attribute named attr_name, for object o, to the value v. Raise an exception and return
// -1 on failure; return 0 on success. This is the equivalent of the Python statement o.attr_name = v.
//
// If v is NULL, the attribute is deleted. This behavior is deprecated in favor of using PyObject_DelAttr(), but there
// are currently no plans to remove it.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_SetAttr
func (pyObject *PyObject) SetAttr(attr_name *PyObject, v *PyObject) int {
	return int(C.PyObject_SetAttr(toc(pyObject), toc(attr_name), toc(v)))
}

// SetAttrString is the same as PyObject.SetAttr(), but attr_name is specified as a const char* UTF-8 encoded bytes
// string, rather than a PyObject*.
//
// If v is NULL, the attribute is deleted, but this feature is deprecated in favor of using PyObject.DelAttrString().
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_SetAttrString
func (pyObject *PyObject) SetAttrString(attr_name string, v *PyObject) int {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return int(C.PyObject_SetAttrString(toc(pyObject), cattr_name, toc(v)))
}

// DelAttr deletes attribute named attr_name, for object o. Returns -1 on failure. This is the equivalent of the Python
// statement del o.attr_name.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_DelAttr
func (pyObject *PyObject) DelAttr(attr_name *PyObject) int {
	return int(C._go_PyObject_DelAttr(toc(pyObject), toc(attr_name)))
}

// DelAttrString is the same as PyObject.DelAttr(), but attr_name is specified as a const char* UTF-8 encoded bytes
// string, rather than a PyObject*.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_DelAttrString
func (pyObject *PyObject) DelAttrString(attr_name string) int {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return int(C._go_PyObject_DelAttrString(toc(pyObject), cattr_name))
}

// RichCompare compares the values of o1 and o2 using the operation specified by opid, which must be one of Py_LT,
// Py_LE, Py_EQ, Py_NE, Py_GT, or Py_GE, corresponding to <, <=, ==, !=, >, or >= respectively. This is the equivalent
// of the Python expression o1 op o2, where op is the operator corresponding to opid. Returns the value of the
// comparison on success, or NULL on failure.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_RichCompare
func (pyObject *PyObject) RichCompare(o *PyObject, opid int) *PyObject {
	return togo(C.PyObject_RichCompare(toc(pyObject), toc(o), C.int(opid)))
}

// RichCompareBool compares the values of o1 and o2 using the operation specified by opid, like PyObject_RichCompare(),
// but returns -1 on error, 0 if the result is false, 1 otherwise.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_RichCompareBool
func (pyObject *PyObject) RichCompareBool(o *PyObject, opid int) int {
	return int(C.PyObject_RichCompareBool(toc(pyObject), toc(o), C.int(opid)))
}

// Repr compute a string representation of object o. Returns the string representation on success, NULL on failure.
// This is the equivalent of the Python expression repr(o). Called by the repr() built-in function.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_Repr
func (pyObject *PyObject) Repr() *PyObject {
	return togo(C.PyObject_Repr(toc(pyObject)))
}

// ASCII is similar to PyObject.Repr(), computes a string representation of object o, but escapes the non-ASCII
// characters in the string returned by PyObject_Repr() with \x, \u or \U escapes. This generates a string similar to
// that returned by PyObject.Repr() in Python 2. Called by the ascii() built-in function.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_ASCII
func (pyObject *PyObject) ASCII() *PyObject {
	return togo(C.PyObject_ASCII(toc(pyObject)))
}

// Str computes a string representation of object o. Returns the string representation on success, NULL on failure.
// This is the equivalent of the Python expression str(o). Called by the str() built-in function and, therefore,
// by the print() function.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_Str
func (pyObject *PyObject) Str() *PyObject {
	return togo(C.PyObject_Str(toc(pyObject)))
}

// Bytes computes a bytes representation of object o. NULL is returned on failure and a bytes object on success.
// This is equivalent to the Python expression bytes(o), when o is not an integer. Unlike bytes(o), a TypeError is
// raised when o is an integer instead of a zero-initialized bytes object.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_Bytes
func (pyObject *PyObject) Bytes() *PyObject {
	return togo(C.PyObject_Bytes(toc(pyObject)))
}

// IsSubclass returns 1 if the class derived is identical to or derived from the class cls, otherwise return 0.
// In case of an error, return -1.
//
// If cls is a tuple, the check will be done against every entry in cls. The result will be 1 when at least one of the
// checks returns 1, otherwise it will be 0.
//
// If cls has a __subclasscheck__() method, it will be called to determine the subclass status as described in PEP 3119.
// Otherwise, derived is a subclass of cls if it is a direct or indirect subclass, i.e. contained in cls.__mro__.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_IsSubclass
func (pyObject *PyObject) IsSubclass(cls *PyObject) int {
	return int(C.PyObject_IsSubclass(toc(pyObject), toc(cls)))
}

// IsInstance returns 1 if inst is an instance of the class cls or a subclass of cls, or 0 if not. On error, returns -1
// and sets an exception.
//
// If cls is a tuple, the check will be done against every entry in cls. The result will be 1 when at least one of the
// checks returns 1, otherwise it will be 0.
//
// If cls has a __instancecheck__() method, it will be called to determine the subclass status as described in PEP 3119.
// Otherwise, inst is an instance of cls if its class is a subclass of cls.
//
// An instance inst can override what is considered its class by having a __class__ attribute.
//
// An object cls can override if it is considered a class, and what its base classes are, by having a __bases__
// attribute (which must be a tuple of base classes).
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_IsInstance
func (pyObject *PyObject) IsInstance(cls *PyObject) int {
	return int(C.PyObject_IsInstance(toc(pyObject), toc(cls)))
}

// PyCallable_Check returns 1 if inst is an instance of the class cls or a subclass of cls, or 0 if not. On error,
// returns -1 and sets an exception.
//
// If cls is a tuple, the check will be done against every entry in cls. The result will be 1 when at least one of the
// checks returns 1, otherwise it will be 0.
//
// If cls has a __instancecheck__() method, it will be called to determine the subclass status as described in PEP 3119.
// Otherwise, inst is an instance of cls if its class is a subclass of cls.
//
// An instance inst can override what is considered its class by having a __class__ attribute.
//
// An object cls can override if it is considered a class, and what its base classes are, by having a __bases__
// attribute (which must be a tuple of base classes).
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyCallable_Check
func PyCallable_Check(o *PyObject) bool {
	return C.PyCallable_Check(toc(o)) == 1
}

// Call calls a callable Python object callable, with arguments given by the tuple args, and named arguments given by
// the dictionary kwargs. The function returns the result of the call on success, or raise an exception and return NULL
// on failure.
//
// args must not be NULL; use an empty tuple if no arguments are needed. If no named arguments are needed, kwargs
// can be NULL.
//
// This is the equivalent of the Python expression: callable(*args).
//
// Reference: https://docs.python.org/3/c-api/call.html#c.PyObject_Call
func (pyObject *PyObject) Call(args *PyObject, kwargs *PyObject) *PyObject {
	return togo(C.PyObject_Call(toc(pyObject), toc(args), toc(kwargs)))
}

// CallObject calls a callable Python object callable, with a variable number of C arguments. The C arguments are
// described using a Py_BuildValue() style format string. The format can be NULL, indicating that no arguments are
// provided. The function returns the result of the call on success, or raise an exception and return NULL on failure.
//
// This is the equivalent of the Python expression: callable(*args).
//
// Note that if you only pass PyObject* args, PyObject_CallFunctionObjArgs() is a faster alternative.
//
// Reference: https://docs.python.org/3/c-api/call.html#c.PyObject_CallObject
func (pyObject *PyObject) CallObject(args *PyObject) *PyObject {
	return togo(C.PyObject_CallObject(toc(pyObject), toc(args)))
}

// CallFunctionObjArgs calls a callable Python object callable, with a variable number of PyObject* arguments. The
// arguments are provided as a variable number of parameters followed by NULL. The function returns the result of the
// call on success, or raise an exception and return NULL on failure.
//
// This is the equivalent of the Python expression: callable(arg1, arg2, ...).
//
// Reference: https://docs.python.org/3/c-api/call.html#c.PyObject_CallFunctionObjArgs
func (pyObject *PyObject) CallFunctionObjArgs(args ...*PyObject) *PyObject {
	if len(args) > MaxVariadicLength {
		panic("CallFunctionObjArgs: too many arrguments")
	}

	if len(args) == 0 {
		return togo(C._go_PyObject_CallFunctionObjArgs(toc(pyObject), 0, (**C.PyObject)(nil)))
	}

	cargs := make([]*C.PyObject, len(args))

	for i, arg := range args {
		cargs[i] = toc(arg)
	}

	return togo(C._go_PyObject_CallFunctionObjArgs(toc(pyObject), C.int(len(args)), (**C.PyObject)(unsafe.Pointer(&cargs[0]))))
}

// CallMethodObjArgs calls a method of the Python object obj, where the name of the method is given as a Python string
// object in name. It is called with a variable number of PyObject* arguments. The arguments are provided as a variable
// number of parameters followed by NULL. The function returns the result of the call on success, or raise an exception
// and return NULL on failure.
//
// Reference: https://docs.python.org/3/c-api/call.html#c.PyObject_CallMethodObjArgs
func (pyObject *PyObject) CallMethodObjArgs(name *PyObject, args ...*PyObject) *PyObject {
	if len(args) > MaxVariadicLength {
		panic("CallMethodObjArgs: too many arguments")
	}

	if len(args) == 0 {
		return togo(C._go_PyObject_CallMethodObjArgs(toc(pyObject), toc(name), 0, (**C.PyObject)(nil)))
	}

	cargs := make([]*C.PyObject, len(args))

	for i, arg := range args {
		cargs[i] = toc(arg)
	}

	return togo(C._go_PyObject_CallMethodObjArgs(toc(pyObject), toc(name), C.int(len(args)), (**C.PyObject)(unsafe.Pointer(&cargs[0]))))
}

// CallMethodArgs is the same as PyObject.CallMethodObjArgs but with name as go string
func (pyObject *PyObject) CallMethodArgs(name string, args ...*PyObject) *PyObject {
	pyName := PyUnicode_FromString(name)
	defer pyName.DecRef()

	return pyObject.CallMethodObjArgs(pyName, args...)
}

// Hash computes and returns the hash value of an object o. On failure, return -1. This is the equivalent of the Python
// expression hash(o).
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_Hash
func (pyObject *PyObject) Hash() int {
	return int(C.PyObject_Hash(toc(pyObject)))
}

// HashNotImplemented sets a TypeError indicating that type(o) is not hashable and return -1. This function receives
// special treatment when stored in a tp_hash slot, allowing a type to explicitly indicate to the interpreter that it is
// not hashable.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_HashNotImplemented
func (pyObject *PyObject) HashNotImplemented() int {
	return int(C.PyObject_HashNotImplemented(toc(pyObject)))
}

// IsTrue returns 1 if the object o is considered to be true, and 0 otherwise. This is equivalent to the Python
// expression `not not o`. On failure, return -1.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_IsTrue
// nolint: dupword
func (pyObject *PyObject) IsTrue() int {
	return int(C.PyObject_IsTrue(toc(pyObject)))
}

// Not returns 0 if the object o is considered to be true, and 1 otherwise. This is equivalent to the Python expression
// not o. On failure, return -1.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_Not
func (pyObject *PyObject) Not() int {
	return int(C.PyObject_Not(toc(pyObject)))
}

// Type returns a type object corresponding to the object type of object o when o is non-NULL. On failure, raises
// SystemError and returns NULL. This is equivalent to the Python expression type(o). This function creates a new strong
// reference to the return value. There's really no reason to use this function instead of the Py_TYPE() function, which
// returns a pointer of type PyTypeObject*, except when a new strong reference is needed.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_Type
func (pyObject *PyObject) Type() *PyObject {
	return togo(C.PyObject_Type(toc(pyObject)))
}

// Length returns the length of object o. If the object o provides either the sequence and mapping protocols, the
// sequence length is returned. On error, -1 is returned. This is the equivalent to the Python expression len(o).
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_Length
func (pyObject *PyObject) Length() int {
	return int(C.PyObject_Length(toc(pyObject)))
}

// LengthHint return an estimated length for the object o. First try to return its actual length, then an estimate using
// __length_hint__(), and finally return the default value. On error return -1. This is the equivalent to the Python
// expression operator.length_hint(o, defaultvalue).
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_LengthHint
func (pyObject *PyObject) LengthHint(pyDefault int) int {
	return int(C.PyObject_LengthHint(toc(pyObject), C.Py_ssize_t(pyDefault)))
}

// GetItem returns element of o corresponding to the object key or NULL on failure. This is the equivalent of the Python
// expression o[key].
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_GetItem
func (pyObject *PyObject) GetItem(key *PyObject) *PyObject {
	return togo(C.PyObject_GetItem(toc(pyObject), toc(key)))
}

// SetItem maps the object key to the value v. Raise an exception and return -1 on failure; return 0 on success. This is
// the equivalent of the Python statement o[key] = v. This function does not steal a reference to v.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_SetItem
func (pyObject *PyObject) SetItem(key, v *PyObject) int {
	return int(C.PyObject_SetItem(toc(pyObject), toc(key), toc(v)))
}

// DelItem removes the mapping for the object key from the object o. Return -1 on failure. This is equivalent to the
// Python statement del o[key].
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_DelItem
func (pyObject *PyObject) DelItem(key *PyObject) int {
	return int(C.PyObject_DelItem(toc(pyObject), toc(key)))
}

// Dir is equivalent to the Python expression dir(o), returning a (possibly empty) list of strings appropriate for the
// object argument, or NULL if there was an error. If the argument is NULL, this is like the Python dir(), returning the
// names of the current locals; in this case, if no execution frame is active then NULL is returned but PyErr_Occurred()
// will return false.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_Dir
func (pyObject *PyObject) Dir() *PyObject {
	return togo(C.PyObject_Dir(toc(pyObject)))
}

// GetIter is equivalent to the Python expression iter(o). It returns a new iterator for the object argument, or the
// object itself if the object is already an iterator. Raises TypeError and returns NULL if the object cannot be
// iterated.
//
// Reference: https://docs.python.org/3/c-api/object.html#c.PyObject_GetIter
func (pyObject *PyObject) GetIter() *PyObject {
	return togo(C.PyObject_GetIter(toc(pyObject)))
}
