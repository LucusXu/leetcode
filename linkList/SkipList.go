package linkList
// 跳表，Redis Sorted Set实现原理

import (
	"fmt"
)

// 初始化创建一个跳表
func Create(array []int, length, n int) *LinkNode {
	if array == nil || length <= 0 || n <= 0 {
		return nil
	}
	head := new(LinkNode)

	fmt.Printf("%v", head)
	return head
}

func Insert(array []int, length, n int) bool {
	return true
}

func Delete(array []int, length, n int) bool {
	return true
}
