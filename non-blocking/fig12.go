package main

import (
	"log"
	"time"
	"context"
)

// Intension: run until timeout or canceled
// * If `dur` is greater than 0, set timeout to be `dur`
// * Otherwise, there's no timeout, run until canceled.
func timeout(ctx context.Context, dur time.Duration) {
	timer := time.NewTimer(0)
	if dur > 0 {
		timer = time.NewTimer(dur)
	}
	select {
	case <- timer.C:
		log.Print("Timeout")
	case <-ctx.Done():
		log.Print("Stopped")
	}
}

func main() {
	log.Print("Start")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	timeout(ctx, 0)
}