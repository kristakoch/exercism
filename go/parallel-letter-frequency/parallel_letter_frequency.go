package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text
// and returns the data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency performs frequency calculation concurrently
// on a slice of strings.
func ConcurrentFrequency(strs []string) FreqMap {
	freqMap := make(FreqMap)
	freqChan := make(chan FreqMap)

	// Create a goroutine for each element in ss.
	for _, str := range strs {
		go func(s string) {
			freqChan <- Frequency(s)
		}(str)
	}

	// Wait on each result, collect it, and add its values
	// to the map of total frequencies.
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
