package sort

import (
	// "fmt"
)

/**
 * 二分查找算法循环实现
 */
func BinSearchLoop(array []int, length, target int) bool {
	if length <= 0 || array == nil {
		return false
	}
	Bubble(array, length)
	var l int = 0
	var h int = length - 1
	// 防止溢出，右移比/要更快
	var mid int = l + (h - l) >> 1
	for l <= h {
		if array[mid] == target {
			return true
		} else if array[mid] < target {
			l = mid + 1
		} else if array[mid] > target {
			h = mid - 1
		}
		mid = l + (h - l) >> 1
	}
	return false
}

/**
 * 二分查找算法递归实现
 */
func BinSearchRescu(array []int, length, target int) bool {
	if length <= 0 || array == nil {
		return false
	}
	Qsort(array, length)
	var l int = 0
	var h int = length - 1
	// 防止溢出，右移比/要更快
	var mid int = l + (h - l) >> 1
	if array[mid] == target {
		return true
	}
	if l == h {
		return false
	}
	if array[mid] < target {
		BinSearchRescu(array[mid + 1:], length - mid - 1, target)
	} else if array[mid] > target {
		BinSearchRescu(array[:mid], mid, target)
	}
	return false
}
