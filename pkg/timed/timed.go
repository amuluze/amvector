// Package timed
// Date: 2024/06/10 14:31:29
// Author: Amu
// Description:
package timed

import (
	"time"
)

type Ticker interface {
	Chan() <-chan time.Time
	Stop()
}

type realTicker struct {
	*time.Ticker
}

// NewTicker returns a Ticker.
func NewTicker(d time.Duration) Ticker {
	return &realTicker{
		Ticker: time.NewTicker(d),
	}
}

func (r *realTicker) Chan() <-chan time.Time {
	return r.C
}

func (r *realTicker) Stop() {
	r.Ticker.Stop()
}
