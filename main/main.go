package main

import (
	"fmt"

	"github.com/fredericlemoine/gostats"
)

func main() {
	d1, _ := gostats.Dirichlet1(10.0, 10)
	sum1 := 0.0
	for _, v := range d1 {
		sum1 += v
	}
	fmt.Println(d1)
	fmt.Println(sum1)

	d2, _ := gostats.Dirichlet(10.0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	sum2 := 0.0
	for _, v := range d2 {
		sum2 += v
	}
	fmt.Println(d2)
	fmt.Println(sum2)
}
