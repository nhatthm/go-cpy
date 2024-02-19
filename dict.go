package cpy

/*
#include "Python.h"
#include "macro.h"
*/
import "C"

import (
	"unsafe"
)

// Dict is an instance of PyTypeObject represents the Python dictionary type. This is the same object as `dict` in the
// Python layer.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_Type
var Dict = togo((*C.PyObject)(unsafe.Pointer(&C.PyDict_Type)))

// PyDict_Check returns true if p is a dict object or an instance of a subtype of the dict type. This function always
// succeeds.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_Check
func PyDict_Check(p *PyObject) bool {
	return C._go_PyDict_Check(toc(p)) != 0
}

// PyDict_CheckExact returns true if p is a dict object, but not an instance of a subtype of the dict type. This
// function always succeeds.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_CheckExact
func PyDict_CheckExact(p *PyObject) bool {
	return C._go_PyDict_CheckExact(toc(p)) != 0
}

// PyDict_New returns a new empty dictionary, or NULL on failure.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_New
func PyDict_New() *PyObject {
	return togo(C.PyDict_New())
}

// PyDictProxy_New returns a types.MappingProxyType object for a mapping which enforces read-only behavior. This is
// normally used to create a view to prevent modification of the dictionary for non-dynamic class types.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDictProxy_New
func PyDictProxy_New(mapping *PyObject) *PyObject {
	return togo(C.PyDictProxy_New(toc(mapping)))
}

// PyDict_Clear empties an existing dictionary of all key-value pairs.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_Clear
func PyDict_Clear(p *PyObject) {
	C.PyDict_Clear(toc(p))
}

// PyDict_Contains determines if dictionary p contains key. If an item in p is matches key, return 1, otherwise return
// 0. On error, return -1. This is equivalent to the Python expression key in p.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_Contains
func PyDict_Contains(p, key *PyObject) int {
	return int(C.PyDict_Contains(toc(p), toc(key)))
}

// PyDict_Copy returns a new dictionary that contains the same key-value pairs as p.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_Copy
func PyDict_Copy(p *PyObject) *PyObject {
	return togo(C.PyDict_Copy(toc(p)))
}

// PyDict_SetItem : https://docs.python.org/3/c-api/dict.html#c.PyDict_SetItem
func PyDict_SetItem(p, key, val *PyObject) int {
	return int(C.PyDict_SetItem(toc(p), toc(key), toc(val)))
}

// PyDict_SetItemString inserts val into the dictionary p with a key of key. key must be hashable; if it isn't,
// TypeError will be raised. Return 0 on success or -1 on failure. This function does not steal a reference to val.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_SetItemString
func PyDict_SetItemString(p *PyObject, key string, val *PyObject) int {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	return int(C.PyDict_SetItemString(toc(p), ckey, toc(val)))
}

// PyDict_DelItem removes the entry in dictionary p with the key. key must be hashable; if it isn't, TypeError is
// raised. If key is not in the dictionary, KeyError is raised. Return 0 on success or -1 on failure.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_DelItem
func PyDict_DelItem(p, key *PyObject) int {
	return int(C.PyDict_DelItem(toc(p), toc(key)))
}

// PyDict_DelItemString is the same as PyDict_DelItem(), but key is specified as a const char* UTF-8 encoded bytes
// string, rather than a PyObject*.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_DelItemString
func PyDict_DelItemString(p *PyObject, key string) int {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	return int(C.PyDict_DelItemString(toc(p), ckey))
}

// PyDict_GetItem returns the object from dictionary p which has the key. Return NULL if the key is not present, but
// without setting an exception.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_GetItem
func PyDict_GetItem(p, key *PyObject) *PyObject {
	return togo(C.PyDict_GetItem(toc(p), toc(key)))
}

// PyDict_GetItemWithError is a variant of PyDict_GetItem() that does not suppress exceptions. Return NULL with an
// exception set if an exception occurred. Return NULL without an exception set if the key wasn't present.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_GetItemWithError
func PyDict_GetItemWithError(p, key *PyObject) *PyObject {
	return togo(C.PyDict_GetItemWithError(toc(p), toc(key)))
}

// PyDict_GetItemString is the same as PyDict_GetItem(), but key is specified as a const char* UTF-8 encoded bytes
// string, rather than a PyObject*.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_GetItemString
func PyDict_GetItemString(p *PyObject, key string) *PyObject {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	return togo(C.PyDict_GetItemString(toc(p), ckey))
}

// PyDict_SetDefault is the same as the Python-level dict.setdefault(). If present, it returns the value corresponding
// to key from the dictionary p. If the key is not in the dict, it is inserted with value defaultobj and defaultobj
// is returned. This function evaluates the hash function of key only once, instead of evaluating it independently for
// the lookup and the insertion.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_SetDefault
func PyDict_SetDefault(p, key, pyDefault *PyObject) *PyObject {
	return togo(C.PyDict_SetDefault(toc(p), toc(key), toc(pyDefault)))
}

// PyDict_Items returns a PyListObject containing all the items from the dictionary.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_Items
func PyDict_Items(p *PyObject) *PyObject {
	return togo(C.PyDict_Items(toc(p)))
}

// PyDict_Keys returns a PyListObject containing all the items from the dictionary.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_Keys
func PyDict_Keys(p *PyObject) *PyObject {
	return togo(C.PyDict_Keys(toc(p)))
}

// PyDict_Values returns a PyListObject containing all the values from the dictionary p.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_Values
func PyDict_Values(p *PyObject) *PyObject {
	return togo(C.PyDict_Values(toc(p)))
}

// PyDict_Size returns the number of items in the dictionary. This is equivalent to len(p) on a dictionary.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_Size
func PyDict_Size(p *PyObject) int {
	return int(C.PyDict_Size(toc(p)))
}

// PyDict_Next iterates over all key-value pairs in the dictionary p. The Py_ssize_t referred to by ppos must be
// initialized to 0 prior to the first call to this function to start the iteration; the function returns true for each
// pair in the dictionary, and false once all pairs have been reported. The parameters pkey and pvalue should either
// point to PyObject* variables that will be filled in with each key and value, respectively, or may be NULL.
// Any references returned through them are borrowed. ppos should not be altered during iteration. Its value represents
// offsets within the internal dictionary structure, and since the structure is sparse, the offsets are not consecutive.
//
// Reference: https://docs.python.org/3/c-api/dict.html#c.PyDict_Next
func PyDict_Next(p *PyObject, ppos *int, pkey, pvalue **PyObject) bool {
	cpos := C.Py_ssize_t(*ppos)
	ckey := toc(*pkey)
	cvalue := toc(*pvalue)

	res := C.PyDict_Next(toc(p), &cpos, &ckey, &cvalue) != 0 //nolint: gocritic

	*ppos = int(cpos)
	*pkey = togo(ckey)
	*pvalue = togo(cvalue)

	return res
}
