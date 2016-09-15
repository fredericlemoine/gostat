/**
This package provides Statistical functions
*/
package gostats

import (
	"math"
	"sort"
	// "strconv"
)

func Mean_int(a []int) float64 {
	var mean float64 = 0
	for _, v := range a {
		mean += float64(v)
	}
	return mean / float64(len(a))
}

func Mean_float(a []float64) float64 {
	var mean float64 = 0
	for _, v := range a {
		mean += v
	}
	return mean / float64(len(a))
}

func Median_int(a []int) float64 {
	/* we don't want to modify the original vector, so work on a copy that is going
	   to be sorted: */
	tmp := make([]int, len(a))
	copy(tmp, a)
	sort.Ints(tmp)
	return float64(tmp[int(math.Floor(float64(len(tmp))/2.0))]+tmp[int(math.Ceil(float64(len(tmp))/2.0))]) / 2.0
}

func Median_float(a []float64) float64 {
	/* we don't want to modify the original vector, so work on a copy that is going
	   to be sorted: */
	tmp := make([]float64, len(a))
	copy(tmp, a)
	sort.Float64s(tmp)
	return float64(tmp[int(math.Floor(float64(len(tmp))/2.0))]+tmp[int(math.Ceil(float64(len(tmp))/2.0))]) / 2.0
}

func Sum_int(a []int) int {
	var sum int = 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func Sum_float(a []float64) float64 {
	var sum float64 = 0
	for _, v := range a {
		sum += v
	}
	return sum
}

// Log factorial
func Log_fact(n int) float64 {
	var lf float64 = 0.0
	for i := 2; i <= n; i++ {
		lf = lf + math.Log(float64(i))
	}
	return lf
}

// Log factorial using Ramanujan approximation
func Factorial_log_rmnj(n int) float64 {
	if n == 0 {
		return 0.0
	} else if n <= 100 {
		return (Log_fact(n))
	} else {
		var accu float64 = 0.0
		accu += math.Log(float64(n)*(1.0+4.0*float64(n)*(1.0+2.0*float64(n)))+1.0/30.0-11.0/(240.0*float64(n))) / 6.0
		accu += math.Log(math.Pi) / 2.0
		accu -= float64(n)
		accu += float64(n) * math.Log(float64(n))
		return (accu)
	}
}

func Sigma(values []float64) float64 {
	var vari float64 = 0.0
	mean := Mean_float(values)

	for _, v := range values {
		vari += math.Pow((v - mean), 2)
	}
	return (math.Sqrt(vari))
}
