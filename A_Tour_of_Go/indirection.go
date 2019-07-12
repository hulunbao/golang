package main

import "fmt"

type Vertex8 struct {
	X,Y float64
}

func (v *Vertex8) Scala(f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScalaFun(v *Vertex8, f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex8{3,4}
	v.Scala(2)
	ScalaFun(&v,10)

	p := &Vertex8{4,3}
	p.Scala(3)
	ScalaFun(p,8)

	fmt.Println(v,p)
}