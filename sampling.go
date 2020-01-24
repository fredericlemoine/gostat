package gostats

import (
	"fmt"
	"math/rand"
)

// This function builds a bootstrap sample by returning
// "sample_size" sampled indices in [0,"sampled_size"[
// with replacement
func Bootstrap(sample_size int) (out_indices []int) {
	return SampleWithReplacement(sample_size, sample_size)
}

// This function builds a new sample by returning
// "output_size" sampled indices with replacement in [0,"input_size"[
func SampleWithReplacement(input_size, output_size int) (out_indices []int) {
	out_indices = make([]int, output_size)

	for i := 0; i < output_size; i++ {
		out_indices[i] = rand.Intn(input_size)
	}
	return out_indices
}

// This function builds a new sample by returning
// "output_size" sampled indices with replacement in [0,length(weights)[
// An index i has a probability of weight[i]/sum(weight) of being sampled
//
// This function uses the Alias Method (https://en.wikipedia.org/wiki/Alias_method)
// The code is inspired from :
// https://github.com/stephaneguindon/phyml/blob/a25f3decdb4c4d5d15bbbf1e2d1275d0b0493ca7/src/stats.c#L4484
func SampleWithReplacementWeighted(weights []float64, output_size int) (out_indices []int, err error) {
	var alias []int
	var small, large []int
	var prob, p []float64
	var i, n int
	var sum float64
	var num_small, num_large int
	var a, g, length int
	var v float64

	length = len(weights)
	out_indices = make([]int, output_size)
	alias = make([]int, length)
	prob = make([]float64, length)
	p = make([]float64, length)
	small = make([]int, length)
	large = make([]int, length)

	sum = .0
	for i, v = range weights {
		if v < 0 {
			err = fmt.Errorf("A weight is < 0")
			return
		}
		sum += v
	}

	if sum == 0. {
		err = fmt.Errorf("The Sum of weights is 0")
		return
	}

	for i, v = range weights {
		p[i] = v * float64(length) / sum
	}

	num_small = 0
	num_large = 0
	for i = length - 1; i >= 0; i-- {
		if p[i] < 1 {
			small[num_small] = i
			num_small++
		} else {
			large[num_large] = i
			num_large++
		}
	}

	for num_small > 0 && num_large > 0 {
		num_small--
		num_large--
		a = small[num_small]
		g = large[num_large]
		prob[a] = p[a]
		alias[a] = g
		p[g] = p[g] + p[a] - 1
		if p[g] < 1 {

			small[num_small] = g
			num_small++
		} else {
			large[num_large] = g
			num_large++
		}
	}

	for num_large > 0 {
		num_large--
		prob[large[num_large]] = 1
	}

	for num_small > 0 {
		num_small--
		prob[small[num_small]] = 1
	}

	var r1, r2 float64
	for n = 0; n < output_size; n++ {
		r1 = rand.Float64()
		r2 = rand.Float64()
		i = int(float64(length) * r1)
		if r2 < prob[i] {
			out_indices[n] = i
		} else {
			out_indices[n] = alias[i]
		}
	}

	return
}
