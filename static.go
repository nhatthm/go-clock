package clock

import (
	"time"
)

type staticClock struct {
	timestamp time.Time
}

// Now returns a fixed timestamp.
func (c staticClock) Now() time.Time {
	return c.timestamp
}

// Fix creates a fixed clock with a given timestamp.
func Fix(timestamp time.Time) Clock {
	return staticClock{
		timestamp: timestamp,
	}
}
