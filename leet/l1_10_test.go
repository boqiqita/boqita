package leet

import (
	"bytes"
	"fmt"
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

func Test_l4(t *testing.T) {
	//assert.Equal(t, float64(2), findMedianSortedArrays([]int{1, 3}, []int{2}))
	//assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, 2.5, findMedianSortedArrays([]int{2, 4}, []int{1, 3}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1, length2, m, n := len(nums1), len(nums2), 0, 0
	for a1, a2, b1, b2 := 0, length1-1, 0, length2-1; a1 <= a2 || b1 <= b2; {
		if a1 <= a2 && b1 <= b2 {
			if nums1[a1] <= nums2[b1] {
				if nums1[a2] >= nums2[b2] {
					m, n, a1, a2 = nums1[a1], nums1[a2], a1+1, a2-1
				} else {
					m, n, a1, b2 = nums1[a1], nums2[b2], a1+1, b2-1
				}
			} else {
				if nums1[a2] <= nums2[b2] {
					m, n, b1, b2 = nums2[b1], nums2[b2], b1+1, b2-1
				} else {
					m, n, b1, a2 = nums2[b1], nums1[a2], b1+1, a2-1
				}
			}
		} else if a1 <= a2 {
			m, n, a1, a2 = nums1[a1], nums1[a2], a1+1, a2-1
		} else {
			m, n, b1, b2 = nums2[b1], nums2[b2], b1+1, b2-1
		}
	}
	return float64(m+n) / 2
}

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

func Test_l7(t *testing.T) {
	assert.Equal(t, 321, reverse(123))
	assert.Equal(t, -321, reverse(-123))
	assert.Equal(t, 21, reverse(120))
	assert.Equal(t, 0, reverse(1534236469))
}

func reverse(x int) int {
	if x > -10 && x < 10 {
		return x
	}

	m, y := 1, 0
	if x < 0 {
		m, x = -1, -x
	}

	for x != 0 {
		y, x = y*10+x%10, x/10
	}
	x = m * y
	if x <= -2147483648 || x >= 2147483647 {
		return 0
	}
	return x
}

func Test_l8(t *testing.T) {
	assert.Equal(t, 42, myAtoi("42"))
	assert.Equal(t, -42, myAtoi(" -042"))
	assert.Equal(t, 1337, myAtoi("1337c0d3"))
	assert.Equal(t, 0, myAtoi("words and 987"))
}

func myAtoi(s string) int {
	i, length, b, num := 0, len(s), true, 0
	for ; i < length; i++ {
		if s[i] != ' ' {
			if s[i] == '-' {
				b = false
				i++
			} else if s[i] == '+' {
				i++
			}
			break
		}
	}

	for ; i < length; i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		num = num*10 + int(s[i]-'0')
		if b && num > 2147483647 {
			return 2147483647
		} else if !b && num > 2147483648 {
			return -2147483648
		}
	}
	if !b {
		return -num
	}
	return num
}

func Test_l9(t *testing.T) {
	assert.True(t, isPalindrome(121))
	assert.False(t, isPalindrome(-121))
	assert.False(t, isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 10 && x >= 0 {
		return true
	}
	str := fmt.Sprintf("%d", x)
	for l, r := 0, len(str)-1; l < r; l, r = l+1, r-1 {
		if str[l] != str[r] {
			return false
		}
	}
	return true
}

func Test_l10(t *testing.T) {
	assert.True(t, isMatch10("aa", "a"))
	assert.True(t, isMatch10("aa", "a*"))
	assert.False(t, isMatch10("ab", ".*"))
}

func isMatch10(s string, p string) bool {
	return false
}
