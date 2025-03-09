package generator

type Generator[T any] interface {
	Generate(<-chan bool) <-chan T
}
