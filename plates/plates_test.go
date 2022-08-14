package plates

import (
	"reflect"
	"testing"
)

func TestA(t *testing.T) {
	type test struct {
		input    float64
		expected []float64
	}
	tests := []test{
		{20, []float64{20}},
		{42, []float64{25, 15, 1, 1}},
	}

	for _, candidate := range tests {
		res := GetPlatesForWeight(candidate.input)
		if !reflect.DeepEqual(res, candidate.expected) {
			t.Fatalf(`Wanted %v got %v`, candidate.expected, res)
		}
	}
}

func TestB(t *testing.T) {
	type test struct {
		input    float64
		expected []float64
	}
	tests := []test{
		{25, []float64{2.5}},
		{45, []float64{10, 2.5}},
		{55, []float64{15, 2.5}},
		{60, []float64{20}},
	}

	for _, candidate := range tests {
		res := GetPlatesForBar(candidate.input)
		if !reflect.DeepEqual(res, candidate.expected) {
			t.Fatalf(`Wanted %v got %v`, candidate.expected, res)
		}
	}
}
