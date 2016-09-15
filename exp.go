package gostats

import (
	"math"
	"math/rand"
)

// Generate from exponential distribution
func Exp(lambda float64) float64 {
	exp := rand.Float64()
	exp = -math.Log(1-exp) / lambda
	return (exp)
}
