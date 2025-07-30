package cpy_test

import (
	"os"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	runtime.LockOSThread()

	ret := m.Run()

	runtime.UnlockOSThread()

	os.Exit(ret)
}
