package des

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	rg := NewNormalRandGenerator(1)

	assert.NotNil(t, rg)

	c1 := 0
	c2 := 0
	for i := 0; i < 1000000; i++ {
		r := rg.Next()
		if r < 0.5 {
			c1++
		} else {
			c2++
		}
	}

	assert.Equal(t, c1+c2, 1000000)
	t.Log(c1)
	t.Log(c2)
}
