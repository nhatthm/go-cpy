package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestGIL(t *testing.T) {
	cpy3.Py_Initialize()

	gil := cpy3.PyGILState_Ensure()

	assert.True(t, cpy3.PyGILState_Check())

	cpy3.PyGILState_Release(gil)
}

func TestThreadState(t *testing.T) {
	cpy3.Py_Initialize()

	threadState := cpy3.PyGILState_GetThisThreadState()

	threadState2 := cpy3.PyThreadState_Get()

	assert.Equal(t, threadState, threadState2)

	threadState3 := cpy3.PyThreadState_Swap(threadState)

	assert.Equal(t, threadState, threadState3)
}

func TestThreadSaveRestore(t *testing.T) {
	cpy3.Py_Initialize()

	threadState := cpy3.PyEval_SaveThread()

	assert.False(t, cpy3.PyGILState_Check())

	cpy3.PyEval_RestoreThread(threadState)
}
