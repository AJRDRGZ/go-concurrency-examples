package main

import (
	"fmt"
	"sync"
)

func main() {
	courses := make(map[string]string)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(2)

	go func() {
		mu.Lock()
		courses["go desde cero"] = "Intermedio"
		mu.Unlock()
		wg.Done()
	}()

	go func() {
		mu.Lock()
		courses["go concurrencia"] = "Avanzado"
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(courses)
}
