package des

import (
	"fmt"
	"math"
)

//Histogram -----
type Histogram struct {
	slots    int
	minValue float64
	maxValue float64
	delta    float64
	Result   []float64
}

//NewHistogram - ......
func NewHistogram(slots int, minValue, maxValue float64) *Histogram {
	delta := ((maxValue - minValue) / float64(slots-1))
	return &Histogram{
		slots:    slots,
		minValue: minValue,
		maxValue: maxValue,
		delta:    delta,
		Result:   make([]float64, slots),
	}
}

//AddValue - .......s
func (h *Histogram) AddValue(value float64) {
	slot := int(math.Round(value/h.delta)) - 1
	h.Result[slot]++
}

//Print - ........
func (h *Histogram) Print(height int) {
	var maxValue float64
	for _, v := range h.Result {
		if v > maxValue {
			maxValue = v
		}
	}

	step := maxValue / float64(height)

	for r := height; r > 0; r-- {
		for _, v := range h.Result {
			if v >= step*float64(r) {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
