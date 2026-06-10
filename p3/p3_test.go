package p3

import (
	"fmt"
	"testing"
)

type lengthOfLongestSubstringTest struct {
	S    string
	Want int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []lengthOfLongestSubstringTest{
		{S: "abcabcbb", Want: 3},
		{S: "bbbbb", Want: 1},
		{S: "pwwkew", Want: 5},
	}

	for _, s := range tests {
		description := fmt.Sprintf("s: %s, target: %d", s.S, s.Want)
		t.Run(description, func(t *testing.T) {
			got := LengthOfLongestSubstring(s.S)

			if s.Want != got {
				t.Errorf("Expected %v, but got %v\n", s.Want, got)
			}
		})
	}
}
