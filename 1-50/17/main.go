package main

import (
	numwords "ProjectEuler/1-50/17/numwords"
	"fmt"
)

func main() {
	count := 0
	for i := 0; i <= 1000; i++ {
		word := numwords.Word(i)
		n := numwords.Letters(word)
		count += n
	}
	fmt.Println(count)
}
