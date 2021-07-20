package main

import (
	"fmt"
	"time"
)

func main() {
	number := make(chan int)
	signal := make(chan struct{})
	go receive(signal, number)
	send(number)

	signal <- struct{}{}
}

func send(number chan<- int) {
	number <- 1
	number <- 2
	number <- 3
	number <- 4
	number <- 5
	time.Sleep(time.Nanosecond)
	number <- 6
}

func receive(signal <-chan struct{}, number <-chan int) {
	for {
		select {
		case v := <-number:
			fmt.Println(v)
		case <-signal:
			return
		default:
			fmt.Println("🤔")
		}
	}
}
