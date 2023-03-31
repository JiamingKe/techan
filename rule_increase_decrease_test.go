package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreaseRule(t *testing.T) {
	const window = 3

	t.Run("returns false when index < window - 1", func(t *testing.T) {
		rule := NewIncreaseRule(NewClosePriceIndicator(mockTimeSeries()), window)

		assert.False(t, rule.IsSatisfied(window-2, nil))
	})

	t.Run("returns true when increase", func(t *testing.T) {
		series := mockTimeSeries("1", "2", "3")
		rule := NewIncreaseRule(NewClosePriceIndicator(series), window)

		assert.True(t, rule.IsSatisfied(window-1, nil))
	})

	t.Run("returns false when part of it same", func(t *testing.T) {
		series := mockTimeSeries("1", "2", "2")
		rule := NewIncreaseRule(NewClosePriceIndicator(series), window)

		assert.False(t, rule.IsSatisfied(window-1, nil))
	})

	t.Run("returns false when part of it decrease", func(t *testing.T) {
		series := mockTimeSeries("1", "2", "1")
		rule := NewIncreaseRule(NewClosePriceIndicator(series), window)

		assert.False(t, rule.IsSatisfied(window-1, nil))
	})
}

func TestDecreaseRule(t *testing.T) {
	const window = 3

	t.Run("returns false when index < window - 1", func(t *testing.T) {
		rule := NewDecreaseRule(NewClosePriceIndicator(mockTimeSeries()), window)

		assert.False(t, rule.IsSatisfied(window-2, nil))
	})

	t.Run("returns true when decrease", func(t *testing.T) {
		series := mockTimeSeries("3", "2", "1")
		rule := NewDecreaseRule(NewClosePriceIndicator(series), window)

		assert.True(t, rule.IsSatisfied(window-1, nil))
	})

	t.Run("returns false when part of it same", func(t *testing.T) {
		series := mockTimeSeries("2", "2", "1")
		rule := NewDecreaseRule(NewClosePriceIndicator(series), window)

		assert.False(t, rule.IsSatisfied(window-1, nil))
	})

	t.Run("returns false when part of it increase", func(t *testing.T) {
		series := mockTimeSeries("1", "2", "1")
		rule := NewDecreaseRule(NewClosePriceIndicator(series), window)

		assert.False(t, rule.IsSatisfied(1, nil))
	})

}
