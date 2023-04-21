package techan

import (
	"testing"
)

func TestPositiveDirectionalIndicator_Calculate(t *testing.T) {
	const window = 14
	ts := generateTestTimeSeries()
	di := NewPositiveDirectionalIndicator(ts, window)

	decimalEquals(t, 0, di.Calculate(window-1))
	decimalEquals(t, 41.6035, di.Calculate(window))

	decimalEquals(t, 40.8349, di.Calculate(window+3))
	decimalEquals(t, 35.1506, di.Calculate(window+5))
	decimalEquals(t, 31.5614, di.Calculate(window+8))

	decimalEquals(t, 37.3435, di.Calculate(ts.LastIndex()))
}

func TestNegativeDirectionalIndicator_Calculate(t *testing.T) {
	const window = 14
	ts := generateTestTimeSeries()
	di := NewNegativeDirectionalIndicator(ts, window)

	decimalEquals(t, 0, di.Calculate(window-1))
	decimalEquals(t, 9.7569, di.Calculate(window))

	decimalEquals(t, 15.9667, di.Calculate(window+3))
	decimalEquals(t, 20.3379, di.Calculate(window+5))
	decimalEquals(t, 18.8749, di.Calculate(window+8))

	decimalEquals(t, 25.8863, di.Calculate(ts.LastIndex()))
}
