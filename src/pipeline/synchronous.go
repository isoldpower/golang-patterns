package pipeline

import (
	"fmt"
	"go-patterns/internal/generator"
)

type SyncPipeline struct {
}

func (sp *SyncPipeline) Execute() {
	done := make(chan bool)
	defer close(done)

	var floatStream generator.IntGenerator = generator.NewRandomIntGenerator(10000000, 99999999)
	var generationStream <-chan int = floatStream.Generate(done)

	flatStream := primeFinder(done, generationStream)

	for number := range generator.LimitGeneration(done, flatStream, 10) {
		fmt.Println(number)
	}
}
