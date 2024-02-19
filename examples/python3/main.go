package main

import (
	"fmt"
	"os"

	"go.nhat.io/cpy/v3"
)

func main() {
	i, err := cpy.Py_Main(os.Args)
	if err != nil {
		fmt.Printf("error launching the python interpreter: %s\n", err)
		os.Exit(1)
	}

	if i == 1 {
		fmt.Println("The interpreter exited due to an exception")
		os.Exit(1)
	}

	if i == 2 {
		fmt.Println("The parameter list does not represent a valid Python command line")
		os.Exit(1)
	}
}
