package stats

import "github.com/gholib/bank/v2/pkg/types"

// Avg average payment amount
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	paymentByCategories := map[types.Category]types.Money{}

	if len(payments) < 0 {
		return paymentByCategories
	}

	counts := map[types.Category]types.Money{}

	for _, payment := range payments {
		paymentByCategories[payment.Category] += payment.Amount
		counts[payment.Category]++
	}

	for key, value := range paymentByCategories {
		paymentByCategories[key] = value / counts[key]
	}

	return paymentByCategories
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
