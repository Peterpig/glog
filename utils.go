package glog

import (
	"fmt"
	"path"
	"runtime"
)

func formatCaller(rf *runtime.Frame, fullPath bool) string {
	if fullPath {
		return fmt.Sprintf("%s:%d", rf.File, rf.Line)
	}
	return fmt.Sprintf("%s:%d", path.Base(rf.File), rf.Line)
}

// getCaller retrieves the name of the first non-slog calling function
func getCaller(callerSkip int) (fr runtime.Frame, ok bool) {
	pcs := make([]uintptr, 1) // alloc 1 times
	num := runtime.Callers(callerSkip, pcs)
	if num < 1 {
		return
	}

	f, _ := runtime.CallersFrames(pcs).Next()
	return f, f.PC != 0
}
