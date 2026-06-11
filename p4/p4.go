package p4

import "slices"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var v []int

	v = append(v, nums1...)
	v = append(v, nums2...)

	slices.Sort(v)

	n := len(v)

	if n%2 == 1 {
		return float64(v[n/2])
	} else {
		a := float64(v[n/2-1])
		b := float64(v[n/2])

		return (a + b) / 2
	}
}
