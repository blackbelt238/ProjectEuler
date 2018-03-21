package main

import (
	pythag "ProjectEuler/1-50/9/pythag"
	"fmt"
)

func main() {
	for i := 1; i <= 1000; i++ {
		for j := 1; j <= 1000; j++ {
			for k := 1; k <= 1000; k++ {
				if i+j+k == 1000 {
					if pythag.IsTriplet(i, j, k) {
						fmt.Println(i * j * k)
						return
					}
				}
			}
		}
	}
}
