package clock

import "time"

// Clock provides time-functionalities.
type Clock interface {
	// Now returns the current local time.
	Now() time.Time
}

type clock struct{}

// Now returns the current local time.
func (clock) Now() time.Time {
	return time.Now()
}

// New creates a live clock that uses `time` package.
func New() Clock {
	return clock{}
}
