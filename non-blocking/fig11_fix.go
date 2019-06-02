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
		// Add one more check of the stop signal here in case the select statement
		// picked the ticker branch
		select {
		case <-stopCh:
			return
		default:
		}

		// Simulate some expensive code
		time.Sleep(time.Second)
		log.Print("Working hard")

		select {
		case <-ticker.C:
			log.Print("Next")
		case <-stopCh:
			log.Print("Exit")
			return
		}
	}
}