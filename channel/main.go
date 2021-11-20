package main

import (
	"fmt"
	"time"
)

func main(){

	m := make(chan string)
	go chanM(m)
	m <- "M"
	go printA()
 	go printB()
}

func printA(){
	time.Sleep(5 *time.Second)
	fmt.Println("A")
}

func printB(){
	time.Sleep(5 *time.Second)
	fmt.Println("B")
}

func chanM(m chan string){
	fmt.Println(<-m)
}