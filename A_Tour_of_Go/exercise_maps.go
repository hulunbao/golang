package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int{
	m := make(map[string]int)
	c := strings.Fields(s)
	for _,value := range c{
		m[value] += 1
	}
	return m
}

func main(){
	wc.Test(WordCount)
}
