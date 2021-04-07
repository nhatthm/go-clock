package clock_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/nhatthm/go-clock"
)

func TestClock_Now(t *testing.T) {
	t.Parallel()

	c := clock.New()
	timestamp := time.Now()
	now := c.Now()

	assert.True(t, timestamp.Before(now) || timestamp.Equal(now))
}

func TestClock_Clock(t *testing.T) {
	t.Parallel()

	c := clock.New()

	assert.Equal(t, c, c.Clock())
}
