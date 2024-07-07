package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l25(t *testing.T) {
	reverseKGroup(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}, 2)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	pre := &ListNode{Next: head}
	p, l, r, count := pre, pre, pre, 1
	for head != nil {
		head = head.Next //3
		if count == k {
			count, l, r, p.Next = 0, p.Next, p.Next, nil //0->nil,1
			for l != nil && l != head {
				temp := l.Next
				p.Next, l.Next = l, p.Next
				l = temp
			}
			p = r
		}
		count++
	}
	return pre.Next
}

func Test_l28(t *testing.T) {
	assert.Equal(t, 0, strStr("sadbutsad", "sad"))
	assert.Equal(t, -1, strStr("leetcode", "leeto"))
}

func strStr(haystack string, needle string) int {
	l1, l2 := len(haystack), len(needle)
	for i := 0; i < l1; i++ {
		if l1-i < l2 {
			return -1
		}
		b := true
		for j := 0; j < l2; j++ {
			if haystack[i+j] != needle[j] {
				b = false
				break
			}
		}
		if b {
			return i
		}
	}
	return -1
}
