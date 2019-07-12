package main

import (
	"fmt"
	"math"
)

type Vertex7 struct{
	X,Y float64
}

func (v Vertex7) Abs() float64 {
	return math.Abs(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex7) Scala(f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex7{3,4}
	v.Scala(10)
	fmt.Println(v.Abs())
}