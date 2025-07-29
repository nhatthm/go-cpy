package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestErrorSetString(t *testing.T) {
	cpy.Py_Initialize()

	cpy.PyErr_SetString(cpy.PyExc_BaseException, "test message")

	assert.NotNil(t, cpy.PyErr_Occurred())
	cpy.PyErr_Clear()
	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorSetObject(t *testing.T) {
	cpy.Py_Initialize()

	message := cpy.PyUnicode_FromString("test message")
	defer message.DecRef()

	cpy.PyErr_SetObject(cpy.PyExc_BaseException, message)

	assert.NotNil(t, cpy.PyErr_Occurred())
	cpy.PyErr_Print()
	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorSetNone(t *testing.T) {
	cpy.Py_Initialize()

	message := cpy.PyUnicode_FromString("test message")
	defer message.DecRef()

	cpy.PyErr_SetNone(cpy.PyExc_BaseException)

	assert.NotNil(t, cpy.PyErr_Occurred())
	cpy.PyErr_Print()
	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorSetObjectEx(t *testing.T) {
	cpy.Py_Initialize()

	message := cpy.PyUnicode_FromString("test message")
	defer message.DecRef()

	cpy.PyErr_SetObject(cpy.PyExc_BaseException, message)

	assert.NotNil(t, cpy.PyErr_Occurred())
	cpy.PyErr_PrintEx(false)
	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorWriteUnraisable(t *testing.T) {
	cpy.Py_Initialize()

	message := cpy.PyUnicode_FromString("unraisable exception")
	defer message.DecRef()

	cpy.PyErr_WriteUnraisable(message)

	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorBadArgument(t *testing.T) {
	cpy.Py_Initialize()

	cpy.PyErr_BadArgument()

	assert.NotNil(t, cpy.PyErr_Occurred())

	cpy.PyErr_Clear()

	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorNoMemory(t *testing.T) {
	cpy.Py_Initialize()

	cpy.PyErr_NoMemory()

	assert.NotNil(t, cpy.PyErr_Occurred())
	cpy.PyErr_Clear()
	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorBadInternalCall(t *testing.T) {
	cpy.Py_Initialize()

	cpy.PyErr_BadInternalCall()

	assert.NotNil(t, cpy.PyErr_Occurred())
	cpy.PyErr_Clear()
	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorImportError(t *testing.T) {
	cpy.Py_Initialize()

	message := cpy.PyUnicode_FromString("test message")
	defer message.DecRef()

	cpy.PyErr_SetImportError(message, nil, nil)

	assert.NotNil(t, cpy.PyErr_Occurred())
	cpy.PyErr_Clear()
	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorImportErrorSubclass(t *testing.T) {
	cpy.Py_Initialize()

	message := cpy.PyUnicode_FromString("test message")
	defer message.DecRef()

	cpy.PyErr_SetImportErrorSubclass(message, nil, nil, cpy.Dict)

	assert.NotNil(t, cpy.PyErr_Occurred())
	cpy.PyErr_Clear()
	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorSyntax(t *testing.T) {
	cpy.Py_Initialize()

	cpy.PyErr_SetNone(cpy.PyExc_SyntaxError)

	filename := "test.py"
	cpy.PyErr_SyntaxLocation(filename, 0)

	assert.NotNil(t, cpy.PyErr_Occurred())
	cpy.PyErr_Clear()
	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorSyntaxEx(t *testing.T) {
	cpy.Py_Initialize()

	cpy.PyErr_SetNone(cpy.PyExc_SyntaxError)

	filename := "test.py"
	cpy.PyErr_SyntaxLocationEx(filename, 0, 0)

	assert.NotNil(t, cpy.PyErr_Occurred())
	cpy.PyErr_Clear()
	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorSyntaxLocation(t *testing.T) {
	cpy.Py_Initialize()

	cpy.PyErr_SetNone(cpy.PyExc_SyntaxError)

	filename := cpy.PyUnicode_FromString("test.py")
	defer filename.DecRef()

	cpy.PyErr_SyntaxLocationObject(filename, 0, 0)

	assert.NotNil(t, cpy.PyErr_Occurred())
	cpy.PyErr_Clear()
	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorExceptionMatches(t *testing.T) {
	cpy.Py_Initialize()

	cpy.PyErr_SetNone(cpy.PyExc_BufferError)

	assert.True(t, cpy.PyErr_ExceptionMatches(cpy.PyExc_BufferError))

	assert.NotNil(t, cpy.PyErr_Occurred())
	cpy.PyErr_Clear()
	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorGivenExceptionMatches(t *testing.T) {
	cpy.Py_Initialize()

	assert.True(t, cpy.PyErr_GivenExceptionMatches(cpy.PyExc_BufferError, cpy.PyExc_BufferError))
}

func TestErrorRaisedException(t *testing.T) {
	cpy.Py_Initialize()

	cpy.PyErr_SetNone(cpy.PyExc_BufferError)

	exc := cpy.PyErr_GetRaisedException()
	assert.Nil(t, cpy.PyErr_Occurred())

	assert.True(t, cpy.PyErr_GivenExceptionMatches(exc, cpy.PyExc_BufferError))

	cpy.PyErr_SetRaisedException(cpy.PyExc_UnicodeError)

	assert.NotNil(t, cpy.PyErr_Occurred())
	cpy.PyErr_Clear()
	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorGetSetExcInfo(t *testing.T) {
	cpy.Py_Initialize()

	cpy.PyErr_SetNone(cpy.PyExc_BufferError)

	exc, value, traceback := cpy.PyErr_GetExcInfo()

	assert.True(t, cpy.PyErr_GivenExceptionMatches(exc, cpy.Py_None), cpy.PyUnicode_AsUTF8(exc.Repr()))
	assert.Equal(t, "None", cpy.PyUnicode_AsUTF8(value.Repr()))
	assert.Equal(t, "None", cpy.PyUnicode_AsUTF8(traceback.Repr()))

	cpy.PyErr_SetExcInfo(exc, value, traceback)

	cpy.PyErr_Clear()
	assert.Nil(t, cpy.PyErr_Occurred())
}

func TestErrorInterrupt(t *testing.T) {
	// https://docs.python.org/3/c-api/exceptions.html#c.cpy3.PyErr_CheckSignals
	t.Skip("cpy3.PyErr_CheckSignals unconditionally returns 0 in embedded builds")
	cpy.Py_Initialize()

	cpy.PyErr_SetInterrupt()

	assert.Equal(t, -1, cpy.PyErr_CheckSignals())

	exc := cpy.PyErr_Occurred()
	assert.True(t, cpy.PyErr_GivenExceptionMatches(exc, cpy.PyExc_TypeError))

	assert.NotNil(t, cpy.PyErr_Occurred())
	cpy.PyErr_Clear()
	assert.Nil(t, cpy.PyErr_Occurred())
}
