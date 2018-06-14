package main

import (
	collatz "ProjectEuler/1-50/14/collatz"
	"fmt"
)

func main() {
	long := 1
	longlen := 1

	for i := 2; i < 1000000; i++ {
		templen := collatz.SequenceLen(i)
		if templen > longlen {
			long = i
			longlen = templen
		}
	}

	fmt.Printf("%d produces the longest chain (%d)\n", long, longlen)
}
