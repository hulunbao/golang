package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64{
	z := 1.0
	num := 0.0
	for{
		z -= (z*z - x) /(2*z)
		if math.Abs(num -z) < 1e-10{
			return z
		}else{
			num = z
		}
	}
}
func main() {
	fmt.Println(Sqrt(3))
	fmt.Println(math.Sqrt(3))
}
