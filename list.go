package cpy

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import "unsafe"

// List is an instance of PyTypeObject represents the Python list type. This is the same object as list in the Python
// layer.
//
// Reference: https://docs.python.org/3/c-api/list.html#c.PyList_Type
var List = togo((*C.PyObject)(unsafe.Pointer(&C.PyList_Type)))

// PyList_Check returns true if p is a list object or an instance of a subtype of the list type. This function always
// succeeds.
//
// Reference: https://docs.python.org/3/c-api/list.html#c.PyList_Check
func PyList_Check(p *PyObject) bool {
	return C._go_PyList_Check(toc(p)) != 0
}

// PyList_CheckExact returns true if p is a list object or an instance of a subtype of the list type. This function
// always succeeds.
//
// Reference: https://docs.python.org/3/c-api/list.html#c.PyList_CheckExact
func PyList_CheckExact(p *PyObject) bool {
	return C._go_PyList_CheckExact(toc(p)) != 0
}

// PyList_New returns a new list of length len on success, or NULL on failure.
//
// Reference: https://docs.python.org/3/c-api/list.html#c.PyList_New
func PyList_New(length int) *PyObject {
	return togo(C.PyList_New(C.Py_ssize_t(length)))
}

// PyList_Size returns the length of the list object in list; this is equivalent to len(list) on a list object.
//
// Reference: https://docs.python.org/3/c-api/list.html#c.PyList_Size
func PyList_Size(p *PyObject) int {
	return int(C.PyList_Size(toc(p)))
}

// PyList_GetItem returns the object at position index in the list pointed to by list. The position must be
// non-negative; indexing from the end of the list is not supported. If index is out of bounds (<0 or >=len(list)),
// return NULL and set an IndexError exception.
//
// Reference: https://docs.python.org/3/c-api/list.html#c.PyList_GetItem
func PyList_GetItem(p *PyObject, pos int) *PyObject {
	return togo(C.PyList_GetItem(toc(p), C.Py_ssize_t(pos)))
}

// PyList_SetItem sets the item at an index in list to item. Return 0 on success. If index is out of bounds, return -1
// and set an IndexError exception.
//
// Reference: https://docs.python.org/3/c-api/list.html#c.PyList_SetItem
func PyList_SetItem(p *PyObject, pos int, o *PyObject) int {
	return int(C.PyList_SetItem(toc(p), C.Py_ssize_t(pos), toc(o)))
}

// PyList_Insert inserts the item into the list in front of the index. Return 0 if successful; return -1 and set an
// exception if unsuccessful. Analogous to list.insert(index, item).
//
// Reference: https://docs.python.org/3/c-api/list.html#c.PyList_Insert
func PyList_Insert(p *PyObject, index int, item *PyObject) int {
	return int(C.PyList_Insert(toc(p), C.Py_ssize_t(index), toc(item)))
}

// PyList_Append appends the object item at the end of the list. Return 0 if successful; return -1 and set an exception
// if unsuccessful. Analogous to list.append(item).
//
// Reference: https://docs.python.org/3/c-api/list.html#c.PyList_Append
func PyList_Append(p, item *PyObject) int {
	return int(C.PyList_Append(toc(p), toc(item)))
}

// PyList_GetSlice appends the object item at the end of the list. Return 0 if successful; return -1 and set an
// exception if unsuccessful. Analogous to list.append(item).
//
// Reference: https://docs.python.org/3/c-api/list.html#c.PyList_GetSlice
func PyList_GetSlice(p *PyObject, low, high int) *PyObject {
	return togo(C.PyList_GetSlice(toc(p), C.Py_ssize_t(low), C.Py_ssize_t(high)))
}

// PyList_SetSlice sets the slice of list between low and high to the contents of itemlist. Analogous to
// list[low:high] = itemlist. The itemlist may be NULL, indicating the assignment of an empty list (slice deletion).
// Return 0 on success, -1 on failure. Indexing from the end of the list is not supported.
//
// Reference: https://docs.python.org/3/c-api/list.html#c.PyList_SetSlice
func PyList_SetSlice(p *PyObject, low, high int, itemlist *PyObject) int {
	return int(C.PyList_SetSlice(toc(p), C.Py_ssize_t(low), C.Py_ssize_t(high), toc(itemlist)))
}

// PyList_Sort sorts the items of list in place. Return 0 on success, -1 on failure. This is equivalent to list.sort().
//
// Reference: https://docs.python.org/3/c-api/list.html#c.PyList_Sort
func PyList_Sort(list *PyObject) int {
	return int(C.PyList_Sort(toc(list)))
}

// PyList_Reverse sorts the items of list in place. Return 0 on success, -1 on failure. This is equivalent to
// list.sort().
//
// Reference: https://docs.python.org/3/c-api/list.html#c.PyList_Reverse
func PyList_Reverse(list *PyObject) int {
	return int(C.PyList_Reverse(toc(list)))
}

// PyList_AsTuple returns a new tuple object containing the contents of list; equivalent to tuple(list).
//
// Reference: https://docs.python.org/3/c-api/list.html#c.PyList_AsTuple
func PyList_AsTuple(list *PyObject) *PyObject {
	return togo(C.PyList_AsTuple(toc(list)))
}
