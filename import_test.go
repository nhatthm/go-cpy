package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestImportModule(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	os := cpy.PyImport_ImportModule("os")
	defer os.DecRef()

	assert.NotNil(t, os)
}

func TestImportModuleEx(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	queue := cpy.PyImport_ImportModuleEx("queue", nil, nil, nil)
	defer queue.DecRef()

	assert.NotNil(t, queue)
}

func TestImportModuleLevelObject(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	mathName := cpy.PyUnicode_FromString("math")
	defer mathName.DecRef()

	math := cpy.PyImport_ImportModuleLevelObject(mathName, nil, nil, nil, 0)
	defer math.DecRef()

	assert.NotNil(t, math)
}

func TestImportModuleLevel(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	sys := cpy.PyImport_ImportModuleLevel("sys", nil, nil, nil, 0)
	defer sys.DecRef()

	assert.NotNil(t, sys)
}

func TestImportImport(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	platformName := cpy.PyUnicode_FromString("platform")
	defer platformName.DecRef()

	platform := cpy.PyImport_Import(platformName)
	defer platform.DecRef()

	assert.NotNil(t, platform)
}

func TestReloadModule(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	os := cpy.PyImport_ImportModule("os")
	defer os.DecRef()

	assert.NotNil(t, os)

	newOs := cpy.PyImport_ReloadModule(os)
	defer newOs.DecRef()

	assert.NotNil(t, newOs)

	// cpy3.PyImport_ReloadModule return a new reference, pointer should be the same.
	assert.Equal(t, os, newOs)
}

func TestAddModuleObject(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	os := cpy.PyImport_ImportModule("os")
	defer os.DecRef()

	assert.NotNil(t, os)

	pyName := cpy.PyUnicode_FromString("os.new")
	defer pyName.DecRef()

	module := cpy.PyImport_AddModuleObject(pyName)

	assert.NotNil(t, module)
}

func TestAddModule(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	os := cpy.PyImport_ImportModule("os")
	defer os.DecRef()

	assert.NotNil(t, os)

	module := cpy.PyImport_AddModule("os.new")

	assert.NotNil(t, module)
}

func TestExecCodeModule(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	// Fake module.
	source := cpy.PyUnicode_FromString("__version__ = '2.0'")
	defer source.DecRef()

	filename := cpy.PyUnicode_FromString("test_module.py")
	defer filename.DecRef()

	mode := cpy.PyUnicode_FromString("exec")
	defer mode.DecRef()

	// Perform module load.
	builtins := cpy.PyEval_GetBuiltins()

	assert.True(t, cpy.PyDict_Check(builtins))

	compile := cpy.PyDict_GetItemString(builtins, "compile")

	assert.True(t, cpy.PyCallable_Check(compile))

	code := compile.CallFunctionObjArgs(source, filename, mode)
	defer code.DecRef()

	assert.NotNil(t, code)

	module := cpy.PyImport_ExecCodeModule("test_module", code)

	assert.NotNil(t, module)

	testModuleStr := cpy.PyUnicode_FromString("test_module")
	defer testModuleStr.DecRef()

	pyModule := cpy.PyImport_GetModule(testModuleStr)
	defer pyModule.DecRef()

	assert.NotNil(t, pyModule)
}

func TestExecCodeModuleEx(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	// Fake module.
	source := cpy.PyUnicode_FromString("__version__ = '2.0'")
	defer source.DecRef()

	filename := cpy.PyUnicode_FromString("test_module.py")
	defer filename.DecRef()

	mode := cpy.PyUnicode_FromString("exec")
	defer mode.DecRef()

	// Perform module load.
	builtins := cpy.PyEval_GetBuiltins()

	assert.True(t, cpy.PyDict_Check(builtins))

	compile := cpy.PyDict_GetItemString(builtins, "compile")

	assert.True(t, cpy.PyCallable_Check(compile))

	code := compile.CallFunctionObjArgs(source, filename, mode)
	defer code.DecRef()

	assert.NotNil(t, code)

	module := cpy.PyImport_ExecCodeModuleEx("test_module", code, "test_module.py")

	assert.NotNil(t, module)
}

func TestExecCodeModuleWithPathnames(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	// Fake module.
	source := cpy.PyUnicode_FromString("__version__ = '2.0'")
	defer source.DecRef()

	filename := cpy.PyUnicode_FromString("test_module.py")
	defer filename.DecRef()

	mode := cpy.PyUnicode_FromString("exec")
	defer mode.DecRef()

	// Perform module load.
	builtins := cpy.PyEval_GetBuiltins()

	assert.True(t, cpy.PyDict_Check(builtins))

	compile := cpy.PyDict_GetItemString(builtins, "compile")

	assert.True(t, cpy.PyCallable_Check(compile))

	code := compile.CallFunctionObjArgs(source, filename, mode)
	defer code.DecRef()

	assert.NotNil(t, code)

	module := cpy.PyImport_ExecCodeModuleWithPathnames("test_module", code, "test_module.py", "test_module.py")

	assert.NotNil(t, module)
}

func TestExecCodeModuleObject(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	// Fake module.
	source := cpy.PyUnicode_FromString("__version__ = '2.0'")
	defer source.DecRef()

	filename := cpy.PyUnicode_FromString("test_module.py")
	defer filename.DecRef()

	mode := cpy.PyUnicode_FromString("exec")
	defer mode.DecRef()

	// Perform module load.
	builtins := cpy.PyEval_GetBuiltins()

	assert.True(t, cpy.PyDict_Check(builtins))

	compile := cpy.PyDict_GetItemString(builtins, "compile")

	assert.True(t, cpy.PyCallable_Check(compile))

	code := compile.CallFunctionObjArgs(source, filename, mode)
	defer code.DecRef()

	assert.NotNil(t, code)

	moduleName := cpy.PyUnicode_FromString("test_module")
	defer moduleName.DecRef()

	module := cpy.PyImport_ExecCodeModuleObject(moduleName, code, filename, filename)

	assert.NotNil(t, module)
}

func TestGetMagicNumber(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	magicNumber := cpy.PyImport_GetMagicNumber()

	assert.NotNil(t, magicNumber)
}

func TestGetMagicTag(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	magicTag := cpy.PyImport_GetMagicTag()

	assert.NotNil(t, magicTag)
}

func TestGetModuleDict(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	moduleDict := cpy.PyImport_GetModuleDict()

	defer moduleDict.DecRef()

	assert.True(t, cpy.PyDict_Check(moduleDict))
}

func TestGetModule(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	os := cpy.PyImport_ImportModule("os")
	defer os.DecRef()

	assert.NotNil(t, os)

	name := cpy.PyUnicode_FromString("os")
	defer name.DecRef()

	module := cpy.PyImport_GetModule(name)

	assert.Equal(t, module, os)
}

func TestGetImporter(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	paths := cpy.PySys_GetObject("path")
	path := cpy.PyList_GetItem(paths, 0)

	assert.NotNil(t, path)

	importer := cpy.PyImport_GetImporter(path)
	defer importer.DecRef()

	assert.NotNil(t, importer)
}
