package util

import (
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

func TrackExecutionTime(command Command) time.Duration {
	debugger := &ExecutionDebugger{}
	debugger.Start()
	command.Execute()

	debugger.End()
	return debugger.Summarize()
}
