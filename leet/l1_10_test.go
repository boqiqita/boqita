package leet

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l1(t *testing.T) {
	ret := twoSum([]int{2, 7, 11, 15}, 9)
	assert.Equal(t, 0, ret[0])
	assert.Equal(t, 1, ret[1])
}

func twoSum(nums []int, target int) []int {
	tempMp := make(map[int]int, len(nums)/3)
	for i, num := range nums {
		if j, ok := tempMp[target-num]; ok {
			return []int{j, i}
		}
		tempMp[num] = i
	}
	return nil
}

func Test_l2(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	ret := addTwoNumbers(l1, l2)
	assert.Equal(t, 7, ret.Val)
	assert.Equal(t, 0, ret.Next.Val)
	assert.Equal(t, 8, ret.Next.Next.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre, tag := &ListNode{}, false
	p := pre
	for l1 != nil || l2 != nil {
		if l1 == nil {
			if tag {
				if l2.Val < 9 {
					l2.Val++
					tag = false
				} else {
					l2.Val = 0
				}
			}
			p.Next, p, l2 = l2, l2, l2.Next
		} else if l2 == nil {
			if tag {
				if l1.Val < 9 {
					l1.Val++
					tag = false
				} else {
					l1.Val = 0
				}
			}
			p.Next, p, l1 = l1, l1, l1.Next
		} else {
			l1.Val = l1.Val + l2.Val
			if tag {
				l1.Val++
			}
			if l1.Val > 9 {
				tag = true
				l1.Val -= 10
			} else {
				tag = false
			}
			p.Next, p, l1, l2 = l1, l1, l1.Next, l2.Next
		}
	}

	if tag {
		p.Next = &ListNode{1, nil}
	}

	return pre.Next
}

func Test_l3(t *testing.T) {
	ret := lengthOfLongestSubstring("abcabcbb")
	assert.Equal(t, 3, ret)
}

func lengthOfLongestSubstring(s string) int {
	l, r, max := 0, 0, 0
	for i := range s {
		r++
		for j := l; j < r && j < i; j++ {
			if s[j] == s[i] {
				l = j + 1
				break
			}
		}
		temp := r - l
		if max < temp {
			max = temp
		}
	}
	return max
}

//func Test_l4(t *testing.T) {
//	ret := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
//	assert.Equal(t, 2.5, ret)
//}
//
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	len1, len2 := len(nums1), len(nums2)
//	if len1 > 0 && len2 == 0 {
//		pair := len1 / 2
//		if pair*2 == len1 {
//			return (float64)(nums1[pair]+nums1[pair+1]) / 2
//		}
//		return (float64)(nums1[pair+1])
//	}
//	if len1 == 0 && len2 > 0 {
//		pair := len2 / 2
//		if pair*2 == len2 {
//			return (float64)(nums2[pair]+nums2[pair+1]) / 2
//		}
//		return (float64)(nums2[pair+1])
//	}
//
//	m1, m2 := nums1[len1/2], nums2[len2/2]
//
//}

func Test_l5(t *testing.T) {
	ret := longestPalindrome("babad")
	assert.Equal(t, "bab", ret)
	ret = longestPalindrome("cbbd")
	assert.Equal(t, "bb", ret)

	ret = longestPalindrome("bbb")
	assert.Equal(t, "bbb", ret)
}

func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	l, r, max, q, p, b, temp := 0, 1, 1, 0, 0, false, 0
	for mid := 1; mid < length; mid++ {
		q, p, b = mid-1, mid+1, false
		for ; q >= 0 && p < length && s[q] == s[p]; q, p = q-1, p+1 {
			b = true
		}
		if b {
			q, p = q+1, p
			temp = p - q
			if max < temp {
				l, r, max = q, p, temp
			}
		}

		if s[mid-1] == s[mid] {
			q, p, b = mid-1, mid, false
			for ; q >= 0 && p < length && s[q] == s[p]; q, p = q-1, p+1 {
				b = true
			}
			if b {
				q, p = q+1, p
				temp = p - q
				if max < temp {
					l, r, max = q, p, temp
				}
			}
		}
	}
	return s[l:r]
}

func Test_l6(t *testing.T) {
	ret := convert("PAYPALISHIRING", 3)
	assert.Equal(t, "PAHNAPLSIIGYIR", ret)
	ret = convert("PAYPALISHIRING", 4)
	assert.Equal(t, "PINALSIGYAHRPI", ret)
	ret = convert("A", 1)
	assert.Equal(t, "A", ret)
	ret = convert("AB", 1)
	assert.Equal(t, "AB", ret)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	num, length := numRows-1, len(s)
	step := num * 2
	var ret bytes.Buffer
	ret.Grow(length)
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < length; j += step {
			ret.WriteRune(rune(s[j+i]))
			if i != 0 && i != num {
				temp := j + step - i
				if temp < length {
					ret.WriteRune(rune(s[temp]))
				}
			}
		}
	}
	return ret.String()
}

//func Test_l7(t *testing.T) {
//
//}
//
//func reverse(x int) int {
//
//}
