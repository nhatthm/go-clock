package service

import "github.com/nhatthm/go-clock"

// ClockProvider provides clock.Clock.
type ClockProvider interface {
	Clock() clock.Clock
}
