package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
