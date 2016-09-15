package gostats

import (
	"errors"
	"fmt"
	"github.com/fredericlemoine/gostats/io"
	"math"
	"math/rand"
)

// Poisson returns a random number of possion distribution
func Poisson(lambda float64) int64 {
	if !(lambda > 0.0) {
		io.ExitWithMessage(errors.New(fmt.Sprintf("Invalid lambda: %.2f", lambda)))
	}
	return poisson(lambda)
}

func poisson(lambda float64) int64 {
	// algorithm given by Knuth
	L := math.Pow(math.E, -lambda)
	var k int64 = 0
	var p float64 = 1.0

	for p > L {
		k++
		p *= rand.Float64()
	}
	return (k - 1)
}
