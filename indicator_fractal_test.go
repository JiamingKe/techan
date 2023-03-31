package techan

import (
	"math"
	"testing"
)

func TestUpFractalIndicator_Calculate(t *testing.T) {
	const window = 2
	ts := generateTestTimeSeries()
	fractal := NewUpFractalIndicator(ts, window)

	decimalEquals(t, math.MaxInt, fractal.Calculate(0))
	decimalEquals(t, math.MaxInt, fractal.Calculate(4))
	decimalEquals(t, ts.Candles[7].MaxPrice.Float(), fractal.Calculate(9))
	decimalEquals(t, ts.Candles[7].MaxPrice.Float(), fractal.Calculate(11))
	decimalEquals(t, ts.Candles[7].MaxPrice.Float(), fractal.Calculate(16))
	decimalEquals(t, ts.Candles[15].MaxPrice.Float(), fractal.Calculate(17))
}

func TestDownFractalIndicator_Calculate(t *testing.T) {
	const window = 2
	ts := generateTestTimeSeries()
	fractal := NewDownFractalIndicator(ts, window)

	decimalEquals(t, -math.MaxInt, fractal.Calculate(0))
	decimalEquals(t, -math.MaxInt, fractal.Calculate(4))
	decimalEquals(t, ts.Candles[9].MinPrice.Float(), fractal.Calculate(11))
	decimalEquals(t, ts.Candles[9].MinPrice.Float(), fractal.Calculate(13))
	decimalEquals(t, ts.Candles[12].MinPrice.Float(), fractal.Calculate(14))
}
