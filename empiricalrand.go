package des

import "sort"

type empericalRandGenerator struct {
	rg         RandomNumberGenerator
	sampleData []float64
}

//NewEmpericalRandGenerator - .........
func NewEmpericalRandGenerator(seed int64, sampleData []float64) RandomNumberGenerator {
	erg := empericalRandGenerator{}
	erg.rg = NewNormalRandGenerator(seed)
	erg.sampleData = append(erg.sampleData, sampleData...)
	sort.Sort(sort.Float64Slice(erg.sampleData))
	return &erg
}

func (erg *empericalRandGenerator) Next() float64 {
	u := erg.rg.Next()
	p := u * (float64(len(erg.sampleData)) - 1)
	j := int(p)
	xj := erg.sampleData[j]
	xk := erg.sampleData[j+1]

	return xj + (p-float64(j))*(xk-xj)
}
