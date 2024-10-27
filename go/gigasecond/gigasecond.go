// Package gigasecond determines the moment after a gigasecond has passed.
package gigasecond

import (
	"time"
)

const (
	gigaFactor = 1000000000
)

// AddGigasecond returns the given time with an added gigasecond.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * gigaFactor)
}
