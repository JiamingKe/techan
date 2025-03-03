package techan

// NewCrossUpIndicatorRule returns a new rule that is satisfied when the lower indicator has crossed above the upper indicator.
func NewCrossUpIndicatorRule(upper, lower Indicator, window int) Rule {
	return crossRule{
		upper:  upper,
		lower:  lower,
		cmp:    1,
		window: window,
	}
}

// NewCrossDownIndicatorRule returns a new rule that is satisfied when the upper indicator has crossed below the lower indicator.
func NewCrossDownIndicatorRule(upper, lower Indicator, window int) Rule {
	return crossRule{
		upper:  lower,
		lower:  upper,
		cmp:    -1,
		window: window,
	}
}

type crossRule struct {
	upper  Indicator
	lower  Indicator
	cmp    int
	window int
}

func (cr crossRule) IsSatisfied(index int, record *TradingRecord) bool {
	i := index
	first := Max(index-cr.window+1, 0)

	if i == 0 {
		return false
	}

	if cmp := cr.lower.Calculate(i).Cmp(cr.upper.Calculate(i)); cmp == 0 || cmp == cr.cmp {
		for ; i >= first; i-- {
			if cmp = cr.lower.Calculate(i).Cmp(cr.upper.Calculate(i)); cmp == 0 || cmp == -cr.cmp {
				return true
			}
		}
	}

	return false
}
