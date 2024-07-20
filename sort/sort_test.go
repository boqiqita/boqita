package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_Sort(t *testing.T) {
	assert.True(t, sortCheck(bubbleSort([]int{8, 2, 3, 4, 3, 6, 6, 3, 1})))    //两两交换，先找出最大值或最小值
	assert.True(t, sortCheck(selectionSort([]int{8, 2, 3, 4, 3, 6, 6, 3, 1}))) //首位与其他位比较，先找出最大或最小值
	assert.True(t, sortCheck(insertionSort([]int{8, 2, 3, 4, 3, 6, 6, 3, 1}))) // 最新位与已排序的数组倒序比较并插入
	assert.True(t, sortCheck(shellSort([]int{8, 2, 3, 4, 3, 6, 6, 3, 1})))     // 分组按列插入排序
	assert.True(t, sortCheck(mergeSort([]int{8, 2, 3, 4, 3, 6, 6, 3, 1})))     //拆分成最小数据，再2个数组合并到一个新数据
	assert.True(t, sortCheck(mergeSort1([]int{8, 2, 3, 4, 3, 6, 6, 3, 1})))
	assert.True(t, sortCheck(quickSort([]int{8, 2, 3, 4, 3, 6, 6, 3, 2, 1}))) //取基准值并交换到数组左边，快慢指针将小于基准值的数交换到左边
	assert.True(t, sortCheck(heapSort([]int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9, 17, 44})))
	assert.True(t, sortCheck(bucketSort([]int{29, 25, 3, 49, 9, 37, 21, 43})))                          //数据按范围拆分为多个数组分别排序
	assert.True(t, sortCheck(countSort([]int{8, 2, 3, 4, 3, 6, 6, 3, 1})))                              //新数组下标值为旧数组的数，新数组数为旧数据相同数的个数
	assert.True(t, sortCheck(radixSort([]int{882, 3, 5, 345, 254, 606, 588, 808, 535, 784, 715, 710}))) // 按位数排序
}

// 冒泡排序 O(n) -> O(n^2)  重复地走访过要排序的数列，一次比较两个元素，如果它们的顺序错误就把它们交换过来
func bubbleSort(nums []int) []int {
	for i, j, flag := 0, 0, true; i < len(nums)-1 && flag; i++ {
		for j, flag = 0, false; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1], flag = nums[j+1], nums[j], true
			}
		}
	}
	return nums
}

// 选择排序 O(n^2) 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾
func selectionSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

// 插入排序 O(n) -> O(n^2) 工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入
func insertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j-1] > nums[j]; j-- { // 5，3
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
	return nums
}

// 希尔排序(插入排序的改进）O(n log n)到O(n2) 把记录按下标的一定增量分组，对每组使用直接插入排序算法排序
// 13 14 94 33 82
// 25 59 94 65 23
// 45 27
func shellSort(nums []int) []int {
	for step := len(nums) / 2; step > 0; step /= 2 { //外层步长控制
		//分组遍历
		for i := step; i < len(nums); i += step {
			//分组插入排序
			for j := i; j-step >= 0 && nums[j-step] > nums[j]; j-- {
				nums[j-step], nums[j] = nums[j], nums[j-step]
			}
		}
	}
	return nums
}

// 归并排序（迭代）O(n log n) 双指针合并到一个新数组里，最后返回新数据
func mergeSort(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}

	mid := length / 2
	left, right := nums[:mid], nums[mid:]
	leftLen, rigthLen := mid, length-mid

	if leftLen > 1 {
		left = mergeSort(left)
	}
	if rigthLen > 1 {
		right = mergeSort(right)
	}

	result := make([]int, length)
	for l, r, sub := 0, 0, 0; sub < length; sub++ {
		if l < leftLen && (r == rigthLen || left[l] <= right[r]) {
			result[sub], l = left[l], l+1
		} else {
			result[sub], r = right[r], r+1
		}
	}
	return result
}

// 归并排序(递归)
func mergeSort1(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	mid := length / 2
	return merge(mergeSort1(nums[:mid]), mergeSort1(nums[mid:]))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	for l, r, i := 0, 0, 0; i < len(result); i++ {
		if l < len(left) && (r == len(right) || left[l] <= right[r]) {
			result[i], l = left[l], l+1
		} else {
			result[i], r = right[r], r+1
		}
	}
	return result
}

