package generator

import "math/rand/v2"

func generateFloatInRange(from float64, to float64) float64 {
	return rand.Float64()*(to-from) + from
}

type FloatGenerator interface {
	Generate(<-chan bool) <-chan float64
}

type RandomFloatGenerator struct {
	from      float64
	to        float64
	anonymous *AnonymousGenerator[float64]
}

func NewRandomFloatGenerator(from float64, to float64) *RandomFloatGenerator {
	return &RandomFloatGenerator{
		from: from,
		to:   to,
		anonymous: NewAnonymousGenerator[float64](func() float64 {
			return generateFloatInRange(from, to)
		}),
	}
}

func (g *RandomFloatGenerator) Generate(done <-chan bool) <-chan float64 {
	return g.anonymous.Generate(done)
}
