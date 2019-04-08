package des

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestERG(t *testing.T) {
	sampleData := []float64{1.0, 2.0, 3.0}
	erg := NewEmpericalRandGenerator(1, sampleData)
	var sum float64
	n := 10000000
	lowerBound := sampleData[0]
	upperBound := sampleData[len(sampleData)-1]
	hist := NewHistogram(20, lowerBound, upperBound)

	for i := 0; i < n; i++ {
		r := erg.Next()

		if r < lowerBound {
			assert.Failf(t, "", "r shouldn't be less than %f. Actual: %f", lowerBound, r)
		}
		if r > upperBound {
			assert.Failf(t, "", "r shouldn't be more than %f. Actual: %f", upperBound, r)
		}

		sum += r
		hist.AddValue(r)
	}

	t.Logf("average: %f\n", sum/float64(n))
	hist.Print(20)
}
