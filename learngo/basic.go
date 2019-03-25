package main

import "fmt"

func consts(){
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b,kb,mb,gb,tb,pb)
}
func main() {
	fmt.Println("hello world！")
	consts()
}
