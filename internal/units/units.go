package units

import (
	"errors"
	"fmt"
	"math"
)

const (
	// Inputs we accept (strict).
	MinDelayMS = 100
	MaxDelayMS = 2000

	// Derived interval bounds (ms between repeats).
	MinIntervalMS = 20
	MaxIntervalMS = 2000

	// CPS bounds implied by interval constraints: 1000 / 50 = 20ms, 1000 / 0.5 = 2000ms
	MinCPS = 0.5
	MaxCPS = 50.0
)

var (
	ErrInvalidDelay = errors.New("delay_ms must be within [100,2000]")
	ErrInvalidRate  = errors.New("rate_cps must be within (0.5, 50.0] and produce an interval within [20,2000] ms")
)

// RateToIntervalMS converts characters-per-second to millisecond interval,
// rounding to the nearest integer (half away from zero).
func RateToIntervalMS(cps float64) (int, error) {
	if math.IsNaN(cps) || math.IsInf(cps, 0) || cps <= 0 {
		return 0, ErrInvalidRate
	}
	if cps < MinCPS || cps > MaxCPS {
		return 0, ErrInvalidRate
	}
	interval := int(math.Round(1000.0 / cps))
	if interval < MinIntervalMS || interval > MaxIntervalMS {
		return 0, ErrInvalidRate
	}
	return interval, nil
}

// Validate checks delay_ms and rate_cps, and returns the derived interval_ms if valid.
// Use this in the CLI to compute the OS-specific value (GNOME/Windows expect interval_ms).
func Validate(delayMS int, rateCPS float64) (intervalMS int, err error) {
	if delayMS < MinDelayMS || delayMS > MaxDelayMS {
		return 0, fmt.Errorf("%w: got %d", ErrInvalidDelay, delayMS)
	}
	interval, e := RateToIntervalMS(rateCPS)
	if e != nil {
		return 0, fmt.Errorf("%w: got cps=%v", ErrInvalidRate, rateCPS)
	}
	return interval, nil
}
