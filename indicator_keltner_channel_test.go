package techan

import (
	"testing"
)

func TestKeltnerChannel(t *testing.T) {
	t.Run("Upper", func(t *testing.T) {
		upper := NewKeltnerChannelUpperIndicator(mockedTimeSeries, 3)

		expectedValues := []float64{
			0,
			0,
			0,
			67.0211,
			67.1374,
			67.0649,
			67.4216,
			67.5919,
			67.2417,
			67.2863,
			66.9885,
			66.3561,
		}

		indicatorEquals(t, expectedValues, upper)
	})

	t.Run("Lower", func(t *testing.T) {
		lower := NewKeltnerChannelLowerIndicator(mockedTimeSeries, 3)

		expectedValues := []float64{
			0,
			0,
			0,
			60.7989,
			60.3226,
			59.8551,
			59.9484,
			59.9431,
			59.4758,
			59.4424,
			57.7059,
			57.5011,
		}

		indicatorEquals(t, expectedValues, lower)
	})
}
