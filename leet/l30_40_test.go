package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l34(t *testing.T) {
	assert.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 4))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 11))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{}, 0))
}

func searchRange(nums []int, target int) []int {
	length := len(nums)
	l, r, mid := 0, length-1, -1
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			l, r = mid-1, mid+1
			for l > -1 && nums[l] == target {
				l--
			}
			for r < length && nums[r] == target {
				r++
			}
			return []int{l + 1, r - 1}
		} else if l == mid || r == mid {
			return []int{-1, -1}
		} else if nums[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}

	return []int{-1, -1}
}

func Test_l35(t *testing.T) {
	assert.Equal(t, 3, searchInsert([]int{1, 3, 4, 5, 6, 8, 10}, 5))
	assert.Equal(t, 2, searchInsert([]int{1, 4, 6, 7, 8, 9}, 6))
	assert.Equal(t, 1, searchInsert([]int{1, 2, 3, 4, 5, 10}, 2))
	assert.Equal(t, 1, searchInsert([]int{1, 3}, 3))
	assert.Equal(t, 1, searchInsert([]int{1, 3}, 2))
	assert.Equal(t, 2, searchInsert([]int{1, 3, 5, 6}, 5))
	assert.Equal(t, 1, searchInsert([]int{1, 3, 5, 6}, 2))
	assert.Equal(t, 4, searchInsert([]int{1, 3, 5, 6}, 7))
	assert.Equal(t, 2, searchInsert([]int{1, 2, 4, 6, 7}, 3))
}

func searchInsert(nums []int, target int) int {
	length := len(nums)
	if length == 0 || target <= nums[0] {
		return 0
	} else if target == nums[length-1] {
		return length - 1
	} else if target > nums[length-1] {
		return length
	}
	l, r := 0, length-1
	for l < r {
		mid := l + (r-l)/2 + 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if nums[l] >= target {
		return l
	}
	return l + 1
}
