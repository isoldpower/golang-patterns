package generator

type AnonymousGenerator[T any] struct {
	generateCallback func() T
}

func NewAnonymousGenerator[T any](fetcher func() T) *AnonymousGenerator[T] {
	return &AnonymousGenerator[T]{
		generateCallback: fetcher,
	}
}

func (g *AnonymousGenerator[T]) Generate(done <-chan bool) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- g.generateCallback():
				break
			}
		}
	}()

	return stream
}
