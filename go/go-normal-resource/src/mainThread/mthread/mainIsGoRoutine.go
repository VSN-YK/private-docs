package mthread

import (
	"fmt"
	"runtime"
)

func MainIsGoRoutine() {
	fmt.Printf("Active GoRutine:%d\n", runtime.NumGoroutine())
}
