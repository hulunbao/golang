package main

import "fmt"

type Vertex6 struct{
	Lat,Long float64
}

var m2 = map[string]Vertex6{
	"Bell Labs" : {40.63844,-74.39967},
	"Google" : {37.43234,-122.09473},
}

func main() {
	fmt.Println(m2)
}
