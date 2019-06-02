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
		// Holding the lock and block, the other goroutine never get a chance to acquire the lock
		ch <- 42 
		log.Printf("Sending complete")
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