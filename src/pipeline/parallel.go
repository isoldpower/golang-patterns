package pipeline

import (
	"fmt"
	"go-pipeline/internal/generator"
)

type ParallelPipeline struct {
}

func (pp *ParallelPipeline) Execute() {
	done := make(chan bool)
	defer close(done)

	var floatStream generator.IntGenerator = generator.NewRandomIntGenerator(10000000, 99999999)
	var generationStream <-chan int = floatStream.Generate(done)

	finderRoutine := func() <-chan int {
		return primeFinder(done, generationStream)
	}

	primeStreams := distributeInStreams(finderRoutine)
	flatStream := flatStreams(done, primeStreams...)

	for number := range generator.LimitGeneration(done, flatStream, 10) {
		fmt.Println(number)
	}
}
