package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"golang/learngo/functional/adder/fib"
	"os"
)

func tryDefer () {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	fmt.Println(4)
}

func writeFile(filename string){
	file, err := os.OpenFile(filename,os.O_EXCL|os.O_CREATE,0666)
	err = errors.New("this is a custom error")
	if err != nil{
		if PathError, ok := err.(*os.PathError); !ok {
			panic(err)
		}else{
			fmt.Println(PathError.Op,PathError.Path,PathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()

	for i := 0;i < 20;i++{
		fmt.Fprintln(writer,f())
	}


}

func main() {
	writeFile("fib.txt")
}
