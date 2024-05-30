package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l2981(t *testing.T) {
	assert.Equal(t, 2, maximumLength("aaaa"))
	assert.Equal(t, -1, maximumLength("abcdef"))
	assert.Equal(t, 1, maximumLength("abcaba"))
	assert.Equal(t, 1, maximumLength("abcaba"))
}

func Test_l2982(t *testing.T) {
	assert.Equal(t, 2, maximumLength("aaaa"))
	assert.Equal(t, -1, maximumLength("abcdef"))
	assert.Equal(t, 1, maximumLength("abcaba"))
	assert.Equal(t, 3, maximumLength("abcccccdddd"))
	assert.Equal(t, 2, maximumLength("ereerrrererrrererre"))
	assert.Equal(t, 3, maximumLength("akphhppppp"))
	assert.Equal(t, 3, maximumLength("dkkkdhhhhh"))
}

func maximumLength(s string) int {
	length := len(s)
	if length < 3 {
		return -1
	}

	sub, mp, count := 0, map[string]int{}, -1
	for i := 1; i <= length; i++ {
		if i == length || s[sub] != s[i] {
			temp := s[sub:i]
			tempLength := len(temp)
			if count < tempLength {
				for j := tempLength; j > count && j > 0; j-- {
					k := temp[:j]
					mp[k] = mp[k] + tempLength - j + 1
					if mp[k] >= 3 && count < len(k) {
						count = len(k)
						break
					}
				}
			}
			sub = i
		}
	}
	return count
}
