package cpy_test

import (
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	m.Run()
}
