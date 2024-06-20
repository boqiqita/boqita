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
	assert.False(t, isMatch44("aa", "a"))
	assert.True(t, isMatch44("aa", "*"))
	assert.False(t, isMatch44("cb", "?a"))
	assert.True(t, isMatch44("adceb", "*a*b"))
	assert.True(t, isMatch44("", "******"))
	assert.False(t, isMatch44("bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab", "b*b*ab**ba*b**b***bba"))
}

// 动态规划、贪心
func isMatch44(s string, p string) bool {
	slen, plen := len(s), len(p)
	dp := make([][]bool, slen+1)
	for i := 0; i <= slen; i++ {
		dp[i] = make([]bool, plen+1)
	}

	dp[0][0] = true
	for i := 1; i <= plen; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		}
	}

	for i := 1; i <= slen; i++ {
		for j := 1; j <= plen; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[slen][plen]
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

// x的n次幂
func Test_l50(t *testing.T) {
	assert.Equal(t, 1024.00000, myPow(2.00000, 10))
	assert.Equal(t, 1024.00000, myPow(-2.00000, 10))
	assert.Equal(t, -32.00000, myPow(-2.00000, 5))
	assert.Equal(t, 9.261000000000001, myPow(2.10000, 3))
	assert.Equal(t, 0.25000, myPow(2.00000, -2))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		x, n = 1/x, -n
	}

	num := float64(1)
	for n > 1 {
		temp := n / 2
		if temp*2 != n {
			num *= x
		}
		x, n = x*x, temp
	}
	return x * num
}
