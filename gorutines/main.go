package main

import (
	"fmt"
	"time"
)

func hello() int {
	fmt.Println("hola, Comunidad EDteam ๐")
	return 1
}

func main() {
	// go hello()
	go func() {
		fmt.Println("hola, Comunidad EDteam desde funciรณn anonima ๐")
	}()

	time.Sleep(time.Millisecond)
	fmt.Println("Hola, Gophers ๐")
}
