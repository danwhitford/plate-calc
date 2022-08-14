package main

import (
	"reflect"
	"testing"
)

func Test20(t *testing.T) {
	type test struct {
		input    float32
		expected []float32
	}
	tests := []test{
		{20, []float32{20}},
		{42, []float32{25, 15, 1, 1}},
	}

	for _, candidate := range tests {
		res := GetPlatesForWeight(candidate.input)
		if !reflect.DeepEqual(res, candidate.expected) {
			t.Fatalf(`Wanted %v got %v`, candidate.expected, res)
		}
	}
}
