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
	var base, tail *ListNode
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		node := new(ListNode)
		node.Val = (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10

		if tail == nil {
			base = node
		} else {
			tail.Next = node
		}
		tail = node

	}
	return base
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
