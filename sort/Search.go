package sort

import (
	// "fmt"
	"fmt"
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

/**
 * 查找出现次数大于n/k的数
 */
func MoreThanK(array []int, n, k int) {
	if n <= k || array == nil {
		return
	}
	hashMap := make(map[int]int, k)
	cnt := 0
	fmt.Printf("%d, %d, %d\n", n, k, n / k)
	fmt.Printf("%d\n", len(hashMap))

	for _, i := range array {
		// fmt.Printf("%d, %d\n", i, idx)
		hashMap[i] = hashMap[i] + 1
		count := countHash(hashMap)
		fmt.Printf("%d\n", count)

		// 清理一次hashmap
		if count == k {
			for key, v := range hashMap {
				if v > 0 {
					hashMap[key] = v - 1
				}
			}
			cnt++;
			if cnt > n / k +  1 {
				break
			}
		}
	}
	fmt.Printf("%d\n", cnt)
	if cnt < n / k + 1 {
		for key, v := range hashMap {
			if v > 0 {
				hashMap[key] = v - 1
			}
		}
	}
	for key, v := range hashMap {
		if v > 0 {
			fmt.Printf("%d, %d\n", key, v)
		}
	}
}

func countHash(hashMap map[int]int) int {
	c := 0
	for _, v := range hashMap {
		if v > 0 {
			c++
		}
	}
	return c
}
