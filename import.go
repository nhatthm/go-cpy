package cpy3

/*
#include "Python.h"
*/
import "C"

import (
	"unsafe"
)

// PyImport_ImportModule is a wrapper around PyImport_Import() which takes a const char* as an argument instead of
// a PyObject*.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_ImportModule
func PyImport_ImportModule(name string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PyImport_ImportModule(cname))
}

// PyImport_ImportModuleEx imports a module. This is best described by referring to the built-in Python
// function __import__().
//
// The return value is a new reference to the imported module or top-level package, or NULL with an exception set on
// failure. Like for __import__(), the return value when a submodule of a package was requested is normally the
// top-level package, unless a non-empty fromlist was given.
//
// Failing imports remove incomplete module objects, like with PyImport_ImportModule().
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_ImportModuleEx
func PyImport_ImportModuleEx(name string, globals, locals, fromlist *PyObject) *PyObject {
	return PyImport_ImportModuleLevel(name, globals, locals, fromlist, 0)
}

// PyImport_ImportModuleLevelObject imports a module. This is best described by referring to the built-in Python
// function __import__(), as the standard __import__() function calls this function directly.
//
// The return value is a new reference to the imported module or top-level package, or NULL with an exception set on
// failure. Like for __import__(), the return value when a submodule of a package was requested is normally the
// top-level package, unless a non-empty fromlist was given.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_ImportModuleLevelObject
func PyImport_ImportModuleLevelObject(name, globals, locals, fromlist *PyObject, level int) *PyObject {
	return togo(C.PyImport_ImportModuleLevelObject(toc(name), toc(globals), toc(locals), toc(fromlist), C.int(level)))
}

// PyImport_ImportModuleLevel is similar to PyImport_ImportModuleLevelObject(), but the name is a UTF-8 encoded string
// instead of a Unicode object.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_ImportModuleLevel
func PyImport_ImportModuleLevel(name string, globals, locals, fromlist *PyObject, level int) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PyImport_ImportModuleLevel(cname, toc(globals), toc(locals), toc(fromlist), C.int(level)))
}

// PyImport_Import is a higher-level interface that calls the current "import hook function" (with an explicit level of
// 0, meaning absolute import). It invokes the __import__() function from the __builtins__ of the current globals.
// This means that the import is done using whatever import hooks are installed in the current environment.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_Import
func PyImport_Import(name *PyObject) *PyObject {
	return togo(C.PyImport_Import(toc(name)))
}

// PyImport_ReloadModule reloads a module. Return a new reference to the reloaded module, or NULL with an exception set
// on failure (the module still exists in this case).
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_ReloadModule
func PyImport_ReloadModule(name *PyObject) *PyObject {
	return togo(C.PyImport_ReloadModule(toc(name)))
}

// PyImport_AddModuleObject returns the module object corresponding to a module name. The name argument may be of the
// form package.module. First check the modules dictionary if there's one there, and if not, create a new one and insert
// it in the modules dictionary. Return NULL with an exception set on failure.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_AddModuleObject
func PyImport_AddModuleObject(name *PyObject) *PyObject {
	return togo(C.PyImport_AddModuleObject(toc(name)))
}

// PyImport_AddModule is similar to PyImport_AddModuleObject(), but the name is a UTF-8 encoded string instead of a
// Unicode object.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_AddModule
func PyImport_AddModule(name string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PyImport_AddModule(cname))
}

// PyImport_ExecCodeModule loads a module given a module name (possibly of the form package.module) and a code object
// read from a Python bytecode file or obtained from the built-in function compile(). Return a new reference to the
// module object, or NULL with an exception set if an error occurred. name is removed from sys.modules in error cases,
// even if name was already in sys.modules on entry to PyImport_ExecCodeModule(). Leaving incompletely initialized
// modules in sys.modules is dangerous, as imports of such modules have no way to know that the module object is an
// unknown (and probably damaged with respect to the module author's intents) state.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_ExecCodeModule
func PyImport_ExecCodeModule(name string, co *PyObject) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PyImport_ExecCodeModule(cname, toc(co)))
}

