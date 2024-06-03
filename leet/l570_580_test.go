package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l575(t *testing.T) {
	assert.Equal(t, 3, distributeCandies575([]int{1, 1, 2, 2, 3, 3}))
	assert.Equal(t, 2, distributeCandies575([]int{1, 1, 2, 3}))
	assert.Equal(t, 1, distributeCandies575([]int{6, 6, 6, 6}))
}

func distributeCandies575(candyType []int) int {
	mp, length := make(map[int]int), len(candyType)/2
	for _, v := range candyType {
		mp[v]++
		if len(mp) > length {
			return length
		}
	}
	return len(mp)
}
