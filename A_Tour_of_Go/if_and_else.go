package main

import (
	"fmt"
	"math"
)

func pow_else(x, n, lim float64) float64 {
	if v := math.Pow(x,n); v < lim{
		return v
	}else{
		fmt.Printf("%g >= %g\n",v,lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow_else(3,2,10),
		pow_else(3,3,20),
		)
}
