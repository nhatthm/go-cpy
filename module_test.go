package cpy_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestModuleCheck(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	name := "test_module"

	module := cpy.PyModule_New(name)
	defer module.DecRef()

	assert.True(t, cpy.PyModule_Check(module))
	assert.True(t, cpy.PyModule_CheckExact(module))
}

func TestModuleNew(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	name := "test_module"

	module := cpy.PyModule_New(name)
	defer module.DecRef()

	assert.NotNil(t, module)
}

func TestModuleNewObject(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	name := "test_module"

	pyName := cpy.PyUnicode_FromString(name)
	defer pyName.DecRef()

	assert.NotNil(t, pyName)

	module := cpy.PyModule_NewObject(pyName)
	defer module.DecRef()

	assert.NotNil(t, module)
}

func TestModuleGetDict(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	name := "sys"

	pyName := cpy.PyUnicode_FromString(name)
	defer pyName.DecRef()

	sys := cpy.PyImport_ImportModule(name)
	defer sys.DecRef()

	dict := cpy.PyModule_GetDict(sys)

	assert.True(t, cpy.PyDict_Check(dict))
}

func TestModuleGetName(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	name := "sys"

	pyName := cpy.PyUnicode_FromString(name)
	defer pyName.DecRef()

	sys := cpy.PyImport_ImportModule(name)
	defer sys.DecRef()

	assert.Equal(t, name, cpy.PyModule_GetName(sys))
}

func TestModuleGetNameObject(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	name := "sys"

	pyName := cpy.PyUnicode_FromString(name)
	defer pyName.DecRef()

	sys := cpy.PyImport_ImportModule(name)
	defer sys.DecRef()

	assert.Equal(t, 1, pyName.RichCompareBool(cpy.PyModule_GetNameObject(sys), cpy.Py_EQ))
}

func TestModuleGetState(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	name := "sys"

	pyName := cpy.PyUnicode_FromString(name)
	defer pyName.DecRef()

	sys := cpy.PyImport_ImportModule(name)
	defer sys.DecRef()

	state := cpy.PyModule_GetState(sys)

	assert.Nil(t, state)
}

func TestModuleGetFilenameObject(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	name := "queue"

	queue := cpy.PyImport_ImportModule(name)
	defer queue.DecRef()

	pyFilename := cpy.PyModule_GetFilenameObject(queue)
	filename := cpy.PyUnicode_AsUTF8(pyFilename)

	assert.NotNil(t, pyFilename)
	assert.True(t, strings.HasSuffix(filename, "/queue.py"))
}
