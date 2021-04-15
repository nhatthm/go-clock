package clock_test

import (
	"testing"
	"time"

	"github.com/nhatthm/go-clock"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		input          string
		expectedResult time.Time
		expectedError  string
	}{
		{
			scenario:      "rfc3339 - error",
			input:         "2020-01-02T03-",
			expectedError: `parsing time "2020-01-02T03-" as "2006-01-02T15:04:05Z07:00": cannot parse "-" as ":"`,
		},
		{
			scenario:       "rfc3339 - success",
			input:          "2020-01-02T03:04:05.000Z",
			expectedResult: time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC),
		},
		{
			scenario:      "ymd - error",
			input:         "2020-01-0",
			expectedError: `parsing time "2020-01-0" as "2006-01-02": cannot parse "0" as "02"`,
		},
		{
			scenario:       "ymd - success",
			input:          "2020-01-02",
			expectedResult: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			result, err := clock.Parse(tc.input)

			assert.Equal(t, tc.expectedResult, result)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}
