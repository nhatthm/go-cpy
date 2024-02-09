package cpy3

/*
#include "Python.h"
*/
import "C"

// PyThreadState represents the state of a single thread.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.PyThreadState
type PyThreadState C.PyThreadState

// PyGILState is an opaque "handle" to the thread state when PyGILState_Ensure() was called, and must be passed to
// PyGILState_Release() to ensure Python is left in the same state
type PyGILState C.PyGILState_STATE

// PyEval_SaveThread releases the global interpreter lock (if it has been created) and reset the thread state to NULL,
// returning the previous thread state (which is not NULL). If the lock has been created, the current thread must have
// acquired it.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.PyEval_SaveThread
func PyEval_SaveThread() *PyThreadState {
	return (*PyThreadState)(C.PyEval_SaveThread())
}

// PyEval_RestoreThread acquires the global interpreter lock (if it has been created) and set the thread state to
// tstate, which must not be NULL. If the lock has been created, the current thread must not have acquired it, otherwise
// deadlock ensues.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.PyEval_RestoreThread
func PyEval_RestoreThread(tstate *PyThreadState) {
	C.PyEval_RestoreThread((*C.PyThreadState)(tstate))
}

// PyThreadState_Get returns the current thread state. The global interpreter lock must be held. When the current thread
// state is NULL, this issues a fatal error (so that the caller needn't check for NULL).
//
// Reference: https://docs.python.org/3/c-api/init.html#c.PyThreadState_Get
func PyThreadState_Get() *PyThreadState {
	return (*PyThreadState)(C.PyThreadState_Get())
}

// PyThreadState_Swap swaps the current thread state with the thread state given by the argument tstate, which may be
// NULL. The global interpreter lock must be held and is not released.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.PyThreadState_Swap
func PyThreadState_Swap(tstate *PyThreadState) *PyThreadState {
	return (*PyThreadState)(C.PyThreadState_Swap((*C.PyThreadState)(tstate)))
}

// PyOS_AfterFork_Child updates internal interpreter state after a process fork. This must be called from the child
// process after calling fork(), or any similar function that clones the current process, if there is any chance the
// process will call back into the Python interpreter. Only available on systems where fork() is defined.
//
// Reference: https://docs.python.org/3/c-api/sys.html#c.PyOS_AfterFork_Child
func PyOS_AfterFork_Child() {
	C.PyOS_AfterFork_Child()
}

// PyGILState_Ensure ensures that the current thread is ready to call the Python C API regardless of the current state
// of Python, or of the global interpreter lock. This may be called as many times as desired by a thread as long as each
// call is matched with a call to PyGILState_Release(). In general, other thread-related APIs may be used between
// PyGILState_Ensure() and PyGILState_Release() calls as long as the thread state is restored to its previous state
// before the Release(). For example, normal usage of the Py_BEGIN_ALLOW_THREADS and Py_END_ALLOW_THREADS macros is
// acceptable.
//
// The return value is an opaque “handle” to the thread state when PyGILState_Ensure() was called, and must be passed to
// PyGILState_Release() to ensure Python is left in the same state. Even though recursive calls are allowed, these
// handles cannot be shared - each unique call to PyGILState_Ensure() must save the handle for its call to
// PyGILState_Release().
//
// When the function returns, the current thread will hold the GIL and be able to call arbitrary Python code. Failure
// is a fatal error.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.PyGILState_Ensure
func PyGILState_Ensure() PyGILState {
	return PyGILState(C.PyGILState_Ensure())
}

// PyGILState_Release releases any resources previously acquired. After this call, Python’s state will be the same as it
// was prior to the corresponding PyGILState_Ensure() call (but generally this state will be unknown to the caller,
// hence the use of the GILState API).
//
// Every call to PyGILState_Ensure() must be matched by a call to PyGILState_Release() on the same thread.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.PyGILState_Release
func PyGILState_Release(state PyGILState) {
	C.PyGILState_Release(C.PyGILState_STATE(state))
}

// PyGILState_GetThisThreadState gets the current thread state for this thread. May return NULL if no GILState API has
// been used on the current thread. Note that the main thread always has such a thread-state, even if no
// auto-thread-state call has been made on the main thread. This is mainly a helper/diagnostic function.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.PyGILState_GetThisThreadState
func PyGILState_GetThisThreadState() *PyThreadState {
	return (*PyThreadState)(C.PyGILState_GetThisThreadState())
}

// PyGILState_Check returns 1 if the current thread is holding the GIL and 0 otherwise. This function can be called from
// any thread at any time. Only if it has had its Python thread state initialized and currently is holding the GIL will
// it return 1. This is mainly a helper/diagnostic function. It can be useful for example in callback contexts or memory
// allocation functions when knowing that the GIL is locked can allow the caller to perform sensitive actions or
// otherwise behave differently.
//
// Reference: https://docs.python.org/3/c-api/init.html#c.PyGILState_Check
func PyGILState_Check() bool {
	return C.PyGILState_Check() == 1
}
