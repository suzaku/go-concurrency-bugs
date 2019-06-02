// TODO: I don't know how to reproduce the goroutine leak as described in https://github.com/system-pclub/go-concurrency-bugs/blob/e7fc37200aa9b35483ec482148734148712a1830/blocking-bugs/etcd/cb9a3e04b141d150a370f1e6329428b621742041.txt
package main

import (
	"log"
	"time"
	"context"
	"net/http"
	_ "net/http/pprof"
)

func test(ctx context.Context) {
	<-ctx.Done()
	log.Print("Done")
}

func main() {
	go func() { log.Fatal(http.ListenAndServe(":4000", nil)) }()
	
	for i := 1; i < 128; i++ {
		ctx := context.Background()	
		hctx, hcancel := context.WithCancel(ctx)
		if i > 0 {
			hctx, hcancel = context.WithTimeout(ctx, time.Second)
		}
		go test(hctx)
		hcancel()
	}

	log.Print("Visit http://localhost:4000/debug/pprof/goroutine?debug=1 to check goroutines")
	time.Sleep(30 * time.Minute)
}