package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	maxGoroutines := 3
	semaphore := make(chan struct{}, maxGoroutines)
	urls := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	for _, url := range urls {
		semaphore <- struct{}{}
		wg.Add(1)
		go func(u string) {
			fmt.Printf("proccessed %s\n", u)
			wg.Done()
			<-semaphore
		}(url)
	}

	wg.Wait()
	close(semaphore)
	fmt.Printf("All goroutines completed.")
}
