package main

import (
	"fmt"
	"time"
)

func hello() int {
	fmt.Println("hola, Comunidad EDteam 🖐")
	return 1
}

func main() {
	// go hello()
	go func() {
		fmt.Println("hola, Comunidad EDteam desde función anonima 🖐")
	}()

	time.Sleep(time.Millisecond)
	fmt.Println("Hola, Gophers 😎")
}
