// Package gigasecond contains operations with gigaseconds
package gigasecond

import "time"

const Gigasecond = 1e9 * time.Second

// AddGigasecond adds a gigasecond to an input date
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
