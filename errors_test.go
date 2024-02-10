package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestErrorSetString(t *testing.T) {
	cpy3.Py_Initialize()

	cpy3.PyErr_SetString(cpy3.PyExc_BaseException, "test message")

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_Clear()
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorSetObject(t *testing.T) {
	cpy3.Py_Initialize()

	message := cpy3.PyUnicode_FromString("test message")
	defer message.DecRef()

	cpy3.PyErr_SetObject(cpy3.PyExc_BaseException, message)

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_Print()
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorSetNone(t *testing.T) {
	cpy3.Py_Initialize()

	message := cpy3.PyUnicode_FromString("test message")
	defer message.DecRef()

	cpy3.PyErr_SetNone(cpy3.PyExc_BaseException)

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_Print()
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorSetObjectEx(t *testing.T) {
	cpy3.Py_Initialize()

	message := cpy3.PyUnicode_FromString("test message")
	defer message.DecRef()

	cpy3.PyErr_SetObject(cpy3.PyExc_BaseException, message)

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_PrintEx(false)
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorWriteUnraisable(t *testing.T) {
	cpy3.Py_Initialize()

	message := cpy3.PyUnicode_FromString("unraisable exception")
	defer message.DecRef()

	cpy3.PyErr_WriteUnraisable(message)

	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorBadArgument(t *testing.T) {
	cpy3.Py_Initialize()

	cpy3.PyErr_BadArgument()

	assert.NotNil(t, cpy3.PyErr_Occurred())

	cpy3.PyErr_Clear()

	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorNoMemory(t *testing.T) {
	cpy3.Py_Initialize()

	cpy3.PyErr_NoMemory()

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_Clear()
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorBadInternalCall(t *testing.T) {
	cpy3.Py_Initialize()

	cpy3.PyErr_BadInternalCall()

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_Clear()
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorImportError(t *testing.T) {
	cpy3.Py_Initialize()

	message := cpy3.PyUnicode_FromString("test message")
	defer message.DecRef()

	cpy3.PyErr_SetImportError(message, nil, nil)

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_Clear()
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorImportErrorSubclass(t *testing.T) {
	cpy3.Py_Initialize()

	message := cpy3.PyUnicode_FromString("test message")
	defer message.DecRef()

	cpy3.PyErr_SetImportErrorSubclass(message, nil, nil, cpy3.Dict)

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_Clear()
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorSyntax(t *testing.T) {
	cpy3.Py_Initialize()

	cpy3.PyErr_SetNone(cpy3.PyExc_SyntaxError)

	filename := "test.py"
	cpy3.PyErr_SyntaxLocation(filename, 0)

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_Clear()
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorSyntaxEx(t *testing.T) {
	cpy3.Py_Initialize()

	cpy3.PyErr_SetNone(cpy3.PyExc_SyntaxError)

	filename := "test.py"
	cpy3.PyErr_SyntaxLocationEx(filename, 0, 0)

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_Clear()
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorSyntaxLocation(t *testing.T) {
	cpy3.Py_Initialize()

	cpy3.PyErr_SetNone(cpy3.PyExc_SyntaxError)

	filename := cpy3.PyUnicode_FromString("test.py")
	defer filename.DecRef()

	cpy3.PyErr_SyntaxLocationObject(filename, 0, 0)

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_Clear()
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorExceptionMatches(t *testing.T) {
	cpy3.Py_Initialize()

	cpy3.PyErr_SetNone(cpy3.PyExc_BufferError)

	assert.True(t, cpy3.PyErr_ExceptionMatches(cpy3.PyExc_BufferError))

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_Clear()
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorGivenExceptionMatches(t *testing.T) {
	cpy3.Py_Initialize()

	assert.True(t, cpy3.PyErr_GivenExceptionMatches(cpy3.PyExc_BufferError, cpy3.PyExc_BufferError))
}

func TestErrorFetchRestore(t *testing.T) {
	cpy3.Py_Initialize()

	cpy3.PyErr_SetNone(cpy3.PyExc_BufferError)

	exc, value, traceback := cpy3.PyErr_Fetch()
	assert.Nil(t, cpy3.PyErr_Occurred())

	assert.True(t, cpy3.PyErr_GivenExceptionMatches(exc, cpy3.PyExc_BufferError))
	assert.Nil(t, value)
	assert.Nil(t, traceback)

	cpy3.PyErr_Restore(exc, value, traceback)

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_Clear()
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorNormalizeExceptionRestore(t *testing.T) {
	cpy3.Py_Initialize()

	cpy3.PyErr_SetNone(cpy3.PyExc_BufferError)

	exc, value, traceback := cpy3.PyErr_Fetch()
	exc, value, traceback = cpy3.PyErr_NormalizeException(exc, value, traceback)
	assert.Nil(t, cpy3.PyErr_Occurred())

	assert.True(t, cpy3.PyErr_GivenExceptionMatches(exc, cpy3.PyExc_BufferError))
	assert.Equal(t, 1, value.IsInstance(exc))
	assert.Nil(t, traceback)

	cpy3.PyErr_Restore(exc, value, traceback)

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_Clear()
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorGetSetExcInfo(t *testing.T) {
	cpy3.Py_Initialize()

	cpy3.PyErr_SetNone(cpy3.PyExc_BufferError)

	exc, value, traceback := cpy3.PyErr_GetExcInfo()

	assert.True(t, cpy3.PyErr_GivenExceptionMatches(exc, cpy3.Py_None), cpy3.PyUnicode_AsUTF8(exc.Repr()))
	assert.Equal(t, "None", cpy3.PyUnicode_AsUTF8(value.Repr()))
	assert.Equal(t, "None", cpy3.PyUnicode_AsUTF8(traceback.Repr()))

	cpy3.PyErr_SetExcInfo(exc, value, traceback)

	cpy3.PyErr_Clear()
	assert.Nil(t, cpy3.PyErr_Occurred())
}

func TestErrorInterrupt(t *testing.T) {
	// https://docs.python.org/3/c-api/exceptions.html#c.cpy3.PyErr_CheckSignals
	t.Skip("cpy3.PyErr_CheckSignals unconditionally returns 0 in embedded builds")
	cpy3.Py_Initialize()

	cpy3.PyErr_SetInterrupt()

	assert.Equal(t, -1, cpy3.PyErr_CheckSignals())

	exc := cpy3.PyErr_Occurred()
	assert.True(t, cpy3.PyErr_GivenExceptionMatches(exc, cpy3.PyExc_TypeError))

	assert.NotNil(t, cpy3.PyErr_Occurred())
	cpy3.PyErr_Clear()
	assert.Nil(t, cpy3.PyErr_Occurred())
}
