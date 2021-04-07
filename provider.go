package clock

// Provider provides clock.Clock.
type Provider interface {
	Clock() Clock
}
