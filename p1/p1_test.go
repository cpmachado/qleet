package p1

import (
	"fmt"
	"reflect"
	"testing"
)

type twoSumTest struct {
	Nums   []int
	Target int
	Want   []int
}

func TestTwoSum(t *testing.T) {
	tests := []twoSumTest{
		{Nums: []int{2, 7, 11, 15}, Target: 9, Want: []int{0, 1}},
		{Nums: []int{3, 2, 4}, Target: 6, Want: []int{1, 2}},
		{Nums: []int{3, 3}, Target: 6, Want: []int{0, 1}},
	}

	for _, s := range tests {
		description := fmt.Sprintf("nums: %v, target: %v", s.Nums, s.Target)
		t.Run(description, func(t *testing.T) {
			got := TwoSum(s.Nums, s.Target)

			if !reflect.DeepEqual(s.Want, got) {
				t.Errorf("expected = %v, target = %d, twoSum(%v) = %v", s.Want, s.Target, s.Nums, got)
			}
		})
	}
}
