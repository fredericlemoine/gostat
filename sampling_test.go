package gostats

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSampleWithReplacementWeighted(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	length := 150000000
	weights := []float64{2., 4., 6., 8., 10.}
	expected := []int{10000000, 20000000, 30000000, 40000000, 50000000}
	counts := make([]int, len(expected))

	indices, err := SampleWithReplacementWeighted(weights, length)

	if err != nil {
		t.Error(err)
	}

	for _, v := range indices {
		counts[v]++
	}

	// The ratio between expected and observed should not be greater than 0.001
	for i, v := range counts {
		ratio := float64(v) / float64(expected[i])
		if ratio > 1.001 || ratio < (1.0/1.001) {
			t.Error(fmt.Errorf("Expected count: %d != Observed count: %d ratio: %f", expected[i], v, ratio))
		}
	}
}

func TestSampleWithReplacement(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	length := 100000000
	expected := []int{20000000, 20000000, 20000000, 20000000, 20000000}
	counts := make([]int, len(expected))

	indices := SampleWithReplacement(len(expected), length)

	for _, v := range indices {
		counts[v]++
	}

	// The ratio between expected and observed should not be greater than 0.001
	for i, v := range counts {
		ratio := float64(v) / float64(expected[i])
		if ratio > 1.001 || ratio < (1.0/1.001) {
			t.Error(fmt.Errorf("Expected count: %d != Observed count: %d", expected[i], v))
		}
	}
}

func TestBootstrap(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	length := 100000000
	expected := []int{20000000, 20000000, 20000000, 20000000, 20000000}
	counts := make([]int, len(expected))

	for is := 0; is < length/len(expected); is++ {
		indices := Bootstrap(len(expected))
		for _, v := range indices {
			counts[v]++
		}
	}

	// The ratio between expected and observed should not be greater than 0.001
	for i, v := range counts {
		ratio := float64(v) / float64(expected[i])
		if ratio > 1.001 || ratio < (1.0/1.001) {
			t.Error(fmt.Errorf("Expected count: %d != Observed count: %d", expected[i], v))
		}
	}
}
