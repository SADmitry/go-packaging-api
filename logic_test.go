package main

import (
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	packSizes = []int{5000, 2000, 1000, 500, 250}

	tests := []struct {
		items    int
		expected map[int]int
	}{
		{12001, map[int]int{5000: 2, 2000: 1, 250: 1}},
		{500, map[int]int{500: 1}},
		{375, map[int]int{250: 2}},
		{0, map[int]int{}},
	}

	for _, tt := range tests {
		result := calculatePacks(tt.items)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("For %d items, expected %v, got %v", tt.items, tt.expected, result)
		}
	}
}
