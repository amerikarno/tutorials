package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Solution(561892))
}

func Solution(N int) int {
	// write your code in Go 1.4
	a := strconv.FormatInt(int64(N), 2)
	a0 := strings.Split(a,"0")
	var r []string
	for _, str := range a0 {
		if str != "" {
			r = append(r, str)
		}
	}
	
	if len(r) <= 1{
		return 0
	}
	
	a2 := strings.Split(a,"")
	a3 := a2[len(a2)-1]
	// fmt.Println(len(a2))
	
	a1 := strings.Split(a,"1")

	c := 0
	if a3 == "0" {

		for i, j := range a1 {
			// fmt.Println(i)
			if i == len(a1)-1 {
				return c
			}

			if len(j) > c {
				c = len(j)
			}
		}
		return c
	}

	for _, j := range a1 {
		if len(j) > c {
			c = len(j)
		}
	}
	
	return c
}