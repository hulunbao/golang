package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int{
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s){
		if lastI, ok := lastOccurred[ch]; ok && lastOccurred[ch] >= start{
			start = lastI + 1
		}
		if i - start + 1 > maxLength{
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcsf"))
	fmt.Println(lengthOfNonRepeatingSubStr("fds"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdefa"))
	fmt.Println(lengthOfNonRepeatingSubStr("我爱北京天安门！"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))
}
