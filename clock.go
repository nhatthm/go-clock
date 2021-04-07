package clock

import "time"

// Clock provides time-functionalities.
type Clock interface {
	// Now returns the current local time.
	Now() time.Time
}

// TimeClock is a clock that uses `time` package as the source.
type TimeClock struct{}

// Now returns the current local time.
func (TimeClock) Now() time.Time {
	return time.Now()
}

// Clock provides clock.Clock.
func (c TimeClock) Clock() Clock {
	return c
}

// New creates a live clock that uses `time` package.
func New() TimeClock {
	return TimeClock{}
}
