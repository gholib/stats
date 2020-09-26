package stats

import (
	"github.com/gholib/bank/v2/pkg/types"
)

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

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	periodsDynamic := map[types.Category]types.Money{}

	if len(first) == 0 && len(second) == 0 {
		return periodsDynamic
	}

	currentPeriods := first
	if len(first) < len(second) {
		currentPeriods = second
	}

	for key := range currentPeriods {
		periodsDynamic[key] = second[key] - first[key]
	}

	return periodsDynamic
}
