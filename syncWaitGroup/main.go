package main

import (
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"http://localhost:1234?duration=3s",
	"http://localhost:1234?duration=1s",
	"http://localhost:1234?duration=5s",
}

func main() {
	fetchSequential(urls)
}

func fetchSequential(urls []string) {
	for _, url := range urls {
		fetch(url)
	}
}

func fetchConcurrent(urls []string) {
	var wg sync.WaitGroup
	wg.Add(3)

	for _, url := range urls {
		go func(u string) {
			fetch(u)
			wg.Done()
		}(url)
	}

	wg.Wait()
}

func fetch(url string) {
	resp, err := http.Head(url)
	if err != nil {
		log.Fatalf("failed url: %s, err: %v", url, err)
	}
	log.Println(url, ": ", resp.StatusCode)
}
