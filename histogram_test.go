package des

import "testing"

func TestHistogram(t *testing.T) {
	h := NewHistogram(10, 1, 10)

	for i := 1; i <= 10; i++ {
		for j := 0; j < i; j++ {
			h.AddValue(float64(i))
		}
	}

	h.Print(5)
}
