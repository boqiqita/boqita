package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l2951(t *testing.T) {
	assert.Equal(t, 0, len(findPeaks([]int{2, 4, 4})))
	assert.Equal(t, 2, len(findPeaks([]int{1, 4, 3, 8, 5})))
	assert.Equal(t, 0, len(findPeaks([]int{1, 1, 3})))
}

func findPeaks(mountain []int) []int {
	arr, length := []int{}, len(mountain)
	for i := 1; i < length-1; i++ {
		if (mountain[i-1] < mountain[i]) && (mountain[i] > mountain[i+1]) {
			arr = append(arr, i)
		}
	}
	return arr
}
