package main

import "fmt"

func main() {
	fmt.Println(Solution(10, 85, 30))
}

func Solution(X int, Y int, D int) int {
	i := 0
	a := X
	for a < Y {
		i++
		a = X + (D * i)
	}
	return i
}