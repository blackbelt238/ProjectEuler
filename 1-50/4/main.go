package main

import (
	palindrome "ProjectEuler/1-50/4/palindrome"
	"fmt"
)

func main() {
	pprod := palindrome.Products(3)
	fmt.Println(pprod[len(pprod)-1])
}
