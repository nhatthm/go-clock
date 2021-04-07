package mock_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/nhatthm/go-clock/mock"
)

func TestClock_Now(t *testing.T) {
	t.Parallel()

	timestamp := time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)

	c := mock.Mock(func(c *mock.Clock) {
		c.On("Now").Return(timestamp)
	})(t)

	assert.Equal(t, timestamp, c.Now())
}
