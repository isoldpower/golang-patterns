package util

import (
	"fmt"
	"time"
)

type ExecutionDebugger struct {
	start time.Time
	end   time.Time
}

func (ed *ExecutionDebugger) Start() {
	ed.start = time.Now()
}

func (ed *ExecutionDebugger) End() {
	ed.end = time.Now()
}

func (ed *ExecutionDebugger) Summarize() time.Duration {
	duration := ed.end.Sub(ed.start)

	return duration
}

func DebugTime[T func(...any) any](routine T) {
	start := time.Now()
	routine()

	elapsed := time.Since(start)
	fmt.Println("Execution time: ", elapsed)
}
