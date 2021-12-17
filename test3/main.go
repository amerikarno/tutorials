package main

import "fmt"

func main() {

	A := []int{3, 8, 9, 7, 6}

	fmt.Println(Solution(A,2))

}

func Solution(A []int, K int) []int {
	var a []int
	for i := 0; i < K; i++{
		A = Rotation(A)
	}
	a = A
	return a
}

func Rotation(A []int) []int {
	var r []int
	// a := len(A)

	
	for i, j := range A {
		if i == len(A) - 1 {
			r = append([]int{j}, r...)
		} else {
			r = append(r, j)
		}
		// fmt.Println(j)
	}

	return r
}