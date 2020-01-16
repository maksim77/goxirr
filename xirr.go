package goxirr

import (
	"math"
	"time"
)

type Transaction struct {
	Date time.Time
	Cash float64
}

func Xirr(transactions []Transaction) float64 {
	var years []float64
	for _, ta := range transactions {
		years = append(years, (ta.Date.Sub(transactions[0].Date).Hours()/24)/365)
	}
	residual := 1.0
	step := 0.05
	guess := 0.05
	epsilon := 0.0001
	limit := 10000

	for math.Abs(residual) > epsilon && limit > 0 {
		limit -= 1
		residual = 0.0
		for i, trans := range transactions {
			residual += trans.Cash / math.Pow(guess, years[i])
		}
		if math.Abs(residual) > epsilon {
			if residual > 0 {
				guess += step
			} else {
				guess -= step
				step /= 2.0
			}
		}
	}
	return (guess - 1) * 100
}
