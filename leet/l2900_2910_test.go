package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l2903(t *testing.T) {
	ret := findIndices([]int{5, 1, 4, 1}, 2, 4)
	assert.Equal(t, 0, ret[0])
	assert.Equal(t, 3, ret[1])
}

var ret = []int{-1, -1}

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	length, temp := len(nums), 0
	for j, i := 0, indexDifference; j < length && i < length; j, i = j+1, j+1+indexDifference {
		for ; i < length; i++ {
			temp = nums[j] - nums[i]
			if temp >= valueDifference || temp <= -valueDifference {
				return []int{j, i}
			}
		}
	}
	return ret
}
