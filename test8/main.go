package main

import (
	"fmt"
	"strings"
)

func main() {
	s := Solution(1,4)
	fmt.Println(s)
}

func Solution(A int, B int) string {

 	var c, ca, cb int
	var ans []string

	C := A+B
	for c < C {
		switch c < C {
		case A > B:
			if ca < 2 {
				A--
				ca++
				c++
				cb = 0
				ans = append(ans, "a")
			} else {
				B--
				cb++
				c++
				ca = 0
				ans = append(ans, "b")
			}
		case A < B:
			if ca < 2 {
				A--
				ca++
				c++
				cb = 0
				ans = append(ans, "a")
			} else {
				B--
				cb++
				c++
				ca = 0
				ans = append(ans, "b")
			}
		case A == B:
			if cb < 2 {
				B--
				cb++
				c++
				ca = 0
				ans = append(ans, "b")
			} else {
				A--
				ca++
				c++
				cb = 0
				ans = append(ans, "a")
			}
		}
	}
	fmt.Println(ans)

	answer := strings.Join(ans, "")
	return answer

}