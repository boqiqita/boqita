package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l242(t *testing.T) {
	assert.True(t, isAnagram("anagram", "nagaram"))
	assert.False(t, isAnagram("rat", "car"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	length, arr := len(s), make([]int, 27)
	for i := 0; i < length; i++ {
		arr[int(s[i]-'a')]++
		arr[int(t[i]-'a')]--
	}
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}
