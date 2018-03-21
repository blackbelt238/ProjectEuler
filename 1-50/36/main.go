package main

import (
	palindrome "ProjectEuler/1-50/4/palindrome"
	"fmt"
)

func main() {
	var sum int
	for i := 0; i < 1000000; i++ {
		if palindrome.IsPalindromic(i, 10) {
			if palindrome.IsPalindromic(i, 2) {
				sum += i
			}
		}
	}
	fmt.Println(sum)
}
