package service

import "go.nhat.io/clock"

// ClockProvider provides clock.Clock.
type ClockProvider interface {
	Clock() clock.Clock
}
