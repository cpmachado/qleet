package p2

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}

func (l1 *ListNode) Equal(l2 *ListNode) bool {
	p, q := l1, l2

	for p != nil && q != nil && p.Val == q.Val {
		p = p.Next
		q = q.Next
	}

	return p != nil && q != nil
}

func (l1 *ListNode) String() string {
	s := &strings.Builder{}
	p := l1

	s.WriteString("[")

	for p != nil {
		fmt.Fprintf(s, "%d", p.Val)
		p = p.Next
		if p != nil {
			s.WriteString(" ")
		}
	}

	s.WriteString("]")

	return s.String()
}
