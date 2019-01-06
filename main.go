package main

import "fmt"
import "leetcode/sort"

func testSort() {
	var array = []int{1,3,7,6,4,2,8,9,10,5}
	// sort.Bubble(array, len(array))
	// fmt.Printf("%v", array)
	sort.Qsort(array, len(array))
	fmt.Printf("%v", array)
}

func main() {
	// 排序算法
	// testSort()
	var a = []int{1,2,3,4,5}
	fmt.Printf("%v", a[:1])
	fmt.Printf("%v", a[1:3])
}