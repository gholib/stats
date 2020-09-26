package stats

import (
	"reflect"
	"testing"

	"github.com/gholib/bank/v2/pkg/types"
)

func TestCategoriesAvg_empty(t *testing.T) {
	var payments = []types.Payment{}
	result := CategoriesAvg(payments)

	if len(result) != 0 {
		t.Error("Invalid result")
	}
}

func TestCategoriesAvg_correct(t *testing.T) {
	var payments = []types.Payment{
		{
			ID:       1,
			Category: "auto",
			Amount:   1000_00,
		},
		{
			ID:       1,
			Category: "auto",
			Amount:   3000_00,
		},
		{
			ID:       1,
			Category: "food",
			Amount:   1000_00,
		},
		{
			ID:       1,
			Category: "fun",
			Amount:   1000_00,
		},
	}
	expected := map[types.Category]types.Money{
		"auto": 2000_00,
		"food": 1000_00,
		"fun":  1000_00,
	}
	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result:  %v", result)
	}
}
