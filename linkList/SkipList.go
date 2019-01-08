package linkList
// 跳表，Redis Sorted Set实现原理

import (
	"fmt"
)

func Create(array []int, length, n int) *LinkNode {
	if array == nil || length <= 0 || n <= 0 {
		return nil
	}
	head := new(LinkNode)

	record := make(map[int]bool, 0)
	fmt.Printf("%v", record)
	return head
}

func Search(array []int, length, n int) bool {
	return true
}

func Insert(array []int, length, n int) bool {
	return true
}

func Delete(array []int, length, n int) bool {
	return true
}
