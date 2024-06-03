package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 1、2、3、4、5、6、7、8、9

// 0、1、33、69，8
func Test_l246(t *testing.T) {
	assert.True(t, isStrobogrammatic("69"))
	assert.True(t, isStrobogrammatic("88"))
	assert.False(t, isStrobogrammatic("962"))
	assert.True(t, isStrobogrammatic("1"))
	assert.False(t, isStrobogrammatic("33"))
}

func isStrobogrammatic(num string) bool {
	length := len(num)
	if length == 0 {
		return false
	}
	l, r := 0, length-1
	for ; l < r; l, r = l+1, r-1 {
		if (num[l] == '0' && num[r] == '0') || (num[l] == '1' && num[r] == '1') || (num[l] == '8' && num[r] == '8') || (num[l] == '9' && num[r] == '6') || (num[l] == '6' && num[r] == '9') {
			continue
		}
		return false
	}
	if l == r {
		return num[l] == '0' || num[l] == '1' || num[l] == '8'
	}
	return true
}

//func Test_l247(t *testing.T) {
//
//}
//
//// 0,1,6,8,9
//func findStrobogrammatic(n int) []string {
//	switch n {
//	case 1:
//		return []string{"0", "1", "8"}
//	case 2:
//		return []string{"11", "69", "88", "96"}
//	case 3:
//		return []string{"101", "609", "808", "906", "111", "619", "818", "916", "181", "689", "888", "986"}
//	default:
//		l, r := []string{"1", "6", "8", "9"}, []string{"1", "9", "8", "6"}
//
//	}
//}

//func Test_l248(t *testing.T) {
//
//}
//
//func strobogrammaticInRange(low string, high string) int {
//
//}
