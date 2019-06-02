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
	var tc <- chan time.Time
	if dur > 0 {
		tc = time.NewTimer(dur).C
	}
	select {
	case <- tc:
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