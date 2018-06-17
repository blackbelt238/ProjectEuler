package main

import (
	longmath "ProjectEuler/1-50/13/longmath"
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./problem-13.txt")
	if err != nil {
		fmt.Printf("could not open file: %v\n", err)
		return
	}
	defer file.Close()

	var n1, n2 string
	n1 = "0"

	// go through each number in the file, summing it with the sum of all the previous numbers
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n2 = scanner.Text()
		n1 = longmath.LongAdd(n1, n2)
	}

	fmt.Printf("The sum is %v...\n", n1[:10])
}
