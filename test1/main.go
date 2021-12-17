package main

// you can also use imports, for example:
import (
	"fmt"
	"sort"
)

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func main() {
	A := []int{1, 2, 3, 6, 4, 1}
	fmt.Println(Solution(A))
}

func Solution(A []int) (i int) {
	// write your code in Go 1.4
	i = 1
	for i <= 100000 {
	sort.Ints(A)
		for _, j := range A {
			// fmt.Print(i)
			if i == j {
				i++
			}
		}
		// i += 1
		return i
	}
	return 1
}
