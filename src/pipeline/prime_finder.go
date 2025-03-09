package pipeline

import "go-patterns/internal/util"

func primeFinder(done <-chan bool, inputStream <-chan int) <-chan int {
	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case input := <-inputStream:
				if util.IsPrime(input) {
					primes <- input
				}
			}
		}
	}()

	return primes
}
