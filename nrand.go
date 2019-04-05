package des

import "math/rand"

type normalRandGenerator struct {
	seed int
	rg   *rand.Rand
}

//NewNormalRandGenerator - .............
func NewNormalRandGenerator(seed int64) RandomNumberGenerator {
	return &normalRandGenerator{
		rg: rand.New(rand.NewSource(seed)),
	}
}

func (nrg normalRandGenerator) Next() float64 {
	return nrg.rg.Float64()
}
