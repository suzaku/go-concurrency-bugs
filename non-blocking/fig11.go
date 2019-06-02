package main

import (
	"log"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	ticker := time.NewTicker(100 * time.Millisecond)

	go func() {
		// Wait enough time for 2 run of the expensive code
		time.Sleep(2 * time.Second)
		close(stopCh)
	}()

	for {
		// Simulate some expensive code
		time.Sleep(time.Second)
		log.Print("Working hard")

		// When multiple cases are available, the select statements would pick one randomly
		// So it's possible that the expensive code continue to run even after we get the stop signal
		select {
		case <-ticker.C:
			log.Print("Next")
		case <-stopCh:
			log.Print("Exit")
			return
		}
	}
}