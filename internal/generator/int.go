package generator

import "math/rand/v2"

func generateIntInRange(from int, to int) int {
	return rand.IntN(to-from) + from
}

type IntGenerator interface {
	Generate(<-chan bool) <-chan int
}

type RandomIntGenerator struct {
	from      int
	to        int
	anonymous *AnonymousGenerator[int]
}

func NewRandomIntGenerator(from int, to int) *RandomIntGenerator {
	return &RandomIntGenerator{
		from: from,
		to:   to,
		anonymous: NewAnonymousGenerator[int](func() int {
			return generateIntInRange(from, to)
		}),
	}
}

func (g *RandomIntGenerator) Generate(done <-chan bool) <-chan int {
	return g.anonymous.Generate(done)
}
