package pipeline

import (
	"fmt"
	"go-patterns/internal/util"
)

type Command struct {
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Execute() {
	var trackedPipeline util.Command = &ParallelPipeline{}
	fmt.Printf("Parallel execution time: %s\n", util.TrackExecutionTime(trackedPipeline))

	trackedPipeline = &SyncPipeline{}
	fmt.Println("Default execution time: ", util.TrackExecutionTime(trackedPipeline))
}
