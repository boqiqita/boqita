package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l83(t *testing.T) {
	ret := deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{2, nil}}})
	assert.Equal(t, 1, ret.Val)
	assert.Equal(t, 2, ret.Next.Val)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, p := head, head.Next
	for p != nil {
		if head.Val != p.Val {
			head.Next, head = p, p
		}
		p = p.Next
	}
	head.Next = nil
	return pre
}
