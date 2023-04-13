package techan

import "testing"

func TestAlligatorIndicator(t *testing.T) {
	const (
		window = 5
		offset = 3
	)
	t.Run("initial value", func(t *testing.T) {

		series := generateTestTimeSeries()
		indicator := NewAlligatorIndicator(series, window, offset)

		decimalEquals(t, 28139.2400, indicator.Calculate(window+offset-1))
	})

	t.Run("second value", func(t *testing.T) {

		series := generateTestTimeSeries()
		indicator := NewAlligatorIndicator(series, window, offset)

		decimalEquals(t, 28145.5320, indicator.Calculate(window+offset))
	})

	t.Run("third value", func(t *testing.T) {

		series := generateTestTimeSeries()
		indicator := NewAlligatorIndicator(series, window, offset)

		decimalEquals(t, 28156.7956, indicator.Calculate(window+offset+1))
	})
}
