package letter

// Uses waitgroups and a sync.Map.

import (
	"fmt"
	"log"
	"sync"
)

// ConcurrentFrequencyFirstTry performs frequency calculation
// concurrently on a slice of strings.
func ConcurrentFrequencyFirstTry(ss []string) FreqMap {
	totalFrequency := make(FreqMap)

	// Create a concurrency-safe map so multiple goroutines can
	// add frequencies concurrently.
	var safeMap sync.Map

	// Create a waitgroup so execution won't continue until all
	// goroutines have finished working.
	var wg sync.WaitGroup
	wg.Add(len(ss))

	for _, s := range ss {
		go func(s string) {
			defer wg.Done()

			fm := Frequency(s)

			// Once the frequency has been found, add it to the
			// thread-safe map.
			if err := store(&safeMap, fm); nil != err {
				log.Fatal(err)
			}

		}(s)
	}
	wg.Wait()

	// Transfer the contents of the thread-safe map
	// over to the map of total frequency results.
	totalFrequency = transfer(&safeMap, totalFrequency)

	return totalFrequency
}

func store(safeMap *sync.Map, fm FreqMap) error {

	for r, v := range fm {
		res, ok := safeMap.Load(r)
		if !ok {
			safeMap.Store(r, v)
			continue
		}

		var intFreq int
		if intFreq, ok = res.(int); !ok {
			return fmt.Errorf("error converting frequency %v to int", res)
		}

		safeMap.Store(r, intFreq+v)
	}

	return nil
}

func transfer(safeMap *sync.Map, fm FreqMap) FreqMap {
	safeMap.Range(func(k interface{}, v interface{}) bool {
		var ok bool

		// Convert values to FreqMap key and value types.
		var key rune
		if key, ok = k.(rune); !ok {
			return false
		}

		var val int
		if val, ok = v.(int); !ok {
			return false
		}

		// Set the values on the FreqMap.
		fm[key] = val

		return true
	})

	return fm
}
