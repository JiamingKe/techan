package techan

import "testing"

func TestAverageTrueRangeIndicator(t *testing.T) {
	atrIndicator := NewAverageTrueRangeIndicator(mockedTimeSeries, 3)

	expectedValues := []float64{
		0,
		0,
		1.3333,
		1.5556,
		1.7037,
		1.8025,
		1.8683,
		1.9122,
		1.9415,
		1.961,
		2.3207,
		2.2138,
	}

	indicatorEquals(t, expectedValues, atrIndicator)
}
