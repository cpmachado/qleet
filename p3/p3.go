package p3

func LengthOfLongestSubstring(s string) int {
	maxCount := 0
	counter := 0
	seen := make(map[rune]int)

	for idx, c := range s {
		if bidx, found := seen[c]; found {
			if maxCount < counter {
				maxCount = counter
			}
			for k, v := range seen {
				if v < bidx {
					delete(seen, k)
				}
			}
			counter = idx - bidx - 1
		}
		counter++
		seen[c] = idx
	}
	if maxCount < counter {
		maxCount = counter
	}
	return maxCount
}
