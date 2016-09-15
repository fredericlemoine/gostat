package gostats

import (
	"math/rand"
)

func Binomial(p float64, nb int) int {
	var binom int = 0
	for i := 0; i < nb; i++ {
		if rand.Float64() < p {
			binom++
		}
	}
	return (binom)
}