// 快速排序 O(n log n) -> O(n^2) 从数列中挑出一个元素，称为 "基准"（pivot）,重新排序数列，所有比基准值小的元素摆放在基准前面，所有比基准值大的元素摆在基准后面,递归排序子序列
func quickSort(nums []int) []int {
	return quick(nums, 0, len(nums)-1)
}

func quick(nums []int, left, right int) []int {
	if left < right {
		pivotNewIndex := partition(nums, left, right) //重新排序数列，所有比基准值小的元素摆放在基准前面
		quick(nums, left, pivotNewIndex-1)
		quick(nums, pivotNewIndex+1, right)
	}
	return nums
}

func partition(nums []int, l, r int) int {
	pivotIndex := l
	l = l + 1
	for l < r {
		for l < r && nums[l] < nums[pivotIndex] {
			l++
		}
		for l < r && nums[r] >= nums[pivotIndex] {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	if nums[l] >= nums[pivotIndex] {
		l--
	}
	nums[pivotIndex], nums[l] = nums[l], nums[pivotIndex]
	return l
}

func partition1(nums []int, l, r int) int {
	pivot, storeIndex := nums[l], l+1 // 从数列中挑出一个元素，称为 "基准"（pivot）
	for i := storeIndex; i <= r; i++ {
		if nums[i] < pivot {
			if storeIndex < i {
				nums[storeIndex], nums[i] = nums[i], nums[storeIndex]
			}
			storeIndex++
		}
	}
	nums[l], nums[storeIndex-1] = nums[storeIndex-1], nums[l]
	return storeIndex - 1
}

// 堆排序O(n log n) 一个近似完全二叉树的结构，子结点的键值或索引总是小于（或者大于）它的父节点（大顶堆:升序 小顶堆：降序）
// 类似选择排序，根节点替换成最大值或者最小值，然后根节点放到最后面，length-1
func heapSort(nums []int) []int {
	length := len(nums)
	for i := length/2 - 1; i >= 0; i-- {
		heap(nums, i, length-1)
	}
	fmt.Println(nums)
	for i := length - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heap(nums, 0, i-1)
	}
	return nums
}

func heap(nums []int, dad, end int) {
	son := 2*dad + 1
	if son > end {
		return
	}
	if son+1 <= end && nums[son] < nums[son+1] {
		son++
	}
	if nums[dad] < nums[son] {
		nums[dad], nums[son] = nums[son], nums[dad]
		heap(nums, son, end)
	}
}

// 桶排序 Ο(n+k)工作的原理是将数组分到有限数量的桶里。每个桶再各自排序（不限制）
func bucketSort(nums []int) []int {
	bucket, maxNum := make([][]int, 10), 0
	for _, num := range nums {
		temp := num / 10
		bucket[temp] = append(bucket[temp], num)
		if maxNum < num {
			maxNum = num
		}
	}
	length := maxNum / 10
	for i, count := 0, 0; i <= length; i++ {
		sort.Ints(bucket[i])
		for j := 0; j < len(bucket[i]); j++ {
			nums[count], count = bucket[i][j], count+1
		}
	}
	return nums
}

// 计数排序 Ο(n+k)
func countSort(nums []int) []int {
	arr, maxnum, count := make([]int, 10), 0, 0 // 排序的数字最大值
	for _, num := range nums {
		arr[num]++
		if maxnum < num {
			maxnum = num
		}
	}

	for i := 0; i <= maxnum; i++ {
		for j := arr[i]; j > 0; j-- {
			nums[count], count = i, count+1
		}
	}
	return nums
}

// 基数排序(计数排序改良) Ο(n * k) 将所有待比较数值（正整数）统一为同样的数位长度，数位较短的数前面补零。然后，从最低位开始，依次进行一次排序。
func radixSort(nums []int) []int {
	radixs, flag, unit := make([][]int, 10), true, 1
	for flag {
		radixs, flag, unit = make([][]int, 10), false, unit*10
		for _, num := range nums {
			temp := (num % unit) / (unit / 10)
			if temp > 0 {
				flag = true
			}
			radixs[temp] = append(radixs[temp], num)
		}
		for count, i := 0, 0; i < len(radixs); i++ {
			for j := 0; j < len(radixs[i]); j++ {
				nums[count], count = radixs[i][j], count+1
			}
		}

	}
	return nums
}

func sortCheck(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			fmt.Println(nums)
			return false
		}
	}
	return true
}
