package leet

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l2928(t *testing.T) {
	assert.Equal(t, 3, distributeCandies(5, 2))
	assert.Equal(t, 10, distributeCandies(3, 3))
}

func distributeCandies(n int, limit int) int {
	num := 0
	for i := 0; i <= limit && i <= n; i++ {
		for j := 0; j <= limit && j <= n-i; j++ {
			if n-i-j <= limit {
				num++
				fmt.Println(i, j, n-i-j)
			}
		}
	}
	return num
}

func distributeCandies1(n int, limit int) int {
	ans := 0
	for i := 0; i <= min(limit, n); i++ {
		if n-i > 2*limit {
			continue
		}
		ans += min(n-i, limit) - max(0, n-i-limit) + 1
	}
	return ans
}

func cal(x int) int {
	if x < 0 {
		return 0
	}
	return x * (x - 1) / 2
}

// 容斥:采用常规计数的思想，用所有方案减去不合法的方案。使用组合数学的容斥原理，用所有的方案数减去至少有一个小朋友分得超过 limit\textit{limit}limit 颗糖果。
func distributeCandies2(n int, limit int) int {
	return cal(n+2) - 3*cal(n-limit+1) + 3*cal(n-(limit+1)*2+2) - cal(n-3*(limit+1)+2)
}
