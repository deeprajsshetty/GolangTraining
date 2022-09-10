package main

import (
	"fmt"
	"sync"
)

const (
	number = 1000
)

var (
	countValue int
	wg         sync.WaitGroup
	mu         sync.Mutex
)

func main() {
	wg.Add(number)
	for i := 0; i < number; i++ {
		// The code m.Lock() within each goroutine protects the global countValue from
		// modification outside of the goroutine in which it is invoked. No other goroutine
		// can change countValue until the m.Unlock() is invoked.
		go func() {
			// This locks the global countValue while each goroutine modifies its value and protects this shared data from being corrupted.
			mu.Lock()
			countValue++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("\ncountValue = %d", countValue)
}

/* Output
countValue = 1000
*/
