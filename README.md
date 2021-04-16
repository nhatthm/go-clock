# Clock

[![Build Status](https://github.com/nhatthm/go-clock/actions/workflows/test.yaml/badge.svg)](https://github.com/nhatthm/go-clock/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/nhatthm/go-clock/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/nhatthm/go-clock)
[![Go Report Card](https://goreportcard.com/badge/github.com/nhatthm/httpmock)](https://goreportcard.com/report/github.com/nhatthm/httpmock)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/github.com/nhatthm/go-clock)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

A library for mocking time in Golang.

## Usage

### Real clock

The clock will return time using `time` package

```go
package mypackage

import "github.com/nhatthm/go-clock"

type Application struct {
	clock clock.Clock
}

func (a *Application) Do() {
	ts := a.clock.Now()

	// Other logic.
}

func New(clock clock.Clock) *Application {
	return &Application{
		clock: clock,
	}
}

var app = New(clock.New())
```

### Static clock

The clock will return a fixed timestamp.

```go
package mypackage

import (
	"time"

	"github.com/nhatthm/go-clock"
)

type Application struct {
	clock clock.Clock
}

func (a *Application) Do() {
	ts := a.clock.Now()
	// ts is "2020-01-02T03:04:05Z"

	// Other logic.
}

func New() *Application {
	return &Application{
		clock: clock.Fix(time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)),
	}
}
```

## Mock

The clock is mocked using [stretchr/testify/mock](https://github.com/stretchr/testify#mock-package)

```go
package mypackage

import (
	"testing"
	"time"

	clockMock "github.com/nhatthm/go-clock/mock"
)

func TestApplication(t *testing.T) {
	t.Parallel()

	ts := time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)

	clock := clockMock.Mock(func(c *mock.Clock) {
		c.On("Now").Return(ts).Once()
	})(t)

	app := New(clock)
	
	// assertions.
}
```

## `cucumber/godog`

See [nhatthm/clockdog](https://github.com/nhatthm/clockdog)

## Donation

If this project help you reduce time to develop, you can give me a cup of coffee :)

### Paypal donation

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or scan this

<img src="https://user-images.githubusercontent.com/1154587/113494222-ad8cb200-94e6-11eb-9ef3-eb883ada222a.png" width="147px" />
