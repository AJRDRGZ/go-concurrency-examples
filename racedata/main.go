package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(1)

	data := 1

	go func() {
		mu.Lock()
		data++
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()
	mu.Lock()
	fmt.Println(data)
	mu.Unlock()
}
