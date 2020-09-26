package stats

import (
	"fmt"

	"github.com/gholib/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000_00,
			Category: "category 1",
			Status:   "OK",
		},
		{
			ID:       2,
			Amount:   50_000_00,
			Category: "category 2",
			Status:   "OK",
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "category 3",
			Status:   "FAIL",
		},
	}

	fmt.Println(Avg(payments))

	//Output:
	// 3000000

}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000_00,
			Category: "category 1",
			Status:   "OK",
		},
		{
			ID:       2,
			Amount:   50_000_00,
			Category: "category 2",
			Status:   "OK",
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "category 3",
			Status:   "OK",
		},
		{
			ID:       1,
			Amount:   20_000_00,
			Category: "category 1",
			Status:   "OK",
		},
		{
			ID:       1,
			Amount:   20_000_00,
			Category: "category 1",
			Status:   "FAIL",
		},
	}

	fmt.Println(TotalInCategory(payments, "category 1"))

	// Output:
	// 3000000

}
