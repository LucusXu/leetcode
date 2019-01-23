package main

import "fmt"
import "leetcode/sort"
import "leetcode/linkList"
import (
	"leetcode/string"
)

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

// 约瑟夫环算法
func testJCircle() {
	var last = linkList.JCircle(10, 5);
	fmt.Printf("%d\n", last)
}

// 约瑟夫环算法
func testLru() {
	var array = []int{1,3,7,6,4,2,8,9,10,5}
	var last = linkList.Lru(array, 10, 4);
	fmt.Printf("%d\n", last)
}

// 测试ip转换成整形
func testIp2Long() {
	var ipLong1, ipLong2 = string.Ip2Long("107.111.56.89");
	var ip = string.Long2Ip(ipLong1);
	fmt.Printf("%d\n", ipLong1, ipLong2, ip)
}

// 测试选择链表
func testRollback() {
	var array = []int{1,3,5,6,8,9,10,15,22,33,34,40,43,46}
	linkList.Rollback(array, 14, 4);
}

// 测试计数
func testUniqCount() {
	var array = []int{-5, -5, -4, -1, 1, 2, 3,4,5,6,8}
	var c = sort.UniqCount(array, 11);
	fmt.Printf("%d\n", c)
}

// 测试数组
func testMoreThanK() {
	var array = []int{1, 2, 2, 2, 2, 3,4,5,5,5,6,8,9,10}
	sort.MoreThanK(array, 14, 3);
}

// 测试数组
func testPalindrome() {
	str := "level"
	linkList.CreateLink(str)
}

func main() {
	// testSort()
	// testSearch()
	// testJCircle()
	// testLru()
	// testRollback()
	// testUniqCount()
	// testMoreThanK()
	testPalindrome()
}