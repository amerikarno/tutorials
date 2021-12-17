package main

import (
	"fmt"
	"time"
)

// var wg sync.WaitGroup
func main(){

	// wg.Add(2)
	m1 := make(chan string)
	m2 := make(chan string)

	go printHi("m1",m1)

	go printHello("m2",m2)

	for {
		select {
		case msg1 := <-m1:
			fmt.Println(msg1)
		case msg2 := <-m2:
			fmt.Println(msg2)
		default:
			break
		}
	}
}

func printHello(n string, m chan string){
	for i := 0; i < 6; i++{
	 	time.Sleep(2 *time.Second)
		a := fmt.Sprintf("Hello %s",n)
		m <- a
	}
	// wg.Done()
	close(m)
}

func printHi(n string, m chan string){
	for i := 0; i < 6; i++{
		time.Sleep(600 *time.Millisecond)
		a := fmt.Sprintf("Hi %s",n)
	m <- a
	}
	// wg.Done()
	// close(m)
}