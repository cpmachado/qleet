package p2

import (
	"fmt"
	"testing"
)

type addTwoNumbersTest struct {
	L1   *ListNode
	L2   *ListNode
	Want *ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []addTwoNumbersTest{
		{
			L1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			L2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			Want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 8,
					},
				},
			},
		},
		{
			L1: &ListNode{
				Val: 0,
			},
			L2: &ListNode{
				Val: 0,
			},
			Want: &ListNode{
				Val: 0,
			},
		},
		{
			L1: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 9,
									},
								},
							},
						},
					},
				},
			},
			L2: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
						},
					},
				},
			},
			Want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 1,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, s := range tests {
		description := fmt.Sprintf("l1: %s, l2: %s, target: %s", s.L1, s.L2, s.Want)
		t.Run(description, func(t *testing.T) {
			got := AddTwoNumbers(s.L1, s.L2)

			if s.Want.String() != got.String() {
				t.Errorf("Expected %v, but got %v\n", s.Want, got)
			}
		})
	}
}
