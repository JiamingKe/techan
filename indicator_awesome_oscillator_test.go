package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAwesomeOscillatorIndicator_Calculate(t *testing.T) {

	series := generateTestTimeSeries()

	awesomeOscillator := NewAwesomeOscillatorIndicator(series).Calculate(40)

	assert.EqualValues(t, "-2.4194", awesomeOscillator.FormattedString(4))
}
