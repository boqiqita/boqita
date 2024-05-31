package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l35(t *testing.T) {
	assert.Equal(t, 2, searchInsert([]int{1, 3, 5, 6}, 5))
	assert.Equal(t, 1, searchInsert([]int{1, 3, 5, 6}, 2))
	assert.Equal(t, 4, searchInsert([]int{1, 3, 5, 6}, 7))
}

func searchInsert(nums []int, target int) int {
	num, length := 0, len(nums)
	if length == 0 || target <= nums[0] {
		return 0
	} else if target >= nums[length-1] {
		return length
	}

	return num
}
