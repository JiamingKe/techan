package techan

import (
	"testing"
)

func TestDirectionalMovementIndicator_Calculate(t *testing.T) {
	const window = 14
	ts := generateTestTimeSeries()
	adx := NewDirectionalIndexIndicator(ts, window)

	decimalEquals(t, 0, adx.Calculate(window-1))
	decimalEquals(t, 62.0062, adx.Calculate(window))

	decimalEquals(t, 43.7808, adx.Calculate(window+3))
	decimalEquals(t, 26.6950, adx.Calculate(window+5))
	decimalEquals(t, 25.1536, adx.Calculate(window+8))

	decimalEquals(t, 18.1200, adx.Calculate(ts.LastIndex()))
}
