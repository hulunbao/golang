package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	first := 0
	second := 1
	return func() int{
		//temp := first
		first,second = second,(first+second)
		return first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
