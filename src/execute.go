package src

import (
	"fmt"
	"go-pipeline/src/pipeline"
)

func Execute() {
	var trackedPipeline Command = &pipeline.ParallelPipeline{}
	fmt.Printf("Parallel execution time: %s\n", trackExecutionTime(trackedPipeline))

	trackedPipeline = &pipeline.SyncPipeline{}
	fmt.Println("Default execution time: ", trackExecutionTime(trackedPipeline))
}
