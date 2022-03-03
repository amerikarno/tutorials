package main

import (
	"fmt"
	"time"
)

func main() {
	// r1 := []int{3, 4, 3, 0, 2, 2, 3, 0, 0}
	// r2 := []int{4, 2, 0}
	// r3 := []int{4, 4, 3, 3, 1, 0}
	// timelayout := "15:04:05"
	S := "01:23:45"
	T := "01:23:55"
	// s,_ := time.Parse(timelayout, S)
	// t,_ := time.Parse(timelayout, T)
	a := Solution(S,T)
	fmt.Println(a)
	// S := Solution(r1)
	// fmt.Println(S)
	// S = Solution(r2)
	// fmt.Println(S)
	// S = Solution(r3)
	// fmt.Println(S)
}

func Solution(S string,T string) []string {
	l := "15:04:05"
	s,_ := time.Parse(l, S)
	t,_ := time.Parse(l, T)
	var tr []string
	for s.Before(t) {


		s.Add(1 *time.Second)
		S := s.String()
		tr = append(tr, S)
	}
	return tr
}

// func Solution(S string, T string) int