package main

import (
	triangle "ProjectEuler/1-50/18/triangle"
	"fmt"
)

func main() {
	t, _ := triangle.CreateTriangle("./tri.txt")
	fmt.Println(triangle.MaxPathSum(t))
}
