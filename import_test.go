package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestImportModule(t *testing.T) {
	cpy3.Py_Initialize()

	os := cpy3.PyImport_ImportModule("os")
	defer os.DecRef()

	assert.NotNil(t, os)
}

func TestImportModuleEx(t *testing.T) {
	cpy3.Py_Initialize()

	queue := cpy3.PyImport_ImportModuleEx("queue", nil, nil, nil)
	defer queue.DecRef()

	assert.NotNil(t, queue)
}

func TestImportModuleLevelObject(t *testing.T) {
	cpy3.Py_Initialize()

	mathName := cpy3.PyUnicode_FromString("math")
	defer mathName.DecRef()

	math := cpy3.PyImport_ImportModuleLevelObject(mathName, nil, nil, nil, 0)
	defer math.DecRef()

	assert.NotNil(t, math)
}

func TestImportModuleLevel(t *testing.T) {
	cpy3.Py_Initialize()

	sys := cpy3.PyImport_ImportModuleLevel("sys", nil, nil, nil, 0)
	defer sys.DecRef()

	assert.NotNil(t, sys)
}

func TestImportImport(t *testing.T) {
	cpy3.Py_Initialize()

	platformName := cpy3.PyUnicode_FromString("platform")
	defer platformName.DecRef()

	platform := cpy3.PyImport_Import(platformName)
	defer platform.DecRef()

	assert.NotNil(t, platform)
}

func TestReloadModule(t *testing.T) {
	cpy3.Py_Initialize()

	os := cpy3.PyImport_ImportModule("os")
	defer os.DecRef()

	assert.NotNil(t, os)

	newOs := cpy3.PyImport_ReloadModule(os)
	defer newOs.DecRef()

	assert.NotNil(t, newOs)

	// cpy3.PyImport_ReloadModule return a new reference, pointer should be the same.
	assert.Equal(t, os, newOs)
}

func TestAddModuleObject(t *testing.T) {
	cpy3.Py_Initialize()

	os := cpy3.PyImport_ImportModule("os")
	defer os.DecRef()

	assert.NotNil(t, os)

	pyName := cpy3.PyUnicode_FromString("os.new")
	defer pyName.DecRef()

	module := cpy3.PyImport_AddModuleObject(pyName)

	assert.NotNil(t, module)
}

func TestAddModule(t *testing.T) {
	cpy3.Py_Initialize()

	os := cpy3.PyImport_ImportModule("os")
	defer os.DecRef()

	assert.NotNil(t, os)

	module := cpy3.PyImport_AddModule("os.new")

	assert.NotNil(t, module)
}

func TestExecCodeModule(t *testing.T) {
	cpy3.Py_Initialize()

	// Fake module.
	source := cpy3.PyUnicode_FromString("__version__ = '2.0'")
	defer source.DecRef()

	filename := cpy3.PyUnicode_FromString("test_module.py")
	defer filename.DecRef()

	mode := cpy3.PyUnicode_FromString("exec")
	defer mode.DecRef()

	// Perform module load.
	builtins := cpy3.PyEval_GetBuiltins()

	assert.True(t, cpy3.PyDict_Check(builtins))

	compile := cpy3.PyDict_GetItemString(builtins, "compile")

	assert.True(t, cpy3.PyCallable_Check(compile))

	code := compile.CallFunctionObjArgs(source, filename, mode)
	defer code.DecRef()

	assert.NotNil(t, code)

	module := cpy3.PyImport_ExecCodeModule("test_module", code)

	assert.NotNil(t, module)

	testModuleStr := cpy3.PyUnicode_FromString("test_module")
	defer testModuleStr.DecRef()

	pyModule := cpy3.PyImport_GetModule(testModuleStr)
	defer pyModule.DecRef()

	assert.NotNil(t, pyModule)
}

