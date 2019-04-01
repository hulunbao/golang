package main

import "fmt"

type Vertex1 struct {
	X int
	Y int
}
func main() {
	v := Vertex1{1,3}
	fmt.Println(v.X)
}
