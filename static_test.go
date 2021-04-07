package clock_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/nhatthm/go-clock"
)

func TestStaticClock_Now(t *testing.T) {
	t.Parallel()

	timestamp := time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)
	c := clock.Fix(timestamp)

	assert.Equal(t, timestamp, c.Now())
}

func TestStaticClock_Clock(t *testing.T) {
	t.Parallel()

	c := clock.Fix(time.Now())

	assert.Equal(t, c, c.Clock())
}
