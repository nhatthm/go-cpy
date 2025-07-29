package cpy

/*
#cgo pkg-config: python-3.12-embed
#include "Python.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// Py_Main is the main program for the standard interpreter. This is made available for programs which embed Python.
// The argc and argv parameters should be prepared exactly as those which are passed to a C program's main() function
// (converted to wchar_t according to the user's locale). It is important to note that the argument list may be modified
// (but the contents of the strings pointed to by the argument list are not). The return value will be 0 if the
// interpreter exits normally (i.e., without an exception), 1 if the interpreter exits due to an exception, or 2 if the
// parameter list does not represent a valid Python command line.
//
// "error" will be set if we fail to call "Py_DecodeLocale" on every "args".
//
// Reference: https://docs.python.org/3/c-api/veryhigh.html?highlight=pycompilerflags#c.Py_Main
func Py_Main(args []string) (int, error) {
	argc := C.int(len(args))
	argv := make([]*C.wchar_t, argc)
	defers := make([]func(), 0, argc*2)

	defer func() {
		for _, def := range defers {
			def()
		}
	}()

	for i, arg := range args {
		carg := C.CString(arg)

		defers = append(defers, func() { C.free(unsafe.Pointer(carg)) })

		warg := C.Py_DecodeLocale(carg, nil)
		if warg == nil {
			return -1, fmt.Errorf("fail to call Py_DecodeLocale on '%s'", arg) //nolint: goerr113
		}

		// Py_DecodeLocale requires a call to PyMem_RawFree to free the memory
		defers = append(defers, func() { C.PyMem_RawFree(unsafe.Pointer(warg)) })

		argv[i] = warg
	}

	return int(C.Py_Main(argc, (**C.wchar_t)(unsafe.Pointer(&argv[0])))), nil
}

// PyRun_AnyFile is a simplified interface to PyRun_AnyFileExFlags() below, leaving closeit set to 0 and flags set to
// NULL.
//
// "error" will be set if we fail to open "filename".
//
// Reference: https://docs.python.org/3/c-api/veryhigh.html?highlight=pycompilerflags#c.PyRun_AnyFile
func PyRun_AnyFile(filename string) (int, error) {
	cfilename := C.CString(filename)

	defer C.free(unsafe.Pointer(cfilename))

	mode := C.CString("r")

	defer C.free(unsafe.Pointer(mode))

	cfile, err := C.fopen(cfilename, mode)
	if err != nil {
		return -1, fmt.Errorf("fail to open '%s': %s", filename, err) //nolint: goerr113,errorlint
	}

	defer C.fclose(cfile)

	// C.PyRun_AnyFile is a macro, using C.PyRun_AnyFileFlags instead.
	return int(C.PyRun_AnyFileFlags(cfile, cfilename, nil)), nil
}

// PyRun_SimpleString : https://docs.python.org/3/c-api/veryhigh.html?highlight=pycompilerflags#c.PyRun_SimpleString
func PyRun_SimpleString(command string) int {
	ccommand := C.CString(command)

	defer C.free(unsafe.Pointer(ccommand))

	// C.PyRun_SimpleString is a macro, using C.PyRun_SimpleStringFlags instead
	return int(C.PyRun_SimpleStringFlags(ccommand, nil))
}
