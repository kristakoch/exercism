package letter

// Uses channels.

// ConcurrentFrequencyThirdTry performs frequency calculation 
// concurrently on a slice of strings.
func ConcurrentFrequencyThirdTry(strs []string) FreqMap {
	freqMap := make(FreqMap)
	freqChan := make(chan FreqMap)

	// Create a goroutine for each element in ss.
	for _, str := range strs {
		go func(str string) {
			freqChan <- Frequency(str)
		}(str)
	}

	// Collect each result and add its values to the
	// map of total frequencies.
	for i := 0; i < len(strs); i++ {
		select {
		case grtFreqMap := <-freqChan:
			for runeVal, freqVal := range grtFreqMap {
				freqMap[runeVal] += freqVal
			}
		}
	}

	return freqMap
}