// PyImport_ExecCodeModuleEx is like PyImport_ExecCodeModule(), but the __file__ attribute of the module object is set
// to pathname if it is non-NULL.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_ExecCodeModuleEx
func PyImport_ExecCodeModuleEx(name string, co *PyObject, pathname string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cpathname := C.CString(pathname)
	defer C.free(unsafe.Pointer(cpathname))

	return togo(C.PyImport_ExecCodeModuleEx(cname, toc(co), cpathname))
}

// PyImport_ExecCodeModuleObject is like PyImport_ExecCodeModuleEx(), but the __cached__ attribute of the module object
// is set to cpathname if it is non-NULL. Of the three functions, this is the preferred one to use.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_ExecCodeModuleObject
func PyImport_ExecCodeModuleObject(name, co, pathname, cpathname *PyObject) *PyObject {
	return togo(C.PyImport_ExecCodeModuleObject(toc(name), toc(co), toc(pathname), toc(cpathname)))
}

// PyImport_ExecCodeModuleWithPathnames is like PyImport_ExecCodeModuleObject(), but name, pathname and cpathname are
// UTF-8 encoded strings. Attempts are also made to figure out what the value for pathname should be from cpathname if
// the former is set to NULL.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_ExecCodeModuleWithPathnames
func PyImport_ExecCodeModuleWithPathnames(name string, co *PyObject, pathname string, cpathname string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cspathname := C.CString(pathname)
	defer C.free(unsafe.Pointer(cspathname))

	ccpathname := C.CString(cpathname)
	defer C.free(unsafe.Pointer(ccpathname))

	return togo(C.PyImport_ExecCodeModuleWithPathnames(cname, toc(co), cspathname, ccpathname))
}

// PyImport_GetMagicNumber returns the magic number for Python bytecode files (a.k.a. .pyc file). The magic number
// should be present in the first four bytes of the bytecode file, in little-endian byte order. Returns -1 on error.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_GetMagicNumber
func PyImport_GetMagicNumber() int {
	return int(C.PyImport_GetMagicNumber())
}

// PyImport_GetMagicTag : https://docs.python.org/3/c-api/import.html#c.PyImport_GetMagicTag
func PyImport_GetMagicTag() string {
	cmagicTag := C.PyImport_GetMagicTag()

	return C.GoString(cmagicTag)
}

// PyImport_GetModuleDict returns the magic tag string for PEP 3147 format Python bytecode file names. Keep in mind
// that the value at sys.implementation.cache_tag is authoritative and should be used instead of this function.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_GetModuleDict
func PyImport_GetModuleDict() *PyObject {
	return togo(C.PyImport_GetModuleDict())
}

// PyImport_GetModule returns the already imported module with the given name. If the module has not been imported yet
// then returns NULL but does not set an error. Returns NULL and sets an error if the lookup failed.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_GetModule
func PyImport_GetModule(name *PyObject) *PyObject {
	return togo(C.PyImport_GetModule(toc(name)))

}

// PyImport_GetImporter returns a finder object for a sys.path/pkg.__path__ item path, possibly by fetching it from the
// sys.path_importer_cache dict. If it wasn't yet cached, traverse sys.path_hooks until a hook is found that can handle
// the path item. Return None if no hook could; this tells our caller that the path based finder could not find a finder
// for this path item. Cache the result in sys.path_importer_cache. Return a new reference to the finder object.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_GetImporter
func PyImport_GetImporter(path *PyObject) *PyObject {
	return togo(C.PyImport_GetImporter(toc(path)))

}

// PyImport_ImportFrozenModuleObject loads a frozen module named name. Return 1 for success, 0 if the module is not
// found, and -1 with an exception set if the initialization failed. To access the imported module on a successful load,
// use PyImport_ImportModule(). (Note the misnomer — this function would reload the module if it was already imported.)
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_ImportFrozenModuleObject
func PyImport_ImportFrozenModuleObject(name *PyObject) int {
	return int(C.PyImport_ImportFrozenModuleObject(toc(name)))

}

// PyImport_ImportFrozenModule is similar to PyImport_ImportFrozenModuleObject(), but the name is a UTF-8 encoded string
// instead of a Unicode object.
//
// Reference: https://docs.python.org/3/c-api/import.html#c.PyImport_ImportFrozenModule
func PyImport_ImportFrozenModule(name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return int(C.PyImport_ImportFrozenModule(cname))

}
