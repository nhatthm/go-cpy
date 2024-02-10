package cpy3

/*
#include "Python.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Py_Initialize initializes the Python interpreter. In an application embedding Python, this should be called before
// using any other Python/C API functions; see Before Python Initialization for the few exceptions.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_Initialize
func Py_Initialize() {
	C.Py_Initialize()
}

// Py_InitializeEx works like Py_Initialize() if initsigs is 1. If initsigs is 0, it skips initialization registration
// of signal handlers, which might be useful when Python is embedded.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_InitializeEx
func Py_InitializeEx(initsigs bool) {
	if initsigs {
		C.Py_InitializeEx(1)
	} else {
		C.Py_InitializeEx(0)
	}
}

// Py_IsInitialized return true (nonzero) when the Python interpreter has been initialized, false (zero) if not.
// After Py_FinalizeEx() is called, this returns false until Py_Initialize() is called again.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_IsInitialized
func Py_IsInitialized() bool {
	return C.Py_IsInitialized() != 0
}

// Py_FinalizeEx undoes all initializations made by Py_Initialize() and subsequent use of Python/C API functions, and
// destroy all sub-interpreters (see Py_NewInterpreter() below) that were created and not yet destroyed since the last
// call to Py_Initialize(). Ideally, this frees all memory allocated by the Python interpreter. This is a no-op when
// called for a second time (without calling Py_Initialize() again first). Normally the return value is 0. If there were
// errors during finalization (flushing buffered data), -1 is returned.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_FinalizeEx
func Py_FinalizeEx() int {
	return int(C.Py_FinalizeEx())
}

// Py_Finalize is a backwards-compatible version of Py_FinalizeEx() that disregards the return value.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_Finalize
func Py_Finalize() {
	C.Py_Finalize()
}

// Py_GetProgramName returns the program name set with Py_SetProgramName(), or the default. The returned string points
// into static storage; the caller should not modify its value.
//
// This function should not be called before Py_Initialize(), otherwise it returns NULL.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetProgramName
func Py_GetProgramName() (string, error) {
	wcname := C.Py_GetProgramName()
	if wcname == nil {
		return "", nil
	}

	cname := C.Py_EncodeLocale(wcname, nil)

	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}

	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

// Py_GetPrefix returns the prefix for installed platform-independent files. This is derived through a number of
// complicated rules from the program name set with Py_SetProgramName() and some environment variables; for example, if
// the program name is '/usr/local/bin/python', the prefix is '/usr/local'. The returned string points into static
// storage; the caller should not modify its value. This corresponds to the prefix variable in the top-level Makefile
// and the --prefix argument to the configure script at build time. The value is available to Python code as sys.prefix.
// It is only useful on Unix. See also the next function.
//
// This function should not be called before Py_Initialize(), otherwise it returns NULL.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetPrefix
func Py_GetPrefix() (string, error) {
	wcname := C.Py_GetPrefix()
	if wcname == nil {
		return "", nil
	}

	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}

	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

// Py_GetExecPrefix
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetExecPrefix
func Py_GetExecPrefix() (string, error) {
	wcname := C.Py_GetExecPrefix()
	if wcname == nil {
		return "", nil
	}

	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}

	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

// Py_GetProgramFullPath
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetProgramFullPath
func Py_GetProgramFullPath() (string, error) {
	wcname := C.Py_GetProgramFullPath()
	if wcname == nil {
		return "", nil
	}

	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}

	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

// Py_GetPath
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetPath
func Py_GetPath() (string, error) {
	wcname := C.Py_GetPath()
	if wcname == nil {
		return "", nil
	}

	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}

	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

// Py_GetVersion
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetVersion
func Py_GetVersion() string {
	cversion := C.Py_GetVersion()

	return C.GoString(cversion)
}

// Py_GetPlatform
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetPlatform
func Py_GetPlatform() string {
	cplatform := C.Py_GetPlatform()

	return C.GoString(cplatform)
}

// Py_GetCopyright
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetCopyright
func Py_GetCopyright() string {
	ccopyright := C.Py_GetCopyright()

	return C.GoString(ccopyright)
}

// Py_GetCompiler
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetCompiler
func Py_GetCompiler() string {
	ccompiler := C.Py_GetCompiler()

	return C.GoString(ccompiler)
}

// Py_GetBuildInfo
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetBuildInfo
func Py_GetBuildInfo() string {
	cbuildInfo := C.Py_GetBuildInfo()

	return C.GoString(cbuildInfo)
}

// Py_GetPythonHome
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetPythonHome
func Py_GetPythonHome() (string, error) {
	wchome := C.Py_GetPythonHome()
	if wchome == nil {
		return "", nil
	}

	chome := C.Py_EncodeLocale(wchome, nil)
	if chome == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}

	defer C.PyMem_Free(unsafe.Pointer(chome))

	return C.GoString(chome), nil
}
