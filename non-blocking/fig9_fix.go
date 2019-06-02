package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	status := "idle"
	var m sync.Mutex
	var wg sync.WaitGroup
	stopped := make(chan struct{})

	go func() {
		// Sleep to allow the other goroutine to acquire the lock first,
		// so that the idle branch is reached
		time.Sleep(50 * time.Millisecond)
		m.Lock()
		status = "stopped"
		defer m.Unlock()
		log.Print("Start waiting")
		wg.Wait()
		log.Print("Done waiting")
		close(stopped)
	}()

	go func() {
		m.Lock()
		defer m.Unlock()

		switch status {
		case "idle":
			// The Add is move out so that it's inside the critical section,
			// now we can be sure when the idle branch is reached,
			// Add would happen before Wait
			wg.Add(1) 
			go func() {
				time.Sleep(time.Second)
				log.Print("Being idle")
				wg.Done()
			}()
		case "stopped":
		}
	}()

	<-stopped
}