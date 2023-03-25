package techan

import "github.com/sdcoffey/big"

// NewVolumeIndicator returns an indicator which returns the volume of a candle for a given index
func NewFractalIndicator(series *TimeSeries, window int) Indicator {
	return fractalIndicator{
		highPrice: NewHighPriceIndicator(series),
		lowPrice:  NewLowPriceIndicator(series),
		window:    window,
	}
}

type fractalIndicator struct {
	highPrice Indicator
	lowPrice  Indicator
	window    int
}

func (fi fractalIndicator) Calculate(index int) big.Decimal {
	highestPrice := fi.highPrice.Calculate(index)
	lowestPrice := fi.lowPrice.Calculate(index)
	for i := index - fi.window; i <= index+fi.window; i++ {
		currentHigh := fi.highPrice.Calculate(i)
		currentLow := fi.lowPrice.Calculate(i)

		if currentHigh.GT(highestPrice) {
			highestPrice = currentHigh
		}

		if currentLow.LT(lowestPrice) {
			lowestPrice = currentLow
		}
	}

	if fi.highPrice.Calculate(index).EQ(highestPrice) {
		return big.ONE
	} else if fi.lowPrice.Calculate(index).EQ(lowestPrice) {
		return big.ONE.Neg()
	}

	return big.ZERO
}
