package clock

import (
	"time"
)

// StaticClock is the clock that has a fixed timestamp.
type StaticClock struct {
	timestamp time.Time
}

// Now returns a fixed timestamp.
func (c StaticClock) Now() time.Time {
	return c.timestamp
}

// Clock provides clock.Clock.
func (c StaticClock) Clock() Clock {
	return c
}

// Fix creates a fixed clock with a given timestamp.
func Fix(timestamp time.Time) StaticClock {
	return StaticClock{
		timestamp: timestamp,
	}
}
