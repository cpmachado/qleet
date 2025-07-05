package p1

// TwoSum retrieves the indices of two numbers within nums, whose sum is target.
//
// Constraints
//   - 2 <= nums.length <= 10^4
//   - -10^9 <= nums[i] <= 10^9
//   - -10^9 <= target <= 10^9
//   - Only one valid answer exists.
func TwoSum(nums []int, target int) []int {
	// map[delta from target]index
	processed := make(map[int]int)

	for i, num := range nums {
		delta := target - num
		if j, found := processed[delta]; found {
			return []int{j, i}
		}
		// this is further constraint imposed by myself.
		// Consider elements are non-unique, and that you seek the lowest index
		if _, found := processed[num]; !found {
			processed[num] = i
		}
	}
	return nil
}
