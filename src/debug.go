package src

import (
	"go-patterns/internal/util"
	"time"
)

func trackExecutionTime(command Command) time.Duration {
	debugger := &util.ExecutionDebugger{}
	debugger.Start()
	command.Execute()

	debugger.End()
	return debugger.Summarize()
}
