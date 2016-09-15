package gostats

import (
	"math/rand"
)

func Float64Range(a, b int) float64 {
	return float64(a) + rand.Float64()*float64(b-a)
}

func Proba(p float64) bool {
	return (rand.Float64() < p)
}
