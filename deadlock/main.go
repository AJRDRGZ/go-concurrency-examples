package main

import "fmt"

func main() {
	message := make(chan string)

	go func() {
		fmt.Println(<-message)
		//message <- "Hello"
	}()

	fmt.Println(<-message)
}
