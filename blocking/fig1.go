package main

import (
	"log"
	"time"
	"sync"
)

var (
	wg sync.WaitGroup
	execTime time.Duration = time.Second
)

func finishReq(timeout time.Duration) int {
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(execTime) // Simulate time spent calculating the result
		// Timeout would occur before we reach the following line,
		// then no goroutine would be reading from ch,
		// the following line would block.
		ch <- 42 
	}()
	select{
	case result := <- ch:
		return result
	case <-time.After(timeout):
		log.Print("Timeout")
		return -1
	}
}

func main() {
	// Set a timeout that's obviously smaller than the execution time
	timeout := 50 * time.Millisecond 
	log.Printf("Result: %d", finishReq(timeout))
	wg.Wait() // fatal error: all goroutines are asleep - deadlock!
}