package main

import (
	"bufio"
	"fmt"
	"golang/learngo/functional/adder/fib"
	"io"
	"strings"
)

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error){
	next := g()
	if next > 10000{
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n",next)
	// TODO: incorrect if p is too small
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fib.Fibonacci()

	printFileContents(f)
}
