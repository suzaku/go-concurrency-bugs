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
		// If the following Wait runs earlier than the Add in the other goroutine,
		// it would return immediately
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
			go func() {
				// Sleep to make the corresponding Wait run before Add
				time.Sleep(200 * time.Millisecond)
				wg.Add(1) 
				time.Sleep(time.Second)
				// We can't see the following message because the Wait returns without waiting
				log.Print("Being idle")
				wg.Done()
			}()
		case "stopped":
		}
	}()

	<-stopped
}