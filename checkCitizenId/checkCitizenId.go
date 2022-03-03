package main

import (
	"fmt"
	"strconv"
)

func main() {
	cid := "3102000378645"
	// cid1 := "1234567890121"
	k := 0
	for i, j := range cid{
		if i == len(cid)-1 {
			break
		}
		k += (int(j)-48) * (13-i)
		fmt.Println(i,j,k)
	}

	l := (11-(k%11))%10

	if strconv.Itoa(l) == string(cid[12]) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	fmt.Println(l)
}