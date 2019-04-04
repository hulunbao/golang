package main

import "fmt"

type Vertex5 struct{
	Lat,Long float64
}

var m1 = map[string]Vertex5{
	"Bell Labs" : Vertex5{
		40.63844,-74.39967,
	},
	"Google" : Vertex5{
		37.43234,-122.09473,
	},
}

func main() {
	fmt.Println(m1)
}
