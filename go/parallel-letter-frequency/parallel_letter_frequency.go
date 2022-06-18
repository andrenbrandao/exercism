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

// BenchmarkSequentialFrequency
// BenchmarkSequentialFrequency-8              2304            479404 ns/op           17728 B/op         15 allocs/op
// BenchmarkConcurrentFrequency
// BenchmarkConcurrentFrequency-8              5232            224873 ns/op           12609 B/op         66 allocs/op
func ConcurrentFrequency(l []string) FreqMap {
	freq := make(FreqMap, 0)
	c := make(chan FreqMap)
	wg := sync.WaitGroup{}

	wg.Add(len(l))

	go func() {
		wg.Wait()
		close(c)
	}()

	for _, s := range l {
		go func(s string, c chan FreqMap) {
			defer wg.Done()

			freq := Frequency(s)
			c <- freq
		}(s, c)
	}

	for m := range c {
		for key, value := range m {
			freq[key] += value
		}
	}

	return freq
}
