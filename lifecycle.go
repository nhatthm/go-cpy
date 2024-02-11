package cpy3

/*
#include "Python.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// ErrEncodeLocaleFailed is returned when we fail to call Py_EncodeLocale.
var ErrEncodeLocaleFailed = fmt.Errorf("fail to call Py_EncodeLocale")

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
		return "", ErrEncodeLocaleFailed
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
		return "", ErrEncodeLocaleFailed
	}

	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

// Py_GetExecPrefix returns the exec-prefix for installed platform-dependent files. This is derived through a number of
// complicated rules from the program name set with Py_SetProgramName() and some environment variables; for example,
// if the program name is '/usr/local/bin/python', the exec-prefix is '/usr/local'. The returned string points into
// static storage; the caller should not modify its value. This corresponds to the exec_prefix variable in the top-level
// Makefile and the --exec-prefix argument to the configure script at build time. The value is available to Python code
// as sys.exec_prefix. It is only useful on Unix.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetExecPrefix
func Py_GetExecPrefix() (string, error) {
	wcname := C.Py_GetExecPrefix()
	if wcname == nil {
		return "", nil
	}

	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", ErrEncodeLocaleFailed
	}

	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

// Py_GetProgramFullPath returns the full program name of the Python executable; this is computed as a side effect of
// deriving the default module search path from the program name (set by Py_SetProgramName() above). The returned string
// points into static storage; the caller should not modify its value. The value is available to Python code as
// `sys.executable`.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetProgramFullPath
func Py_GetProgramFullPath() (string, error) {
	wcname := C.Py_GetProgramFullPath()
	if wcname == nil {
		return "", nil
	}

	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", ErrEncodeLocaleFailed
	}

	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

// Py_GetPath returns the default module search path; this is computed from the program name (set by Py_SetProgramName()
// above) and some environment variables. The returned string consists of a series of directory names separated by a
// platform dependent delimiter character. The delimiter character is ':' on Unix and macOS, ';' on Windows. The
// returned string points into static storage; the caller should not modify its value. The list `sys.path` is
// initialized with this value on interpreter startup; it can be (and usually is) modified later to change the search
// path for loading modules.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetPath
func Py_GetPath() (string, error) {
	wcname := C.Py_GetPath()
	if wcname == nil {
		return "", nil
	}

	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", ErrEncodeLocaleFailed
	}

	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

// Py_GetVersion returns the version of this Python interpreter. This is a string that looks something like:
//
// "3.0a5+ (py3k:63103M, May 12 2008, 00:53:55) \n[GCC 4.2.3]"
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetVersion
func Py_GetVersion() string {
	cversion := C.Py_GetVersion()

	return C.GoString(cversion)
}

// Py_GetPlatform returns the platform identifier for the current platform. On Unix, this is formed from the "official"
// name of the operating system, converted to lower case, followed by the major revision number; e.g., for Solaris 2.x,
// which is also known as SunOS 5.x, the value is 'sunos5'. On macOS, it is 'darwin'. On Windows, it is 'win'. The
// returned string points into static storage; the caller should not modify its value. The value is available to Python
// code as `sys.platform`.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetPlatform
func Py_GetPlatform() string {
	cplatform := C.Py_GetPlatform()

	return C.GoString(cplatform)
}

// Py_GetCopyright returns the official copyright string for the current Python version, for example:
//
// "Copyright 1991-1995 Stichting Mathematisch Centrum, Amsterdam"
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetCopyright
// nolint: misspell
func Py_GetCopyright() string {
	ccopyright := C.Py_GetCopyright()

	return C.GoString(ccopyright)
}

// Py_GetCompiler returns an indication of the compiler used to build the current Python version, in square brackets,
// for example:
//
// "[GCC 2.7.2.2]"
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetCompiler
func Py_GetCompiler() string {
	ccompiler := C.Py_GetCompiler()

	return C.GoString(ccompiler)
}

// Py_GetBuildInfo returns information about the sequence number and build date and time of the current Python
// interpreter instance, for example:
//
// "#67, Aug  1 1997, 22:34:28"
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetBuildInfo
func Py_GetBuildInfo() string {
	cbuildInfo := C.Py_GetBuildInfo()

	return C.GoString(cbuildInfo)
}

// Py_GetPythonHome returns the default "home", that is, the value set by a previous call to Py_SetPythonHome(), or the
// value of the `PYTHONHOME` environment variable if it is set.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.Py_GetPythonHome
func Py_GetPythonHome() (string, error) {
	wchome := C.Py_GetPythonHome()
	if wchome == nil {
		return "", nil
	}

	chome := C.Py_EncodeLocale(wchome, nil)
	if chome == nil {
		return "", ErrEncodeLocaleFailed
	}

	defer C.PyMem_Free(unsafe.Pointer(chome))

	return C.GoString(chome), nil
}
