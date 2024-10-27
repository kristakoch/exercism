package letter

// Uses waitgroups and a mutex.

import (
	"sync"
)

// ConcurrentFrequencySecondTry performs frequency calculation
// concurrently on a slice of strings.
func ConcurrentFrequencySecondTry(ss []string) FreqMap {
	totalfm := make(FreqMap)
	mutex := &sync.Mutex{}

	// Use a waitgroup to know when all goroutines have finished.
	var wg sync.WaitGroup
	wg.Add(len(ss))

	// Create one goroutine for each element of the slice.
	for _, s := range ss {
		go func(str string) {
			defer wg.Done()

			fm := Frequency(str)

			// Lock the total frequency map while we're adding the
			// results for this goroutine, and unlock when done.
			mutex.Lock()
			for rn, fq := range fm {
				totalfm[rn] += fq
			}
			mutex.Unlock()
		}(s)
	}
	wg.Wait()

	return totalfm
}
