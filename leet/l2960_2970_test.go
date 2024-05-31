package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l2965(t *testing.T) {
	ret := findMissingAndRepeatedValues([][]int{{1, 3}, {2, 2}})
	assert.Equal(t, 2, ret[0])
	assert.Equal(t, 4, ret[1])

	ret = findMissingAndRepeatedValues([][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}})
	assert.Equal(t, 9, ret[0])
	assert.Equal(t, 5, ret[1])

	ret = findMissingAndRepeatedValues([][]int{{30, 3, 17, 56, 48, 63, 45, 29}, {54, 58, 13, 16, 20, 18, 61, 26}, {5, 59, 35, 19, 15, 37, 51, 64}, {28, 52, 4, 43, 64, 2, 39, 21}, {22, 47, 31, 53, 8, 57, 36, 38}, {9, 41, 55, 44, 40, 49, 60, 33}, {50, 23, 7, 62, 11, 34, 24, 14}, {10, 46, 32, 42, 27, 12, 1, 6}})
	assert.Equal(t, 64, ret[0])
	assert.Equal(t, 25, ret[1])
}

func findMissingAndRepeatedValues(grid [][]int) []int {
	num, mp, repeated, length := 0, map[int]struct{}{}, 0, len(grid)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			num = num + i*length + j + 1 - grid[i][j]
			if repeated == 0 {
				if _, ok := mp[grid[i][j]]; ok {
					repeated = grid[i][j]
				} else {
					mp[grid[i][j]] = struct{}{}
				}
			}
		}
	}
	return []int{repeated, repeated + num}
}
