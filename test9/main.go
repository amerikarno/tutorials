package main

// you can also use imports, for example:
import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func main() {
	init := time.Now()
	b := Solution(21111158)
	total := time.Now().Sub(init)
	fmt.Println(b)
	// fmt.Println(IsSparse(b))
	fmt.Println(total)
}

func Solution(N int) int {
	if N == 0 {
		return N
	}
	var wg sync.WaitGroup
	wg.Add(N)
	for a := 0; a < N; a++ {

		x := make(chan bool)
		y := make(chan bool)
		n := N-a
		go IsSparse(n,x)
		go IsSparse(a,y)
		if <-x && <-y{
			fmt.Println("True")
			return n
		}
		wg.Done()
	}
	wg.Wait()
	return -1
}

//check this int is sparse or not
func IsSparse(A int, m chan bool) {
	n := strconv.FormatInt(int64(A), 2) //convert int to binary format
	N := strings.Split(n,"")
	a := 0
	for _, j := range N {
		if j == "1" {
			a++
			} else {
				a = 0
			}
			if a > 1 {
				m <- false
			}
		}
		m <- true
		close(m)
	}
