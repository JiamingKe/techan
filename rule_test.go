package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type truthRule struct{}

func (tr truthRule) IsSatisfied(index int, record *TradingRecord) bool {
	return true
}

type falseRule struct{}

func (fr falseRule) IsSatisfied(index int, record *TradingRecord) bool {
	return false
}

func TestAndRule(t *testing.T) {
	t.Run("both truthy", func(t *testing.T) {
		rule := And(truthRule{}, truthRule{})

		assert.True(t, rule.IsSatisfied(0, nil))
	})

	t.Run("both falsey", func(t *testing.T) {
		rule := And(falseRule{}, falseRule{})

		assert.False(t, rule.IsSatisfied(0, nil))
	})

	t.Run("one of each", func(t *testing.T) {
		rule := And(truthRule{}, falseRule{})

		assert.False(t, rule.IsSatisfied(0, nil))
	})

	t.Run("two of each", func(t *testing.T) {
		rule := And(truthRule{}, truthRule{}, falseRule{}, falseRule{})

		assert.False(t, rule.IsSatisfied(0, nil))
	})
}

func TestOrRule(t *testing.T) {
	t.Run("both truthy", func(t *testing.T) {
		rule := Or(truthRule{}, truthRule{})

		assert.True(t, rule.IsSatisfied(0, nil))
	})

	t.Run("both falsey", func(t *testing.T) {
		rule := Or(falseRule{}, falseRule{})

		assert.False(t, rule.IsSatisfied(0, nil))
	})

	t.Run("one of each", func(t *testing.T) {
		rule := Or(truthRule{}, falseRule{})

		assert.True(t, rule.IsSatisfied(0, nil))
	})

	t.Run("two of each", func(t *testing.T) {
		rule := Or(truthRule{}, truthRule{}, falseRule{}, falseRule{})

		assert.True(t, rule.IsSatisfied(0, nil))
	})
}

func TestUnderIndicatorRule(t *testing.T) {
	zeroIndicator := NewConstantIndicator(0)
	oneIndicator := NewConstantIndicator(1)
	twoIndicator := NewConstantIndicator(2)

	t.Run("returns true when first indicator is under second indicator", func(t *testing.T) {
		rule := Under(zeroIndicator, oneIndicator, twoIndicator)

		assert.True(t, rule.IsSatisfied(0, nil))
	})

	t.Run("returns false when first indicator is over second indicator", func(t *testing.T) {
		rule := Under(oneIndicator, zeroIndicator)

		assert.False(t, rule.IsSatisfied(0, nil))
	})
}

func TestPercentChangeRule(t *testing.T) {
	record := NewTradingRecord()

	t.Run("returns false when percent change is less than the amount", func(t *testing.T) {
		series := mockTimeSeries("1", "1.1")
		rule := NewPercentChangeRule(NewClosePriceIndicator(series), 0.25)

		assert.False(t, rule.IsSatisfied(1, record))
	})

	t.Run("returns true when percent change is greater than the amount", func(t *testing.T) {
		series := mockTimeSeries("1", "1.11")
		rule := NewPercentChangeRule(NewClosePriceIndicator(series), 0.1)

		assert.True(t, rule.IsSatisfied(1, record))
	})
}
