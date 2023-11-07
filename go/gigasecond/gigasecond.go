// Package gigasecond adds 1 GigaSecond to a given timestamp
package gigasecond

import (
	"time"
)

// AddGigasecond adds 10^9 seconds to a given time
func AddGigasecond(t time.Time) time.Time {

	return t.Add(1e9 * time.Second)
}
