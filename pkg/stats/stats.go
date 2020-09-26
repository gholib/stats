package stats

import "github.com/gholib/bank/v2/pkg/types"

// Avg average payment amount
func Avg(payments []types.Payment) types.Money {
	if len(payments) < 0 {
		return 0
	}
	var summa, count types.Money
	for _, payment := range payments {
		if payment.Amount < 0 {
			continue
		}
		if payment.Status == types.StatusFail {
			continue
		}
		summa += payment.Amount
		count++
	}

	result := summa / count

	return result
}

//TotalInCategory amount of purchases by category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	if len(payments) < 0 {
		return 0
	}
	var summa types.Money
	for _, payment := range payments {
		if payment.Category == category {
			if payment.Amount < 0 {
				continue
			}
			if payment.Status == types.StatusFail {
				continue
			}
			summa += payment.Amount
		}
	}

	return summa
}
