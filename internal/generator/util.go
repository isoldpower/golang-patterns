package generator

func LimitGeneration[T any](done <-chan bool, generated <-chan T, amount int) <-chan T {
	limit := make(chan T, amount)

	go func() {
		defer close(limit)
		for i := 0; i < amount; i++ {
			select {
			case <-done:
				return
			case limit <- <-generated:
				break
			}
		}
	}()

	return limit
}
