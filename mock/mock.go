package mock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	mockit "github.com/stretchr/testify/mock"

	"github.com/nhatthm/go-clock"
)

// Mocker is Clock mocker.
type Mocker func(tb testing.TB) *Clock

// NoMock is no mock Clock.
var NoMock = Mock()

var (
	_ clock.Clock    = (*Clock)(nil)
	_ clock.Provider = (*Clock)(nil)
)

// Clock is a clock.Clock.
type Clock struct {
	mockit.Mock
}

// Now satisfies clock.Clock.
func (c *Clock) Now() time.Time {
	return c.Called().Get(0).(time.Time)
}

// Clock satisfies clock.Provider.
func (c *Clock) Clock() clock.Clock {
	return c
}

// mock mocks clock.Clock interface.
func mock(mocks ...func(c *Clock)) *Clock {
	c := &Clock{}

	for _, m := range mocks {
		m(c)
	}

	return c
}

// Mock creates Clock mock with cleanup to ensure all the expectations are met.
func Mock(mocks ...func(c *Clock)) Mocker {
	return func(tb testing.TB) *Clock {
		tb.Helper()

		c := mock(mocks...)

		tb.Cleanup(func() {
			assert.True(tb, c.Mock.AssertExpectations(tb))
		})

		return c
	}
}
