package trading

import "math"

type PositionCalculator struct {
	AccountBalance float64
	RiskPercentage float64
}

func NewPositionCalculator(accountBalance, riskPercentage float64) *PositionCalculator {
	return &PositionCalculator{
		AccountBalance: accountBalance,
		RiskPercentage: riskPercentage,
	}
}

func (pc *PositionCalculator) CalculatePositionSize(entryPrice, stopLoss float64) float64 {
	riskAmount := pc.AccountBalance * (pc.RiskPercentage / 100)
	stopLossDistance := math.Abs(entryPrice - stopLoss)
	positionSize := riskAmount / stopLossDistance
	return math.Floor(positionSize*100) / 100 // Round down to 2 decimal places
}
