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

func TestPeriodsDynamic_empty(t *testing.T) {
	var first = map[types.Category]types.Money{}
	var second = map[types.Category]types.Money{}
	result := PeriodsDynamic(first, second)

	if len(result) != 0 {
		t.Error("Invalid result")
	}
}
func TestPeriodsDynamic_correct(t *testing.T) {
	var first = map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	var second = map[types.Category]types.Money{
		"auto": 20,
		"food": 20,
	}

	expected := map[types.Category]types.Money{
		"auto": 10,
		"food": 0,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result:  %v", result)
	}
}

func TestPeriodsDynamic_lessSecondPeriod(t *testing.T) {
	var first = map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	var second = map[types.Category]types.Money{
		"food": 20,
	}

	expected := map[types.Category]types.Money{
		"auto": -10,
		"food": 0,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result:  %v", result)
	}
}

func TestPeriodsDynamic_lessFirstPeriod(t *testing.T) {
	var first = map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	var second = map[types.Category]types.Money{
		"auto":   10,
		"food":   25,
		"mobile": 5,
	}

	expected := map[types.Category]types.Money{
		"auto":   0,
		"food":   5,
		"mobile": 5,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result:  %v", result)
	}
}
