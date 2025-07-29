package main

import (
	"fmt"
	"os"

	"go.nhat.io/cpy/v3"
)

func main() {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	if !cpy.Py_IsInitialized() {
		fmt.Println("Error initializing the python interpreter")
		os.Exit(1)
	}

	err := printList()
	if err != nil {
		fmt.Printf("failed to print the python list: %s\n", err)
	}

	cpy.Py_Finalize()
}

func printList() error {
	list := cpy.PyList_New(5)

	if exc := cpy.PyErr_Occurred(); list == nil && exc != nil {
		return fmt.Errorf("fail to create python list object") //nolint: err113
	}

	defer list.DecRef()

	repr, err := pythonRepr(list)
	if err != nil {
		return fmt.Errorf("fail to get representation of object list") //nolint: err113
	}

	fmt.Printf("python list: %s\n", repr)

	return nil
}

func pythonRepr(o *cpy.PyObject) (string, error) {
	if o == nil {
		return "", fmt.Errorf("object is nil") //nolint: err113
	}

	s := o.Repr()
	if s == nil {
		cpy.PyErr_Clear()
		return "", fmt.Errorf("failed to call Repr object method") //nolint: err113
	}

	defer s.DecRef()

	return cpy.PyUnicode_AsUTF8(s), nil
}
