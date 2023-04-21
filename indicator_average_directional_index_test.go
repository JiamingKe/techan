package techan

import (
	"testing"
)

func TestAverageDirectionalIndexIndicator_Calculate(t *testing.T) {
	const window = 14
	ts := generateTestTimeSeries()
	adx := NewAverageDirectionalIndexIndicator(ts, window)

	decimalEquals(t, 0, adx.Calculate(window-1))
	decimalEquals(t, 4.4290, adx.Calculate(window))

	decimalEquals(t, 15.9932, adx.Calculate(window+3))
	decimalEquals(t, 17.8089, adx.Calculate(window+5))
	decimalEquals(t, 19.0711, adx.Calculate(window+8))

	decimalEquals(t, 23.7423, adx.Calculate(ts.LastIndex()))
}
