package main

import "fmt"

func main() {
	r1 := []int{3, 4, 3, 0, 2, 2, 3, 0, 0}
	r2 := []int{4, 2, 0}
	r3 := []int{4, 4, 3, 3, 1, 0}

	S := Solution(r1)
	fmt.Println(S)
	S = Solution(r2)
	fmt.Println(S)
	S = Solution(r3)
	fmt.Println(S)
}

func Solution(ranks []int) int {
	var report []int
	for _, rank := range ranks {
		supervisor := rank + 1
		for _, s := range ranks {
			if supervisor == s {
				report = append(report, supervisor)
				break
			}
		}
	}
	return len(report)
}