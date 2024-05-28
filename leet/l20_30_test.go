package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l28(t *testing.T) {
	assert.Equal(t, 0, strStr("sadbutsad", "sad"))
	assert.Equal(t, -1, strStr("leetcode", "leeto"))
}

func strStr(haystack string, needle string) int {
	l1, l2 := len(haystack), len(needle)
	for i := 0; i < l1; i++ {
		if l1-i < l2 {
			return -1
		}
		b := true
		for j := 0; j < l2; j++ {
			if haystack[i+j] != needle[j] {
				b = false
				break
			}
		}
		if b {
			return i
		}
	}
	return -1
}
