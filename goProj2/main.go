package main

import (
	"fmt"
	//"strings"
	"time"

	"./readFiles"
)

func main() {
	r := readFiles.Logs("log.txt")
	n := readFiles.Nline(*r,5)

	lt := readFiles.GetTime("SCBAMCRD",n)
	//logTime := lt.Format("15:04:05.000")

	currentTime := time.Now().Format("15:04:05.000")
	ct,_ := time.Parse("15:04:05.000",currentTime)
	fmt.Println("Log Time: ",lt,"\nCurrent Time: ",ct)
	diff := ct.Sub(*lt)
	fmt.Println("Different Time: ",diff)
	u,_ := time.ParseDuration("30m")
	if diff < u {
		fmt.Println("less than hour")
	} else {
		fmt.Println("more than hour")
	}
}