package main

import "fmt"
import "leetcode/sort"

// 排序算法
func testSort() {
	var array = []int{1,3,7,6,4,2,8,9,10,5}
	sort.Bubble(array, len(array))
	fmt.Printf("%v", array)
	sort.Qsort(array, len(array))
	fmt.Printf("%v", array)
}

// 查找算法
func testSearch() {
	var array = []int{1,3,7,6,4,2,8,9,10,5}
	res := sort.BinSearchLoop(array, len(array), 7)
	fmt.Printf("%v", res)
	res = sort.BinSearchRescu(array, len(array), 15)
	fmt.Printf("%v", res)
}

func main() {
	// testSort()
	testSearch()
}