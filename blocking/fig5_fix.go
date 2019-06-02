package main

import (
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	var wg sync.WaitGroup	
	wg.Add(len(nums))
	for _ = range nums {
		go func() {
			defer wg.Done()
		}()
	}
	wg.Wait()
}