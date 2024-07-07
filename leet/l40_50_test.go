package leet

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
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
func Test_l45(t *testing.T) {
	jsonData, _ := json.Marshal(readBinaryWatch(4))
	fmt.Println(string(jsonData))
	jsonData, _ = json.Marshal(readBinaryWatch(5))
	fmt.Println(string(jsonData))
}

func readBinaryWatch(turnedOn int) []string {
	hour, minute, ans := map[int][]string{}, map[int][]string{}, []string{}
	for i := 0; i < 12; i++ {
		count, temp := 0, i
		for temp > 0 {
			if temp&1 == 1 {
				count++
			}
			temp = temp >> 1
		}
		arr := hour[count]
		if arr == nil {
			arr = []string{}
		}
		arr = append(arr, fmt.Sprintf("%d", i))
		hour[count] = arr
	}

	for i := 0; i < 60; i++ {
		count, temp := 0, i
		for temp > 0 {
			if temp&1 == 1 {
				count++
			}
			temp = temp >> 1
		}
		arr := minute[count]
		if arr == nil {
			arr = []string{}
		}
		arr = append(arr, fmt.Sprintf("%02d", i))
		minute[count] = arr
	}

	for i := 0; i <= turnedOn; i++ {
		h, ok1 := hour[i]
		m, ok2 := minute[turnedOn-i]
		if ok1 && ok2 {
			for _, v1 := range h {
				for _, v2 := range m {
					ans = append(ans, fmt.Sprintf("%s:%s", v1, v2))
				}
			}
		}
	}

	return ans
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

func Test_l49(t *testing.T) {
	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		arr := []rune(str)
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
		strKey := string(arr)

		if arr, ok := mp[strKey]; !ok {
			mp[strKey] = []string{str}
		} else {
			mp[strKey] = append(arr, str)
		}
	}

	ans := [][]string{}
	for _, arr := range mp {
		ans = append(ans, arr)
	}
	return ans
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
