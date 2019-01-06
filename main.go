package main

import "fmt"
import "leetcode/sort"

func testSort() {
	var array = []int{1,3,7,6,4,2,8,9,10,5}
	sort.Bubble(array, len(array))
	fmt.Printf("%v", array)
}

func main() {
	// 排序算法
	testSort()
}