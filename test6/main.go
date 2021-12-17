package main

import (
	"fmt"
	"math"
)

func main() {
	A := []int{3, 1, 2, 4, 3}

	fmt.Println(SumSubArray(A))
	fmt.Println(Solution(A))
}

func Solution(A []int) (I int) {
	B := SumSubArray(A)

	for i, v := range B {
		if i == 0 {
			I = v
		} 
		
		if v < I {
			I = v
		}

	}
	return I

}

func SepArray(A []int, I int) ([]int, []int) {
	var bef, aft []int
	for i, j := range A {
		if i < I {
			bef = append(bef, j)
		} else {
			aft = append(aft, j)
		}
	}
	return bef, aft
}

func SumArray(A []int) (r int) {
	for _, v := range A {
		r += v
	}
	return r
}

func SumSubArray(A []int) (L []int) {
	i := 1

	for i < len(A) {
		bef, aft := SepArray(A, i)
		Bef := SumArray(bef)
		Aft := SumArray(aft)
		l := Bef - Aft
		// le := float64(l)
		// les := math.Abs(float64(l))
		// less := int(les)
		less := int(math.Abs(float64(l)))
		L = append(L, less)

		i++
	}
	return L
}
