package pipeline

import (
	"runtime"
	"sync"
)

func transferRoutine[T any](done <-chan bool, waitGroup *sync.WaitGroup, inputStream <-chan T, outputStream chan<- T) {
	defer waitGroup.Done()
	for value := range inputStream {
		select {
		case <-done:
			return
		case outputStream <- value:
			break
		}
	}
}

func distributeInStreams[T any](routine func() <-chan T) []<-chan T {
	cpuCount := runtime.NumCPU()
	outputStream := make([]<-chan T, cpuCount)
	for i := 0; i < cpuCount; i++ {
		outputStream[i] = routine()
	}

	return outputStream
}

func flatStreams[T any](done <-chan bool, streams ...<-chan T) <-chan T {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(len(streams))
	outputStream := make(chan T)

	for _, stream := range streams {
		go transferRoutine(done, waitGroup, stream, outputStream)
	}

	go func() {
		waitGroup.Wait()
		close(outputStream)
	}()

	return outputStream
}
