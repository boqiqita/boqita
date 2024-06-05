package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l67(t *testing.T) {
	assert.Equal(t, "100", addBinary("11", "1"))
	assert.Equal(t, "10101", addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	alen, blen, x, y, z := len(a)-1, len(b)-1, false, false, false
	arr := make([]rune, max(alen+2, blen+2))
	for i := len(arr) - 1; i >= 0; i, alen, blen = i-1, alen-1, blen-1 {
		if alen < 0 || a[alen] == '0' {
			x = false
		} else {
			x = true
		}
		if blen < 0 || b[blen] == '0' {
			y = false
		} else {
			y = true
		}
		if x && y && z {
			arr[i] = '1'
		} else if (x && y) || ((x || y) && z) {
			arr[i] = '0'
			z = true
		} else if x || y || z {
			arr[i] = '1'
			z = false
		} else {
			arr[i] = '0'
		}
	}

	if arr[0] == '0' {
		return string(arr[1:])
	}
	return string(arr)
}
