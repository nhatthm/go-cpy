package cpy3_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestModuleCheck(t *testing.T) {
	cpy3.Py_Initialize()

	name := "test_module"

	module := cpy3.PyModule_New(name)
	assert.True(t, cpy3.PyModule_Check(module))
	assert.True(t, cpy3.PyModule_CheckExact(module))
	defer module.DecRef()
}

func TestModuleNew(t *testing.T) {
	cpy3.Py_Initialize()

	name := "test_module"

	module := cpy3.PyModule_New(name)
	assert.NotNil(t, module)
	defer module.DecRef()
}

func TestModuleNewObject(t *testing.T) {
	cpy3.Py_Initialize()

	name := "test_module"

	pyName := cpy3.PyUnicode_FromString(name)
	assert.NotNil(t, pyName)
	defer pyName.DecRef()

	module := cpy3.PyModule_NewObject(pyName)
	assert.NotNil(t, module)
	defer module.DecRef()
}

func TestModuleGetDict(t *testing.T) {
	cpy3.Py_Initialize()

	name := "sys"
	pyName := cpy3.PyUnicode_FromString(name)
	defer pyName.DecRef()

	sys := cpy3.PyImport_ImportModule(name)
	defer sys.DecRef()

	dict := cpy3.PyModule_GetDict(sys)
	assert.True(t, cpy3.PyDict_Check(dict))
}

func TestModuleGetName(t *testing.T) {
	cpy3.Py_Initialize()

	name := "sys"
	pyName := cpy3.PyUnicode_FromString(name)
	defer pyName.DecRef()

	sys := cpy3.PyImport_ImportModule(name)
	defer sys.DecRef()

	assert.Equal(t, name, cpy3.PyModule_GetName(sys))
}

func TestModuleGetNameObject(t *testing.T) {
	cpy3.Py_Initialize()

	name := "sys"
	pyName := cpy3.PyUnicode_FromString(name)
	defer pyName.DecRef()

	sys := cpy3.PyImport_ImportModule(name)
	defer sys.DecRef()

	assert.Equal(t, 1, pyName.RichCompareBool(cpy3.PyModule_GetNameObject(sys), cpy3.Py_EQ))
}

func TestModuleGetState(t *testing.T) {
	cpy3.Py_Initialize()

	name := "sys"
	pyName := cpy3.PyUnicode_FromString(name)
	defer pyName.DecRef()

	sys := cpy3.PyImport_ImportModule(name)
	defer sys.DecRef()

	state := cpy3.PyModule_GetState(sys)
	assert.True(t, state == nil)
}

func TestModuleGetFilenameObject(t *testing.T) {
	cpy3.Py_Initialize()

	name := "queue"
	queue := cpy3.PyImport_ImportModule(name)
	defer queue.DecRef()

	pyFilename := cpy3.PyModule_GetFilenameObject(queue)
	assert.NotNil(t, pyFilename)
	filename := cpy3.PyUnicode_AsUTF8(pyFilename)

	assert.True(t, strings.HasSuffix(filename, "/queue.py"))
}