func TestExecCodeModuleEx(t *testing.T) {
	cpy3.Py_Initialize()

	// Fake module.
	source := cpy3.PyUnicode_FromString("__version__ = '2.0'")
	defer source.DecRef()

	filename := cpy3.PyUnicode_FromString("test_module.py")
	defer filename.DecRef()

	mode := cpy3.PyUnicode_FromString("exec")
	defer mode.DecRef()

	// Perform module load.
	builtins := cpy3.PyEval_GetBuiltins()

	assert.True(t, cpy3.PyDict_Check(builtins))

	compile := cpy3.PyDict_GetItemString(builtins, "compile")

	assert.True(t, cpy3.PyCallable_Check(compile))

	code := compile.CallFunctionObjArgs(source, filename, mode)
	defer code.DecRef()

	assert.NotNil(t, code)

	module := cpy3.PyImport_ExecCodeModuleEx("test_module", code, "test_module.py")

	assert.NotNil(t, module)
}

func TestExecCodeModuleWithPathnames(t *testing.T) {
	cpy3.Py_Initialize()

	// Fake module.
	source := cpy3.PyUnicode_FromString("__version__ = '2.0'")
	defer source.DecRef()

	filename := cpy3.PyUnicode_FromString("test_module.py")
	defer filename.DecRef()

	mode := cpy3.PyUnicode_FromString("exec")
	defer mode.DecRef()

	// Perform module load.
	builtins := cpy3.PyEval_GetBuiltins()

	assert.True(t, cpy3.PyDict_Check(builtins))

	compile := cpy3.PyDict_GetItemString(builtins, "compile")

	assert.True(t, cpy3.PyCallable_Check(compile))

	code := compile.CallFunctionObjArgs(source, filename, mode)
	defer code.DecRef()

	assert.NotNil(t, code)

	module := cpy3.PyImport_ExecCodeModuleWithPathnames("test_module", code, "test_module.py", "test_module.py")

	assert.NotNil(t, module)
}

func TestExecCodeModuleObject(t *testing.T) {
	cpy3.Py_Initialize()

	// Fake module.
	source := cpy3.PyUnicode_FromString("__version__ = '2.0'")
	defer source.DecRef()

	filename := cpy3.PyUnicode_FromString("test_module.py")
	defer filename.DecRef()

	mode := cpy3.PyUnicode_FromString("exec")
	defer mode.DecRef()

	// Perform module load.
	builtins := cpy3.PyEval_GetBuiltins()

	assert.True(t, cpy3.PyDict_Check(builtins))

	compile := cpy3.PyDict_GetItemString(builtins, "compile")

	assert.True(t, cpy3.PyCallable_Check(compile))

	code := compile.CallFunctionObjArgs(source, filename, mode)
	defer code.DecRef()

	assert.NotNil(t, code)

	moduleName := cpy3.PyUnicode_FromString("test_module")
	defer moduleName.DecRef()

	module := cpy3.PyImport_ExecCodeModuleObject(moduleName, code, filename, filename)

	assert.NotNil(t, module)
}

func TestGetMagicNumber(t *testing.T) {
	cpy3.Py_Initialize()

	magicNumber := cpy3.PyImport_GetMagicNumber()

	assert.NotNil(t, magicNumber)
}

func TestGetMagicTag(t *testing.T) {
	cpy3.Py_Initialize()

	magicTag := cpy3.PyImport_GetMagicTag()

	assert.NotNil(t, magicTag)
}

func TestGetModuleDict(t *testing.T) {
	cpy3.Py_Initialize()

	moduleDict := cpy3.PyImport_GetModuleDict()

	defer moduleDict.DecRef()

	assert.True(t, cpy3.PyDict_Check(moduleDict))
}

func TestGetModule(t *testing.T) {
	cpy3.Py_Initialize()

	os := cpy3.PyImport_ImportModule("os")
	defer os.DecRef()

	assert.NotNil(t, os)

	name := cpy3.PyUnicode_FromString("os")
	defer name.DecRef()

	module := cpy3.PyImport_GetModule(name)

	assert.Equal(t, module, os)
}

func TestGetImporter(t *testing.T) {
	cpy3.Py_Initialize()

	paths := cpy3.PySys_GetObject("path")
	path := cpy3.PyList_GetItem(paths, 0)

	assert.NotNil(t, path)

	importer := cpy3.PyImport_GetImporter(path)
	defer importer.DecRef()

	assert.NotNil(t, importer)
}
