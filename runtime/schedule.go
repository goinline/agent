// Copyright 2023 冯立强 fenglq@tingyun.com.  All rights reserved.
// +build linux
// +build amd64 arm64
// +build cgo

package wrapruntime

import (
	"fmt"
	"runtime"

	"github.com/goinline/agent"
)

//go:noinline
func runtimeSchedule() {
	trampoline.arg1 = *trampoline.idpointer + trampoline.idindex + trampoline.arg1 + trampoline.arg2 + trampoline.arg3 + trampoline.arg4 + trampoline.arg5 + trampoline.arg6 + trampoline.arg7 +
		trampoline.arg8 + trampoline.arg9 + trampoline.arg10 + trampoline.arg11 + trampoline.arg12 + trampoline.arg13 + trampoline.arg14 + trampoline.arg15 + trampoline.arg16 +
		trampoline.arg17 + trampoline.arg18 + trampoline.arg19 + trampoline.arg20
}

//go:noinline
func WrapruntimeSchedule() {

	stackList := [8]uintptr{0, 0, 0, 0, 0, 0, 0, 0}
	count := runtime.Callers(1, stackList[:8])
	out := false
	for i := 0; i < count; i++ {
		name, file := getnameByAddr(stackList[i] - 1)
		if name == "runtime.main" {
			out = true
			fmt.Println("Routine:", tingyun3.GetGID())
		}
		if out {
			fmt.Println(i, name+":"+file)
		}
	}
	runtimeSchedule()
}
func getnameByAddr(p uintptr) (string, string) {
	if r := runtime.FuncForPC(p); r == nil {
		return "", ""
	} else {
		file, line := r.FileLine(p)
		return r.Name(), fmt.Sprintf("%s:%d", file, line)
	}
}
