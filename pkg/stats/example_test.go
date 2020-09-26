package stats

import (
	"fmt"

	"github.com/gholib/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000_00,
			Category: "category 1",
		},
		{
			ID:       2,
			Amount:   50_000_00,
			Category: "category 2",
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "category 3",
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
		},
		{
			ID:       2,
			Amount:   50_000_00,
			Category: "category 2",
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "category 3",
		},
		{
			ID:       1,
			Amount:   20_000_00,
			Category: "category 1",
		},
		{
			ID:       1,
			Amount:   20_000_00,
			Category: "category 1",
		},
	}

	fmt.Println(TotalInCategory(payments, "category 1"))

	// Output:
	// 5000000

}
