package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestGIL(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	gil := cpy.PyGILState_Ensure()

	assert.True(t, cpy.PyGILState_Check())

	cpy.PyGILState_Release(gil)
}

func TestThreadState(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	threadState := cpy.PyGILState_GetThisThreadState()

	threadState2 := cpy.PyThreadState_Get()

	assert.Equal(t, threadState, threadState2)

	threadState3 := cpy.PyThreadState_Swap(threadState)

	assert.Equal(t, threadState, threadState3)
}

func TestThreadSaveRestore(t *testing.T) {
	cpy.Py_Initialize()
	defer cpy.Py_FinalizeEx()

	threadState := cpy.PyEval_SaveThread()

	assert.False(t, cpy.PyGILState_Check())

	cpy.PyEval_RestoreThread(threadState)
}
