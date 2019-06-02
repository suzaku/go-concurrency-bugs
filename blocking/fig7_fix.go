package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan int)
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		log.Printf("Acquired lock, start sending")
		select {
		case ch <- 42:
			log.Printf("Sending complete")
		default:
		}
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		log.Printf("Acquiring lock to receive")
		mu.Lock()
		mu.Unlock()
		log.Printf("Receive %d", <- ch)
	}()

	wg.Wait()
}