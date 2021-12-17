package main

import (
	"fmt"
)

func main() {
	A := []int{9, 3, 9, 3, 9, 7, 9}
	a := Solution(A)
	fmt.Println(a)
}

func Solution(A []int) int {

	k := make(map[int]int)

	for _, v := range A {
		fmt.Println(v)
		c, d := k[v]
		fmt.Println(k[v],c,d)
	
		if d {
			k[v] += 1
		} else {
			k[v] =1
		}
	}

	for i, j := range k {
		if j%2 == 1 {
			return i
		}
	}
	return 0
}
