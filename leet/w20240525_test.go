package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_100309(t *testing.T) {
	assert.Equal(t, 1, duplicateNumbersXOR([]int{1, 2, 1, 3}))
	assert.Equal(t, 1, duplicateNumbersXOR([]int{1, 2, 2, 1}))
}

func duplicateNumbersXOR(nums []int) int {
	f, s, temp := 0, 0, 0
	for _, num := range nums {
		temp = 1 << (num - 1)
		if temp&f == 0 {
			f += temp
		} else {
			s = s ^ num
		}
	}
	return s
}

//func Test_100303(t *testing.T) {
//
//}
//
//func occurrencesOfElement(nums []int, queries []int, x int) []int {
//
//}
//
//func Test_100323(t *testing.T) {
//	assert.Equal(t, 5, numberOfPairs([]int{1, 3, 4}, []int{1, 3, 4}, 1))
//}
//
//func numberOfPairs(nums1 []int, nums2 []int, k int) int {
//
//}

func Test_100307(t *testing.T) {
	assert.Equal(t, 7, minimumChairs("EEEEEEE"))
	assert.Equal(t, 2, minimumChairs("ELELEEL"))
	assert.Equal(t, 2, minimumChairs("ELEELEELLL"))
}

func minimumChairs(s string) int {

}
