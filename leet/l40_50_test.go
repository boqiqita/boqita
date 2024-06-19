package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l42(t *testing.T) {
}

//
//func trap(height []int) int {
//
//}

func Test_l44(t *testing.T) {
	assert.True(t, isMatch44("aa", "a"))
	assert.True(t, isMatch44("aa", "*"))
	assert.False(t, isMatch44("cb", "?a"))
}

func isMatch44(s string, p string) bool {
	//slen, plen, temp, sub := len(s), len(p), '', -1
	//for i := 0; i < slen; i++ {
	//
	//}
	return false
}

func Test_l48(t *testing.T) {
	rotate([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}})
	rotate([][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16}})
}

func rotate(matrix [][]int) {
	length := len(matrix)
	n, pair := length-1, (length+1)/2
	for i := 0; i < pair; i++ {
		length = n - i
		for j := i; j < length; j++ {
			matrix[i][j], matrix[j][n-i], matrix[n-i][n-j], matrix[n-j][i] = matrix[n-j][i], matrix[i][j], matrix[j][n-i], matrix[n-i][n-j]
		}
	}
}
