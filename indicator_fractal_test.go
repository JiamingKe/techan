package techan

import (
	"testing"
)

func TestFractalIndicator_Calculate(t *testing.T) {
	fractal := NewFractalIndicator(generateTestTimeSeries(), 2)

	decimalEquals(t, 0, fractal.Calculate(5))
	decimalEquals(t, 1, fractal.Calculate(7))
	decimalEquals(t, -1, fractal.Calculate(23))
}
