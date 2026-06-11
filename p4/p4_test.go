package p4

import (
	"fmt"
	"testing"
)

type aTest struct {
	Nums1 []int
	Nums2 []int
	Want  float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []aTest{
		{
			Nums1: []int{1, 3},
			Nums2: []int{2},
			Want:  2,
		},
	}

	for _, s := range tests {
		description := fmt.Sprintf("nums1: %v, nums2: %v, want: %v", s.Nums1, s.Nums2, s.Want)
		t.Run(description, func(t *testing.T) {
			got := FindMedianSortedArrays(s.Nums1, s.Nums2)

			if s.Want != got {
				t.Errorf("Expected %v, but got %v\n", s.Want, got)
			}
		})
	}
}
