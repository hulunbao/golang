package main

import (
	"fmt"
	"math"
)

type I1 interface {
	M()
}

type T1 struct {
	S string
}
func (t *T1) M(){
	fmt.Println(t.S)
}

type F float64

func (f F) M(){
	fmt.Println(f)
}

func main()  {
	var i I1
	i = &T1{"Hello world!"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
	
}

func describe(i I1) {
	fmt.Printf("(%v,%T)\n",i,i)
}