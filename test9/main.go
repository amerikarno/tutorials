package main

// you can also use imports, for example:
import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func main() {
	// a := IsSparse(1166)
	// fmt.Println(a)
	b := Solution(1023)
	fmt.Println(b)
	fmt.Println(IsSparse(b))
}

func Solution(N int) int {

	if N == 0 {
		return N
	}
	
	// var wg sync.WaitGroup
	// var ans []int
	for a:=0;a<N;a++ {
		// wg.Add(N)
		n := N-a
		// fmt.Println(n)

			if IsSparse(n) && IsSparse(a) {
			return n
		}
		// wg.Done()
		
		a++
	}
	// wg.Wait()
	return -1
}

// find interger is sparse or not
func IsSparse(A int) bool {
	n := strconv.FormatInt(int64(A), 2)
	N := strings.Split(n,"")
	// fmt.Println(n)

	a := 0
	for _, j := range N {
		
		if j == "1" {
			a++
			} else {
				a = 0
			}
			
		if a > 1 {
			return false
		}
	}

return true
}
