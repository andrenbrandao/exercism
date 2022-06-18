package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.

// Compared the Benchmarks, but the Concurrent version has worse ns/op.
// BenchmarkSequentialFrequency
// BenchmarkSequentialFrequency-8              2437            472813 ns/op           17729 B/op         15 allocs/op
// BenchmarkConcurrentFrequency
// BenchmarkConcurrentFrequency-8               902           1272479 ns/op            4552 B/op         24 allocs/op
func ConcurrentFrequency(l []string) FreqMap {
	freq := make(FreqMap, 0)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, s := range l {
		wg.Add(1)
		go func(s string, freq FreqMap) {
			defer wg.Done()

			for _, r := range s {
				mu.Lock()
				freq[r]++
				mu.Unlock()
			}
		}(s, freq)
	}

	wg.Wait()

	return freq
}
