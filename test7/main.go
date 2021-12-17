package main

import "fmt"

func main() {
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	X := 5
	fmt.Println(Solution(X,A))
}

func Solution(X int, A []int) int {
	for i, j := range A {
		if X == j {
			return i
		}
	}
	return -1
}