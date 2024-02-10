package main

import (
	"fmt"
	"os"

	"go.nhat.io/cpy3"
)

func main() {
	cpy3.Py_Initialize()

	if !cpy3.Py_IsInitialized() {
		fmt.Println("Error initializing the python interpreter")
		os.Exit(1)
	}

	err := printList()
	if err != nil {
		fmt.Printf("failed to print the python list: %s\n", err)
	}

	cpy3.Py_Finalize()
}

func printList() error {
	list := cpy3.PyList_New(5)

	if exc := cpy3.PyErr_Occurred(); list == nil && exc != nil {
		return fmt.Errorf("fail to create python list object") //nolint: goerr113
	}

	defer list.DecRef()

	repr, err := pythonRepr(list)

	if err != nil {
		return fmt.Errorf("fail to get representation of object list") //nolint: goerr113
	}

	fmt.Printf("python list: %s\n", repr)

	return nil
}

func pythonRepr(o *cpy3.PyObject) (string, error) {
	if o == nil {
		return "", fmt.Errorf("object is nil") //nolint: goerr113
	}

	s := o.Repr()
	if s == nil {
		cpy3.PyErr_Clear()
		return "", fmt.Errorf("failed to call Repr object method") //nolint: goerr113
	}

	defer s.DecRef()

	return cpy3.PyUnicode_AsUTF8(s), nil
}
